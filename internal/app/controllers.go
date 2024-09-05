package app

import (
	"database/sql"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
)

func (a *App) ListTask(w http.ResponseWriter, r *http.Request) {
	a.logger.Debug("GET /tasks")

	tasks, err := a.getAllTasks()
	if err != nil {
		a.JSONResponse(w, "server error", 500)
		return
	}

	a.JSONResponse(w, tasks, 200)
}

func (a *App) RetrieveTask(w http.ResponseWriter, r *http.Request) {
	a.logger.Debug("GET /tasks/{id}")

	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		a.JSONResponse(w, "incorrect id, id must be integer", 400)
		return
	}

	task, err := a.getOneTask(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			a.JSONResponse(w, "task was not found", 404)
			return
		}

		a.JSONResponse(w, "server error", 500)
		return
	}

	a.JSONResponse(w, task, 200)
}

func (a *App) CreateTask(w http.ResponseWriter, r *http.Request) {
	a.logger.Debug("POST /tasks")

	defer r.Body.Close()
	data, err := io.ReadAll(r.Body)
	if err != nil {
		a.JSONResponse(w, "server error", 500)
		return
	}

	var obj reqCreateUpdateTask
	if err = json.Unmarshal(data, &obj); err != nil {
		a.JSONResponse(w, "incorrect input data", 400)
		return
	}

	// validate input data

	task, err := a.formAndAddTask(obj)
	if err != nil {
		a.JSONResponse(w, "server error", 500)
		return
	}

	a.JSONResponse(w, task, 201)
}

func (a *App) UpdateTask(w http.ResponseWriter, r *http.Request) {
	a.logger.Debug("PUT /tasks/{id}")

	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		a.JSONResponse(w, "incorrect id, id must be integer", 400)
		return
	}

	defer r.Body.Close()
	data, err := io.ReadAll(r.Body)
	if err != nil {
		a.JSONResponse(w, "server error", 500)
		return
	}

	var obj reqCreateUpdateTask
	if err = json.Unmarshal(data, &obj); err != nil {
		a.JSONResponse(w, "incorrect input data", 400)
		return
	}

	task, err := a.checkExistsAndUpdateTask(id, obj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			a.JSONResponse(w, "task was not found", 404)
			return
		}

		a.JSONResponse(w, "server error", 500)
		return
	}

	a.JSONResponse(w, task, 200)
}

func (a *App) DeleteTask(w http.ResponseWriter, r *http.Request) {
	a.logger.Debug("DELETE /tasks/{id}")

	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		a.JSONResponse(w, "incorrect id, id must be integer", 400)
		return
	}

	if err := a.checkExistsAndDeleteTask(id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			a.JSONResponse(w, "task was not found", 404)
			return
		}

		a.JSONResponse(w, "server error", 500)
		return
	}

	a.JSONResponse(w, nil, 204)

}

func (a *App) JSONResponse(w http.ResponseWriter, data any, code int) {
	var isErr bool
	if code >= 400 {
		isErr = true
	}

	resp := response{
		Error:  isErr,
		Result: data,
	}

	w.Header().Set("Content-Type", "application/json")
	jsonData, err := json.Marshal(resp)
	if err != nil {
		a.logger.Error("json marshaling error: " + err.Error())
		resp = response{
			Error:  true,
			Result: "server error",
		}

		jsonData, _ := json.Marshal(resp)

		w.WriteHeader(500)
		w.Write(jsonData)
		return
	}

	w.WriteHeader(code)
	w.Write(jsonData)
}

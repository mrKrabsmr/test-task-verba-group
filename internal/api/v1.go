package api

import (
	"net/http"

	"github.com/mrKrabsmr/test-task-verba-group/internal/app"
)

func ConfigureV1Routes(router *http.ServeMux, a *app.App) {
	router.HandleFunc("GET /api/v1/tasks", a.ListTask)
	router.HandleFunc("GET /api/v1/tasks/{id}", a.RetrieveTask)
	router.HandleFunc("POST /api/v1/tasks", a.CreateTask)
	router.HandleFunc("PUT /api/v1/tasks/{id}", a.UpdateTask)
	router.HandleFunc("DELETE /api/v1/tasks/{id}", a.DeleteTask)
}

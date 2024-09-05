package app

import "time"

func (a *App) getAllTasks() ([]Task, error) {
	t, err := a.getTasksSQL()
	if err != nil {
		a.logger.Error("Do getTasks | GOT error: " + err.Error())
		return nil, err
	}

	return t, nil
}

func (a *App) getOneTask(id int) (Task, error) {
	t, err := a.getTaskSQL(id)
	if err != nil {
		a.logger.Error("Do getTask | GOT error: " + err.Error())
		return Task{}, err
	}

	return t, nil
}

func (a *App) formAndAddTask(data reqCreateUpdateTask) (Task, error) {
	task := Task{
		Title:       data.Title,
		Description: data.Description,
		DueDate:     data.DueDate,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	id, err := a.createTaskSQL(task)
	if err != nil {
		a.logger.Error("Do createTask | GOT error: " + err.Error())
		return Task{}, err
	}

	task.ID = id
	return task, nil
}

func (a *App) checkExistsAndUpdateTask(id int, data reqCreateUpdateTask) (Task, error) {
	task, err := a.getOneTask(id)
	if err != nil {
		return task, err
	}

	task.Title = data.Title
	task.Description = data.Description
	task.DueDate = data.DueDate
	task.UpdatedAt = time.Now()

	if err := a.updateTaskSQL(task); err != nil {
		a.logger.Error("Do updateTask | GOT error: " + err.Error())
		return Task{}, err
	}

	return task, nil
}

func (a *App) checkExistsAndDeleteTask(id int) error {
	_, err := a.getOneTask(id)
	if err != nil {
		return err
	}

	if err := a.deleteTaskSQL(id); err != nil {
		a.logger.Error("Do deleteTask | GOT error: " + err.Error())
		return err
	}

	return nil
}

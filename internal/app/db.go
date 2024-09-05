package app

func (a *App) InitDB() {
	query := `
	CREATE TABLE tasks(
		id serial PRIMARY KEY,
		title TEXT,
		description TEXT,
		due_date DATE,
		created_at TIMESTAMP,
		updated_at TIMESTAMP 
	)
	`

	if _, err := a.db.Exec(query); err != nil {
		panic(err)
	}
}

func (a *App) getTasksSQL() ([]Task, error) {
	var tasks []Task

	query := "SELECT * FROM tasks"

	if err := a.db.Select(&tasks, query); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (a *App) getTaskSQL(id int) (Task, error) {
	var task Task

	query := "SELECT * FROM tasks where id = $1"

	if err := a.db.Get(&task, query, id); err != nil {
		return task, err
	}

	return task, nil
}

func (a *App) createTaskSQL(task Task) (int, error) {
	query := `
	INSERT INTO tasks(title, description, due_date, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id	
	`
	id := 0
	if err := a.db.QueryRowx(
		query, task.Title, task.Description, task.DueDate, task.CreatedAt, task.UpdatedAt,
	).Scan(&id); err != nil {
		return 0, err
	}

	return int(id), nil
}

func (a *App) updateTaskSQL(task Task) error {
	query := `
	UPDATE tasks SET title = $1, description = $2, due_date = $3, created_at = $4, updated_at=$5 WHERE id = $6	
	`
	if _, err := a.db.Exec(
		query, task.Title, task.Description, task.DueDate, task.CreatedAt, task.UpdatedAt, task.ID,
	); err != nil {
		return err
	}

	return nil
}

func (a *App) deleteTaskSQL(id int) error {
	query := "DELETE FROM tasks where id = $1"

	if _, err := a.db.Exec(query, id); err != nil {
		return err
	}

	return nil
}

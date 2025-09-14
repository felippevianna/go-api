package models

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status   bool   `json:"status"`
}

const  (
	TABLE_NAME = "tasks"
	CREATE_TABLE_QUERY = `CREATE TABLE IF NOT EXISTS tasks (
		id SERIAL PRIMARY KEY,
		title VARCHAR(100) NOT NULL,
		description VARCHAR(255),
		status BOOLEAN DEFAULT FALSE
	);`
	INSERT_TASK_QUERY = `INSERT INTO tasks (title, description, status) VALUES ($1, $2, $3) RETURNING id;`
	GET_TASKS_QUERY = `SELECT id, title, description, status FROM tasks;`
	// GET_TASK_BY_ID_QUERY = `SELECT id, title, description, status FROM tasks WHERE id = $1;`
	// UPDATE_TASK_QUERY = `UPDATE tasks SET title = $1, description = $2, status = $3 WHERE id = $4;`
	// DELETE_TASK_QUERY = `DELETE FROM tasks WHERE id = $1;`
)
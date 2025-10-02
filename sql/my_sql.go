package database

import (
	"database/sql"
	"fmt"
	"tdl-go/config"
	userinterface "tdl-go/user_interface"
)

type Database struct {
	conn *sql.DB
}

func GetConnection(cfg *config.DatabaseConfig) (*Database, error) {
	// Create connection string
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &Database{conn: db}, nil
}

func (d *Database) Close() error {
	return d.conn.Close()
}

func (d *Database) GetAllTasks() ([]*userinterface.Todo, error) {
	query := "SELECT * FROM tasks ORDER BY due_date ASC;"

	rows, err := d.conn.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query tasks: %w", err)
	}
	defer rows.Close()

	var tasks []*userinterface.Todo
	for rows.Next() {
		var task userinterface.Todo

		err := rows.Scan(&task.Id, &task.Description, &task.Completed, &task.DueDate)
		if err != nil {
			return nil, fmt.Errorf("failed to scan task: %w", err)
		}

		tasks = append(tasks, &task)
	}

	return tasks, nil
}

func (d *Database) GetByStatus(status int) ([]*userinterface.Todo, error) {
	query := "SELECT * FROM tasks WHERE completed = ? ORDER BY due_date ASC;"

	rows, err := d.conn.Query(query, status)
	if err != nil {
		return nil, fmt.Errorf("failed to query tasks: %w", err)
	}
	defer rows.Close()

	var tasks []*userinterface.Todo
	for rows.Next() {
		var task userinterface.Todo

		err := rows.Scan(&task.Id, &task.Description, &task.Completed, &task.DueDate)
		if err != nil {
			return nil, fmt.Errorf("failed to scan task: %w", err)
		}

		tasks = append(tasks, &task)
	}

	return tasks, nil
}

func (d *Database) AddTask(task *userinterface.Todo) error {

	query := "INSERT INTO tasks (description, completed, due_date) VALUES (?, ?, ?);"

	_, err := d.conn.Exec(query, task.Description, task.Completed, task.DueDate.Format("2006-01-02"))
	if err != nil {
		return fmt.Errorf("failed to add task: %w", err)
	}

	fmt.Println("Task successfully added to the database!")
	return nil

}

func (d *Database) UpdateTaskCompletion(taskId int) error {
	query := "UPDATE tasks SET completed = 1 WHERE task_id = ? AND completed = 0;"

	result, err := d.conn.Exec(query, taskId)
	if err != nil {
		return fmt.Errorf("failed to update task: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rows == 0 {
		return fmt.Errorf("no incomplete task found with ID: %d", taskId)
	}

	fmt.Println("Task marked as completed!")
	return nil
}

func (d *Database) DeleteTask(taskId int) error {
	query := "DELETE FROM tasks WHERE task_id = ?;"

	result, err := d.conn.Exec(query, taskId)
	if err != nil {
		return fmt.Errorf("failed to delete task: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rows == 0 {
		return fmt.Errorf("could not delete, no task found with that ID: %w", err)
	}

	fmt.Println("Task was successfully deleted!")
	return nil
}

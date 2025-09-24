package database

import (
	"database/sql"
	"fmt"
	userinterface "tdl-go/user_interface"
)

type Database struct {
	conn *sql.DB
}

func GetConnection(username, password, host, port, dbname string) (*Database, error) {
	// Create connection string
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", username, password, host, port, dbname)

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

func (d *Database) AddTask(task *userinterface.Todo) error {

	query := "INSERT INTO tasks (description, completed, due_date) VALUES (?, ?, ?);"

	_, err := d.conn.Exec(query, task.Description, task.Completed, task.DueDate.Format("2006-01-02"))
	if err != nil {
		return fmt.Errorf("failed to add task: %w", err)
	}

	fmt.Println("Task successfully added to the database!")
	return nil

}

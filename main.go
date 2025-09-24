package main

import (
	"fmt"
	"log"
	"strings"
	"tdl-go/config"
	database "tdl-go/sql"
	userinterface "tdl-go/user_interface"

	_ "github.com/go-sql-driver/mysql"
)

var db *database.Database

func main() {

	// Load enviornment variables from .env
	if err := config.LoadEnv(".env"); err != nil {
		log.Printf("Warning: Could not load .env file: %v", err)
		log.Println("Using default/system environment variables...")
	}

	// Get database configurations from env variables
	dbConfig := config.GetDatabaseConfig()

	// Initialize database connection with configuration
	var err error
	db, err = database.GetConnection(
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	fmt.Printf("\t-----WELCOME TO YOUR TO DO LIST-----\n")

	ifContinue := true

	for ifContinue {
		userChoice := userinterface.DisplayMain()

		switch userChoice {
		case 1:
			addNewTask()
		case 2:
			checkOffTask()
		case 3:
			viewAllTasks()
		case 4:
			viewAllPending()
		case 5:
			deleteTask()
		case 0:
			fmt.Println("Goodbye!")
			ifContinue = false
		default:
			fmt.Println("ERROR! Please enter a valid option!!!")
		}
	}

}

func addNewTask() {
	newTask := userinterface.DisplayAddTask()

	// Add task to database
	if err := db.AddTask(newTask); err != nil {
		fmt.Println("Error adding task to database:", err)
		return
	}

	fmt.Println("Description:", newTask.Description)
	fmt.Println("Status: ❌ Pending")
	fmt.Println("Due Date:", newTask.DueDate.Format("01-02-2006"))
}

func checkOffTask() {
	viewAllPending()
	fmt.Println("\n" + strings.Repeat("_", 45))

	fmt.Println("Enter the ID of the task you have completed:")
	var taskId int
	fmt.Scanln(&taskId)

	err := db.UpdateTaskCompletion(taskId)
	if err != nil {
		fmt.Println("ERROR! Could not update task:", err)
	}
}

func viewAllTasks() {
	allTasks, err := db.GetAllTasks()
	if err != nil {
		fmt.Println("Error recieving tasks:", err)
		return
	}

	if len(allTasks) == 0 {
		fmt.Println("No tasks found...")
		return
	}

	printData(allTasks)
}

func viewAllPending() {
	allUnfinished, err := db.GetAllPending()
	if err != nil {
		fmt.Println("Error getting all unfinished tasks:", err)
		return
	}

	if len(allUnfinished) == 0 {
		fmt.Println("You have no unfinished tasks!")
		return
	}

	printData(allUnfinished)
}

func deleteTask() {
	viewAllTasks()

	fmt.Println("Enter the ID of the task you want to delete:")
	var taskId int
	fmt.Scanln(&taskId)

	err := db.DeleteTask(taskId)
	if err != nil {
		fmt.Println("ERROR! Could not delete that task:", err)
	}
}

func printData(tasks []*userinterface.Todo) {
	for _, task := range tasks {
		status := "❌ Pending"
		if task.Completed {
			status = "✅ Completed"
		}

		fmt.Printf("---TASK ID: %d---\n", task.Id)
		fmt.Println("Description:", task.Description)
		fmt.Println("Due Date:", task.DueDate.Format("01-02-2006"))
		fmt.Println("Status:", status)
		fmt.Println(strings.Repeat("_", 45))
	}
}

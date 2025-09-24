package main

import (
	"fmt"
	userinterface "tdl-go/user_interface"
)

func main() {
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
			viewAllCompleted()
		case 0:
			fmt.Println("Goodbye!")
			ifContinue = false
		default:
			fmt.Println("ERROR! Please enter a valid option!!!")
		}
	}

}

func addNewTask() {
	fmt.Println("Add a new task")
}

func checkOffTask() {
	fmt.Println("Check off task")
}

func viewAllTasks() {
	fmt.Println("View all tasks")
}

func viewAllCompleted() {
	fmt.Println("View all completed tasks")
}

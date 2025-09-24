package userinterface

import (
	"fmt"
	"strings"
	"time"
)

func DisplayMain() int {
	fmt.Println("---OPTIONS---")
	fmt.Println(strings.Repeat("_", 20))
	fmt.Println("1 - Add A New Task\n2 - Check Off A Task\n3 - View All To-Do's\n4 - View All Completed Tasks\n0 - Exit")

	fmt.Println("Enter option:")
	var userChoice int
	fmt.Scanln(&userChoice)

	return userChoice
}

type Todo struct {
	Description string
	Completed   bool
	DueDate     time.Time
}

func DisplayAddTask() *Todo {
	fmt.Println("\t---Add A Task---")
	fmt.Println(strings.Repeat("_", 20))

	fmt.Println("Enter a description for your task:")
	var description string
	fmt.Scanln(&description)

	fmt.Println("Enter the day of the month it is due:")
	var dayDue int
	fmt.Scanln(&dayDue)

	fmt.Println("Enter the month (numerically) it is due:")
	var monthDue time.Month
	fmt.Scanln(&monthDue)

	fmt.Println("Enter the year that it is due:")
	var yearDue int
	fmt.Scanln(&yearDue)

	dueDate := time.Date(yearDue, monthDue, dayDue, 0, 0, 0, 0, time.UTC)

	return &Todo{
		Description: description,
		Completed:   false,
		DueDate:     dueDate,
	}
}

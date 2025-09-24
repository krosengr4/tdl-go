package userinterface

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"tdl-go/utils"
	"time"
)

func DisplayMain() int {
	fmt.Println("---OPTIONS---")
	fmt.Println(strings.Repeat("_", 20))
	fmt.Println("1 - Add A New Task\n2 - Check Off A Task\n3 - View All Tasks\n4 - View All Unfinished Tasks\n0 - Exit")

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

// todo: Look at bufio.Scanner to better get user input

func DisplayAddTask() *Todo {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("\t---Add A Task---")
	fmt.Println(strings.Repeat("_", 30))

	fmt.Println("Enter your task description:")
	scanner.Scan()
	description := scanner.Text()

	dayDue := utils.GetValidatedNumber("Enter the day of the month your task is due (numerically):\n", 1, 31)
	monthDue := utils.GetValidatedNumber("Enter the month your task is due (numerically):\n", 1, 12)
	yearDue := utils.GetValidatedNumber("Enter the year your task is due:\n", 1999, 50000)

	dueDate := time.Date(yearDue, time.Month(monthDue), dayDue, 0, 0, 0, 0, time.UTC)

	// Return the time.Time object directly in the struct
	return &Todo{
		Description: description,
		Completed:   false,
		DueDate:     dueDate,
	}
}

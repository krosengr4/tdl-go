package userinterface

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

// todo: Look at bufio.Scanner to better get user input

func DisplayAddTask() *Todo {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("\t---Add A Task---")
	fmt.Println(strings.Repeat("_", 30))

	fmt.Println("Enter your task description:")
	scanner.Scan()
	description := scanner.Text()

	dayDue := getValidatedNumber("Enter the day of the month your task is due (numerically):\n", 1, 31)
	monthDue := getValidatedNumber("Enter the month your task is due (numerically):\n", 1, 12)
	yearDue := getValidatedNumber("Enter the year your task is due:\n", 1999, 50000)

	dueDate := time.Date(yearDue, time.Month(monthDue), dayDue, 0, 0, 0, 0, time.UTC)

	// Return the time.Time object directly in the struct
	return &Todo{
		Description: description,
		Completed:   false,
		DueDate:     dueDate,
	}
}

// Helper function to get a validated number within range
func getValidatedNumber(prompt string, min, max int) int {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(prompt)
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		if num, err := strconv.Atoi(input); err == nil {
			if num >= min && num <= max {
				return num
			}
			fmt.Printf("Number must be between %d and %d. Try again.\n", min, max)
		} else {
			fmt.Println("Invalid number. Please try again.")
		}
	}
}

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
	fmt.Println(strings.Repeat("_", 50))
	fmt.Println("\n---OPTIONS---")
	fmt.Println(strings.Repeat("_", 20))
	fmt.Println("1 - Add A New Task\n2 - Check Off A Task\n3 - View All Tasks\n4 - View All Pending Tasks\n5 - View All Completed Tasks\n6 - Delete A Task\n0 - Exit")

	fmt.Println("Enter option:")

	var userChoice int
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		userChoice, _ = strconv.Atoi(scanner.Text())
	}

	return userChoice
}

type Todo struct {
	Id          int
	Description string
	Completed   bool
	DueDate     time.Time
}

func DisplayAddTask() *Todo {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("\t---Add A Task---")
	fmt.Println(strings.Repeat("_", 30))

	fmt.Println("Enter your task description:")
	scanner.Scan()
	description := scanner.Text()

	var dueDate time.Time
	var err error

	for {
		fmt.Println("Enter the due date (mm-dd-yyyy): ")
		scanner.Scan()
		userDueDate := strings.TrimSpace(scanner.Text())

		dueDate, err = time.Parse("01-02-2006", userDueDate)
		if err == nil {
			break
		} else {
			fmt.Println("Invalid date format. Please use mm-dd-yyyy")
		}

	}

	// Return the time.Time object directly in the struct
	return &Todo{
		Description: description,
		Completed:   false,
		DueDate:     dueDate,
	}
}

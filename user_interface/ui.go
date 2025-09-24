package userinterface

import (
	"fmt"
	"strings"
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

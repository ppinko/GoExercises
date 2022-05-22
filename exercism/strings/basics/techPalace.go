package main

import (
	"fmt"
	"strings"
)

func WelcomeMessage(name string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(name)
}

func AddBorder(welcomeMessage string, numStart int) string {
	stars := strings.Repeat("*", numStart)
	return stars + "\n" + welcomeMessage + "\n" + stars
}

func CleanUpMessage(message string) string {
	return strings.Trim(message, "* \n\t")
}

func main() {
	// welcome message
	fmt.Println(WelcomeMessage("Judy"))

	// add border
	fmt.Println() // print empty line
	fmt.Println(AddBorder("Welcome!", 10))

	// clean up message
	message := `
	**************************
	*    BUY NOW, SAVE 10%   *
	**************************
	`
	fmt.Println() // print empty line
	fmt.Println(CleanUpMessage(message))
}

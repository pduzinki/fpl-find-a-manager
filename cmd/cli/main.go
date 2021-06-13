package main

import "fmt"

func main() {
	fmt.Println("Welcome to 'Find a manager' fpl app!")

	for {
		fmt.Println("Please type the name of the manager you're looking for, or [ctrl+c] to exit:")

		var nameInput string
		fmt.Scanln(&nameInput)

		// TODO remove later
		fmt.Println("hello " + nameInput)
		fmt.Println("No managers found!")

		fmt.Println("--------------------------------")
	}
}

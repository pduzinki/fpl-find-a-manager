package main

import (
	"fmt"
	"log"

	"fpl-find-a-manager/pkg/storage/sqlite"
)

func main() {
	fmt.Println("Welcome to 'Find a manager' fpl app!")

	s, err := sqlite.NewStorage()
	if err != nil {
		log.Fatalln("Failed to create storage!")
	}

	_ = s

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

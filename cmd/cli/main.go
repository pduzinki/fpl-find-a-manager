package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

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
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			nameInput = scanner.Text()
		}

		// TODO remove later
		fmt.Println("hello " + nameInput)
		fmt.Println("No managers found!")
		// TODO press enter to look for someone else, or esc to exit

		fmt.Println("--------------------------------")
	}
}

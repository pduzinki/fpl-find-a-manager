package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"fpl-find-a-manager/pkg/controllers"
	"fpl-find-a-manager/pkg/models"
)

func main() {
	fmt.Println("Welcome to 'Find a manager' fpl app!")

	ms, err := models.NewManagerService()
	if err != nil {
		log.Fatalln("Failed to init models service!")
	}

	mc := controllers.NewManagerController(ms)

	go mc.AddManagers()

	for {
		fmt.Println("Please type the name of the manager " +
			"you're looking for, or [ctrl+c] to exit:")

		var nameInput string
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			nameInput = scanner.Text()
		}

		m, err := mc.MatchManagersByName(nameInput)
		if err != nil {
			fmt.Println("Something went wrong!")
		} else if len(m) == 0 {
			fmt.Println("No managers found!")
		} else {
			fmt.Println(m)
		}
		// TODO press enter to look for someone else, or esc to exit
	}
}

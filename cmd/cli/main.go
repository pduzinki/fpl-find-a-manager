package main

import (
	"bufio"
	"fmt"
	"os"

	"fpl-find-a-manager/pkg/controllers"
	"fpl-find-a-manager/pkg/models"
)

func main() {
	fmt.Println("Welcome to 'Find a manager' fpl app!")

	ms, err := models.NewManagerService()
	if err != nil {
		fmt.Println("Failed to init models service!")
		panic(err)
	}

	mc := controllers.NewManagerController(ms)

	go mc.AddManagers()

	// s, err := sqlite.NewStorage()
	// if err != nil {
	// 	log.Fatalln("Failed to create storage!")
	// }

	// adder := adding.NewService(s)
	// lister := listing.NewService(s)
	// filler := filling.NewService(adder, lister)

	// go filler.Fill()
	// // go adder.AddAllManagers()

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

		// 	m, err := lister.GetManagersByName(nameInput)
		// 	if err != nil {
		// 		fmt.Println("Something went wrong!")
		// 	} else if len(m) == 0 {
		// 		fmt.Println("No managers found!")
		// 	} else {
		// 		fmt.Println(m)
		// 	}

		// 	// TODO press enter to look for someone else, or esc to exit
		// 	fmt.Println("--------------------------------")
	}
}

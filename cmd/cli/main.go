package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"fpl-find-a-manager/pkg/adding"
	"fpl-find-a-manager/pkg/listing"
	"fpl-find-a-manager/pkg/storage/sqlite"
	"fpl-find-a-manager/pkg/wrapper"
)

func main() {
	fmt.Println("Welcome to 'Find a manager' fpl app!")

	s, err := sqlite.NewStorage()
	if err != nil {
		log.Fatalln("Failed to create storage!")
	}

	adder := adding.NewService(s)
	lister := listing.NewService(s)

	wrapper := wrapper.NewWrapper()
	wm, err := wrapper.GetManager(43741)
	if err != nil {
		panic(err)
	}

	am := adding.Manager{
		FplID:    wm.ID,
		FullName: fmt.Sprintf("%s %s", wm.FirstName, wm.LastName),
	}

	adder.AddManager(am)

	for {
		fmt.Println("Please type the name of the manager you're looking for, or [ctrl+c] to exit:")

		var nameInput string
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			nameInput = scanner.Text()
		}

		m, err := lister.GetManagerByName(nameInput)
		if err != nil {
			fmt.Println("No managers found!")
		} else {
			fmt.Println(m)
		}

		// TODO press enter to look for someone else, or esc to exit
		fmt.Println("--------------------------------")
	}
}

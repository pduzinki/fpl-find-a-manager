package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"fpl-find-a-manager/pkg/config"
	"fpl-find-a-manager/pkg/controllers"
	"fpl-find-a-manager/pkg/models"
)

func main() {
	f, err := os.Create("log.txt")
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	log.SetOutput(f)
	// log.SetOutput(ioutil.Discard)

	fmt.Println("Welcome to 'Find a manager' fpl app!")

	cfg := config.Load()

	ms, err := models.NewManagerService(cfg.DBConfig)
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

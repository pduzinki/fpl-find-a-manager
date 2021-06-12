package main

import (
	"fmt"
	"log"

	"fpl-find-a-manager/pkg/storage/sqlite"
)

func main() {
	fmt.Println("hello from fpl-find-a-manager cli")

	s, err := sqlite.NewStorage()
	if err != nil {
		log.Fatalln("Failed to create storage!")
	}

	_ = s
}

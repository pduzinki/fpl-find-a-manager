package main

import (
	"log"
	"net/http"

	"fpl-find-a-manager/pkg/config"
	"fpl-find-a-manager/pkg/controllers"
	"fpl-find-a-manager/pkg/models"
	"fpl-find-a-manager/pkg/rest"
)

func main() {
	log.Println("fpl-find-a-manager app started")

	cfg := config.Load()

	ms, err := models.NewManagerService(cfg.DBConfig)
	if err != nil {
		log.Fatalln("Failed to init models service!")
	}

	mc := controllers.NewManagerController(ms)

	go mc.AddManagers()

	router := rest.Handler()

	log.Println("fpl-find-a-manager app now listening on port 3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}

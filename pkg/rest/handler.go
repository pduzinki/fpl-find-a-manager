package rest

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"fpl-find-a-manager/pkg/controllers"
	"fpl-find-a-manager/pkg/html"
)

//
func Handler(mc *controllers.ManagerController) http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/", Homepage()).Methods("GET")
	r.HandleFunc("/", FindManager(mc)).Methods("POST")

	r.PathPrefix("/static/").Handler(html.StaticFiles())

	return r
}

//
func Homepage() func(w http.ResponseWriter, r *http.Request) {
	page, err := html.NewPage("templates/homepage.html")
	if err != nil {
		log.Fatalln("Failed to parse homepage templates!")
	}

	return func(w http.ResponseWriter, r *http.Request) {

		page.Render(w)
	}
}

//
func FindManager(mc *controllers.ManagerController) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		var form searchForm

		err := parseForm(r, &form)
		if err != nil {
			// TODO handle error
		}

		managers, err := mc.MatchManagersByName(form.ManagerName)
		if err != nil {
			fmt.Println("Something went wrong!")
		} else if len(managers) == 0 {
			fmt.Println("No managers found!")
		} else {
			fmt.Printf("Found %v manager(s):\n", len(managers))
			for i, m := range managers {
				fmt.Printf("%v. %v https://fantasy.premierleague.com/entry/%v/history\n",
					i+1, m.FullName, m.FplID)
			}
		}
		fmt.Printf("\n")

	}
}

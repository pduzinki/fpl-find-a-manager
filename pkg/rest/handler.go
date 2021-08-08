package rest

import (
	"fpl-find-a-manager/pkg/html"
	"net/http"

	"github.com/gorilla/mux"
)

func Handler() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/", Homepage()).Methods("GET")

	r.PathPrefix("/static/").Handler(html.StaticFiles())

	return r
}

func Homepage() func(w http.ResponseWriter, r *http.Request) {
	page, err := html.NewPage("templates/homepage.html")
	if err != nil {
		// TODO proper error handling
		panic(err)
	}

	return func(w http.ResponseWriter, r *http.Request) {

		page.Render(w)
	}
}

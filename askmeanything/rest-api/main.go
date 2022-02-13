package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/skatsaounis/askmeanything/db"
	"github.com/skatsaounis/askmeanything/routes"
)

func main() {
	r := mux.NewRouter().StrictSlash(true)

	// Register routes
	routes.QuestionRoute(r)

	// Connect to database
	db.ConnectDB()

	srv := &http.Server{
		Handler: r,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

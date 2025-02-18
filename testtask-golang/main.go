package main

import (
	"log"
	"net/http"
	"testtask-golang/controllers"
	"testtask-golang/db"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize database connection
	if err := db.Connect(); err != nil {
		log.Fatal(err)
	}

	// Create router
	r := mux.NewRouter()

	// Register routes
	r.HandleFunc("/register", controllers.Register).Methods("POST")
	r.HandleFunc("/login", controllers.Login).Methods("POST")
	r.HandleFunc("/users", controllers.GetUsers).Methods("GET")

	// Start server
	log.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}

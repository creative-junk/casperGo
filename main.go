package main

import (
	"github.com/creative-junk/casperv1/api"
	"github.com/gorilla/handlers"
	"log"
	"net/http"
	"os"
)

func main() {
	//Get the port
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	//Initialize the System Routes
	router := api.NewRouter()

	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})

	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(allowedOrigins, allowedMethods)(router)))
}

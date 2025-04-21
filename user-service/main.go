package main

import (
	"log"
	"net/http"
	"user-service/config"
	"user-service/routes"

	"github.com/gorilla/mux"
)

func main() {

	config.ConnectDB()

	r := mux.NewRouter()
	routes.RegisterUserRoutes(r)

	log.Println("User service running on port : 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

package main

import (
	"log"
	"net/http"
	"oauth-service/routes"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterOAuthRoutes(r)

	log.Println("OAuth Service running on :8082")
	log.Fatal(http.ListenAndServe(":8082", r))
}

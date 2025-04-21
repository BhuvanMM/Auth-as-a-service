package main

import (
	"log"
	"net/http"
	"token-service/routes"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterTokenRoutes(r)

	log.Println("Token Service running on :8081")
	log.Fatal(http.ListenAndServe(":8081", r))
}

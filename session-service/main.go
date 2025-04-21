package main

import (
	"log"
	"net/http"
	"session-service/database"
	"session-service/rabbitmq"
	"session-service/routes"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	database.ConnectMongo()
	rabbitmq.ConsumeSessionEvents()

	routes.RegisterSessionRoutes(r)

	log.Println("🚀 Session Service running on :8083")
	log.Fatal(http.ListenAndServe(":8083", r))
}

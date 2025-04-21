package routes

import (
	"net/http"
	"token-service/controller"

	"github.com/gorilla/mux"
)

func RegisterTokenRoutes(r *mux.Router) {
	r.HandleFunc("/generate", controller.GenerateToken).Methods("POST")
	r.HandleFunc("/refresh", controller.ValidateToken).Methods("POST")
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Token Service is running"))
	}).Methods("GET")
}

package routes

import (
	"net/http"
	"session-service/controller"

	"github.com/gorilla/mux"
)

func RegisterSessionRoutes(r *mux.Router) {
	r.HandleFunc("/sessions", controller.GetSessions).Methods("GET")
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Session service alive"))
	})
}

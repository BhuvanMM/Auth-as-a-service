package routes

import (
	"net/http"
	"oauth-service/controller"

	"github.com/gorilla/mux"
)

func RegisterOAuthRoutes(r *mux.Router) {
	r.HandleFunc("/auth/google/login", controller.GoogleLogin)
	r.HandleFunc("/auth/google/callback", controller.GoogleCallBack)
	r.HandleFunc("/auth/github/login", controller.GithubLogin)
	r.HandleFunc("/auth/github/callback", controller.GithubCallBack)

	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OAuth Service is up"))
	})
}

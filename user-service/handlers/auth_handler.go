package handlers

import (
	"encoding/json"
	"net/http"
	"user-service/models"
	"user-service/services"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)

	err := services.RegisterUser(&user)
	if err != nil {
		http.Error(w, "Signup failed", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "Signup successful"})
}

func Login(w http.ResponseWriter, r *http.Request) {
	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)

	token, err := services.LoginUser(&user)
	if err != nil {
		http.Error(w, "Login failed", http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

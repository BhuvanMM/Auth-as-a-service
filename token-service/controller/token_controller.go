package controller

import (
	"encoding/json"
	"net/http"
	"token-service/models"
	"token-service/utils"
)

func GenerateToken(w http.ResponseWriter, r *http.Request) {
	var req models.TokenRequest
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil || req.UserID == "" {
		http.Error(w, "Invalid Request", http.StatusInternalServerError)
		return
	}

	access, refresh, err := utils.GenerateToken(req.UserID)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"access_token":  access,
		"refresh_token": refresh,
	})
}

func ValidateToken(w http.ResponseWriter, r *http.Request) {
	var req models.RefreshRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || req.RefreshToken == "" {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	claims, err := utils.ValidateToken(req.RefreshToken)
	if err != nil {
		http.Error(w, "Invalid refresh token", http.StatusUnauthorized)
		return
	}

	userID := claims["user_id"].(string)
	access, refresh, err := utils.GenerateToken(userID)
	if err != nil {
		http.Error(w, "Token refresh failed", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"access_token":  access,
		"refresh_token": refresh,
	})
}

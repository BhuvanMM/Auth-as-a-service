package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"oauth-service/config"
	"oauth-service/utils"
)

func GoogleLogin(w http.ResponseWriter, r *http.Request) {
	url := config.GoogleOAuthConfig.AuthCodeURL("state")
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func GoogleCallBack(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	token, err := config.GoogleOAuthConfig.Exchange(context.Background(), code)

	if err != nil {
		http.Error(w, "Token exchange failed", http.StatusInternalServerError)
		return
	}

	client := config.GoogleOAuthConfig.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		http.Error(w, "Failed to get user info", http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	var userInfo map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&userInfo)

	email := userInfo["email"].(string)

	access, refresh, err := utils.GenerateToken(email)
	if err != nil {
		http.Error(w, "JWT creation failed", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"access_token":  access,
		"refresh_token": refresh,
		"email":         email,
	})

}

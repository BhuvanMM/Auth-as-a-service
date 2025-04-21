package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"oauth-service/config"
	"oauth-service/utils"
)

func GithubLogin(w http.ResponseWriter, r *http.Request) {
	url := config.GitHubOAuthConfig.AuthCodeURL("state")
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func GithubCallBack(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	token, err := config.GitHubOAuthConfig.Exchange(context.Background(), code)

	if err != nil {
		http.Error(w, "Token exchange failed", http.StatusInternalServerError)
		return
	}

	client := config.GitHubOAuthConfig.Client(context.Background(), token)

	resp, err := client.Get("https://api.github.com/user/emails")
	if err != nil {
		http.Error(w, "Failed to fetch GitHub user", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var emails []map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&emails)
	email := emails[0]["email"].(string)

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

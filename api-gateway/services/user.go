package services

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type User struct {
	ID    string `json:"_id"`
	Email string `json:"email"`
}

func Signup(email, password string) (*User, error) {
	body, _ := json.Marshal(map[string]string{
		"email":    email,
		"password": password,
	})
	resp, err := http.Post("http://localhost:8080/signup", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var user User
	json.NewDecoder(resp.Body).Decode(&user)
	return &user, nil
}

func GetUser(email string) (*User, error) {
	resp, err := http.Get("http://localhost:8080/user?email=" + email)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var user User
	json.NewDecoder(resp.Body).Decode(&user)
	return &user, nil
}

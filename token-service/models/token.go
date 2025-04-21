package models

type TokenRequest struct {
	UserID string `json:"user_id"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}

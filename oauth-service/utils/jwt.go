package utils

import (
	"oauth-service/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userID string) (string, string, error) {
	accessClaims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(15 * time.Minute).Unix(),
	}

	refreshClaims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(7 * 24 * time.Hour).Unix(),
	}

	access := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	refresh := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)

	accessToken, err := access.SignedString(config.JWTSecret)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := refresh.SignedString(config.JWTSecret)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

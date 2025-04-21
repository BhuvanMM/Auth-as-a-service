package utils

import (
	"time"
	"token-service/config"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userId string) (string, string, error) {
	accessTokenClaims := jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(time.Minute * 15).Unix(),
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	accessTokenString, err := accessToken.SignedString(config.JWTSECRET)
	if err != nil {
		return "", "", err
	}

	refreshTokenClaims := jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(time.Hour * 24 * 7).Unix(),
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	refreshTokenString, err := refreshToken.SignedString(config.JWTSECRET)
	if err != nil {
		return "", "", err
	}

	return accessTokenString, refreshTokenString, nil

}

func ValidateToken(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return config.JWTSECRET, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	return token.Claims.(jwt.MapClaims), nil
}

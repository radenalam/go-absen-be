package utils

import (
	"go-absen-be/internal/entity"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

var jwtKey = []byte(viper.GetString("jwt.secret_key"))

func GenerateJWT(user entity.User) (string, time.Time, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &jwt.RegisteredClaims{
		ID:        user.ID.String(),
		Subject:   user.Username,
		ExpiresAt: jwt.NewNumericDate(expirationTime), 
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", time.Time{}, err
	}

	return tokenString, expirationTime, nil
}
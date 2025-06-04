package model

import "time"

type LoginResponse struct {
	Token string `json:"token"`
	Name string `json:"name"`
	Username string `json:"username"`
	Email *string `json:"email"`
	ExpiresAt time.Time	`json:"expires_at"`
	CreatedAt time.Time	`json:"created_at"`
}
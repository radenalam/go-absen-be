package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserResponse struct {
	ID       uuid.UUID `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    *string `json:"email"`	
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type VerifyUserRequest struct {
	Token string `validate:"required"`
}

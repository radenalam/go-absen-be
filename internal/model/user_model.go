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

type CreateUserRequest struct {
	Name     string  `json:"name" validate:"required"`
	Username string  `json:"username" validate:"required"`
	Email    *string `json:"email" validate:"omitempty,email"`
	Password string  `json:"password" validate:"required,min=8"`
}

type LoginUserRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
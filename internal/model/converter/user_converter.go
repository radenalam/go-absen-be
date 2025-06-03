package converter

import (
	"go-absen-be/internal/entity"
	"go-absen-be/internal/model"
)



func UserToResponse(user *entity.User) *model.UserResponse {
	return &model.UserResponse{
		ID :       user.ID,
		Username:  user.Username,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
	}
}

package usecase

import (
	"context"
	"go-absen-be/internal/entity"
	"go-absen-be/internal/model"
	"go-absen-be/internal/model/converter"
	"go-absen-be/internal/repository"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserUseCase struct {
	DB *gorm.DB	
	Log *logrus.Logger
	Validate *validator.Validate
	UserRepository *repository.UserRepository
}

func NewUserUseCase(db *gorm.DB, logger *logrus.Logger, validate *validator.Validate, userRepository *repository.UserRepository) *UserUseCase {
	return &UserUseCase{
		DB: db,
		Log: logger,
		Validate: validate,
		UserRepository: userRepository,
	}
} 

func(c *UserUseCase) Search(ctx context.Context, request *model.SearchRequest) ([]model.UserResponse, int64, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return nil, 0, fiber.ErrBadRequest
	}

	users, total, err := c.UserRepository.Search(tx, request)
	if err != nil {
		c.Log.Warnf("Failed search user : %+v", err)
		return nil, 0, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, 0, fiber.ErrInternalServerError
	}

	response := make([]model.UserResponse, len(users))
	for i, user := range users {
		response[i] = *converter.UserToResponse(&user)
	}

	return response, total, nil
}

func (c *UserUseCase) Verify(ctx context.Context, request *model.VerifyUserRequest) (*model.Auth, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	err := c.Validate.Struct(request)
	if err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return nil, fiber.ErrBadRequest
	}

	user_token := new(entity.RefreshToken)
	user_token.Token = request.Token

	// if err := c.UserTokenRepository.FindByToken(tx, user_token); err != nil {
	// 	c.Log.Warnf("Failed find user by token : %+v", err)
	// 	return nil, fiber.ErrNotFound
	// }

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return &model.Auth{ID: user_token.UserID}, nil
}
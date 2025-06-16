package usecase

import (
	"context"
	"go-absen-be/internal/entity"
	"go-absen-be/internal/model"
	"go-absen-be/internal/model/converter"
	"go-absen-be/internal/repository"
	"go-absen-be/internal/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserUseCase struct {
	DB *gorm.DB	
	Log *logrus.Logger
	Validate *validator.Validate
	UserRepository *repository.UserRepository
	RefreshTokenRepository repository.RefreshTokenRepository
}

func NewUserUseCase(db *gorm.DB, logger *logrus.Logger, validate *validator.Validate, userRepository *repository.UserRepository) *UserUseCase {
	return &UserUseCase{
		DB: db,
		Log: logger,
		Validate: validate,
		UserRepository: userRepository,
	}
} 

func (c *UserUseCase) Create(ctx context.Context, request *model.CreateUserRequest) (*model.UserResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	err := c.Validate.Struct(request)
	if err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return nil, fiber.ErrBadRequest
	}

	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		c.Log.Warnf("Failed to generate bcrype hash : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	user := &entity.User{
		Name: request.Name,
		Username: request.Username,
		Password: string(password),
		Email: request.Email,
	}

	if err := c.UserRepository.Create(tx, user); err != nil {
		c.Log.Warnf("Failed create user to database : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.UserToResponse(user), nil
}

func (c *UserUseCase) Login(ctx context.Context, request *model.LoginUserRequest) (*model.LoginResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.Warnf("Invalid request body  : %+v", err)
		return nil, fiber.ErrBadRequest
	}

	user := new(entity.User)
	if err := c.UserRepository.FindByUsername(tx, user, request.Username); err != nil {
		c.Log.Warnf("Failed find user by username : %+v", err)		
		return nil, fiber.ErrUnauthorized
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		c.Log.Warnf("Failed to compare user password with bcrype hash : %+v", err)
		return nil, fiber.ErrUnauthorized
	}

	token, expirationTime, err := utils.GenerateJWT(*user)
	if err != nil {
		c.Log.Warnf("Failed to generate JWT token : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	user_token := &entity.RefreshToken{
		ExpiresAt: expirationTime,
		UserID: user.ID,
		Token: token,
	}

	login_response := &model.LoginResponse{
		Token: token,
		Name: user.Name,
		Username: user.Username,
		Email: user.Email,
		ExpiresAt: user_token.ExpiresAt,
		CreatedAt: user_token.CreatedAt,

	}

	if err := c.RefreshTokenRepository.Create(tx, user_token); err != nil {
		c.Log.Warnf("Failed to create user token: %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return login_response, nil
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

	if err := c.RefreshTokenRepository.FindByToken(tx, user_token); err != nil {
		c.Log.Warnf("Failed find user by token : %+v", err)
		return nil, fiber.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return &model.Auth{ID: user_token.UserID}, nil
}

func (u *UserUseCase) HasPermission(ctx context.Context, user_id uuid.UUID, permissionName string) (bool, error) {
	var count int64
	err := u.DB.WithContext(ctx).
		Table("users").
		Joins("JOIN user_roles ur ON ur.user_id = users.id").
		Joins("JOIN role_permissions rp ON rp.role_id = ur.role_id").
		Joins("JOIN permissions p ON p.id = rp.permission_id").
		Where("users.id = ? AND p.name = ?", user_id, permissionName).
		Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

package repository

import (
	"go-absen-be/internal/entity"
	"go-absen-be/internal/model"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserRepository struct {
	Repository[entity.User]
	Log *logrus.Logger
}

func NewUserRepository(log *logrus.Logger) *UserRepository {
	return &UserRepository{
		Log: log,
	}
}

func (r *UserRepository) FindByToken(db *gorm.DB, user *entity.User, token string) error {
	return db.Where("token = ? AND deleted_at IS NULL", token).First(user).Error
}

func (r *UserRepository) FindByUsername(db *gorm.DB, user *entity.User, username string) error {
	return db.Where("username = ? AND deleted_at IS NULL", username).First(user).Error
}

func (r *UserRepository) Search(db *gorm.DB, request *model.SearchRequest) ([]entity.User, int64, error) {
	var user []entity.User
	if err := db.Scopes(r.FilterUser(request)).
	Select("users.*").
	Where("deleted_at IS NULL").
	Offset((request.Page - 1) * request.Size).
	Limit(request.Size).
	Find(&user).Error; err != nil {
	return nil, 0, err
	}

	var total int64 = 0
	if err := db.Model(&entity.User{}).Scopes(r.FilterUser(request)).Where("deleted_at IS NULL").Count(&total).Error; err != nil {
		return nil, 0, err
	}

	return user, total, nil
}

func (r *UserRepository) FilterUser(request *model.SearchRequest) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		return tx
	}
}
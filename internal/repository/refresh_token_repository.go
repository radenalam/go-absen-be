package repository

import (
	"go-absen-be/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type RefreshTokenRepository struct {
	Repository[entity.RefreshToken]
	Log *logrus.Logger
}

func NewRefreshTokenRepository(log *logrus.Logger) *RefreshTokenRepository {
	return &RefreshTokenRepository{
		Log: log,
	}
}

func (r *RefreshTokenRepository) FindByToken(db *gorm.DB, userToken *entity.RefreshToken) error {
	return db.Where("token = ? AND deleted_at IS NULL", userToken.Token).First(userToken).Error
}
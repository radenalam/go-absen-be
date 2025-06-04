package repository

import (
	"go-absen-be/internal/entity"

	"github.com/sirupsen/logrus"
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

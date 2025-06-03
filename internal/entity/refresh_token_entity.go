package entity

import (
	"time"

	"github.com/google/uuid"
)

type RefreshToken struct {
    BaseModelUUID
    UserID    uuid.UUID `gorm:"type:uuid;not null;index"`
    Token     string    `gorm:"not null;uniqueIndex"`
    ExpiresAt time.Time `gorm:"not null"`

    User      User      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}
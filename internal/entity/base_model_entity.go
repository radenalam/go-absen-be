package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModelUUID struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
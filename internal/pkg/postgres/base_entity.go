package postgres

import (
	"time"

	"github.com/google/uuid"
)

type BaseEntity struct {
	Id        uuid.UUID `gorm:"type:uuid;primaryKey"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	Version   int       `gorm:"version"`
}

package postgres

import (
	"github.com/bilgehanaygn/urun/internal/common/postgres"

	"github.com/google/uuid"
)

type CatalogDbEntity struct {
	postgres.BaseEntity
	Name     string
	Products []uuid.UUID `gorm:"type:uuid[]"`
}

func (CatalogDbEntity) TableName() string {
	return "catalog"
} 
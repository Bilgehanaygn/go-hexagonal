package postgres

import (
	"github.com/bilgehanaygn/urun/internal/pkg/postgres"
)

type CatalogDbEntity struct {
	postgres.BaseEntity
	Name     string
	CatalogProducts []CatalogProductDbEntity `gorm:"foreignKey:CatalogId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (CatalogDbEntity) TableName() string {
	return "catalog"
} 
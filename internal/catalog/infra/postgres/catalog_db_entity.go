package postgres

import (
	"github.com/bilgehanaygn/urun/internal/common/postgres"
)

type CatalogDbEntity struct {
	postgres.BaseEntity
	Name     string
	CatalogProducts []CatalogProductDbEntity
}

func (CatalogDbEntity) TableName() string {
	return "catalog"
} 
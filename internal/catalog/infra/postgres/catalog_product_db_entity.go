package postgres

import (
	"github.com/bilgehanaygn/urun/internal/pkg/postgres"
	"github.com/google/uuid"
)

type CatalogProductDbEntity struct {
	postgres.BaseEntity
	CatalogId uuid.UUID
	ProductId uuid.UUID
	Price     float64
}

func (CatalogProductDbEntity) TableName() string {
	return "catalog_product"
} 
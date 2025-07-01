package postgres

import (
	"github.com/bilgehanaygn/urun/internal/common/postgres"
	productpostgres "github.com/bilgehanaygn/urun/internal/product/infra/postgres"
	"github.com/google/uuid"
)

type CatalogProductDbEntity struct {
	postgres.BaseEntity
	CatalogId uuid.UUID
	ProductId uuid.UUID
	Catalog CatalogDbEntity
	Product productpostgres.ProductDbEntity
	Price     float64
}

func (CatalogProductDbEntity) TableName() string {
	return "catalog_product"
} 
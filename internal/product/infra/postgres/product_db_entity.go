package postgres

import "github.com/bilgehanaygn/urun/internal/pkg/postgres"

type ProductDbEntity struct {
	postgres.BaseEntity
	Name  string
	Price float64
}

func (ProductDbEntity) TableName() string {
	return "product"
}

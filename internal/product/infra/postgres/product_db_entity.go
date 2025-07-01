package postgres

import (
	"urun/internal/product/domain"
	"urun/internal/common/postgres"
)

type ProductDbEntity struct {
	postgres.BaseEntity
	Name  string
	Price float64
}

func (ProductDbEntity) TableName() string {
	return "product"
} 
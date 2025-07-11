package postgres

import (
	"github.com/bilgehanaygn/urun/internal/pkg/domain"
	"github.com/bilgehanaygn/urun/internal/pkg/postgres"
)

type ProductDbEntity struct {
	postgres.BaseEntity
	Name  string
	Price float64
	Status domain.ActivenessStatus
}

func (ProductDbEntity) TableName() string {
	return "product"
}

package domain

import (
	"github.com/google/uuid"
)

type CatalogProduct struct {
	Id        uuid.UUID
	CatalogId uuid.UUID
	ProductId uuid.UUID
	Price     float64
}

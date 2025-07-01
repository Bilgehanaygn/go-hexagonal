package response

import (
	"github.com/google/uuid"
)

type CatalogDetailDto struct {
	Id       uuid.UUID   `json:"id"`
	Name     string      `json:"name"`
	CatalogProducts []CatalogProductDto `json:"catalogProducts"`
}

type CatalogProductDto struct {
	Id uuid.UUID `json:"id"`
	CatalogId uuid.UUID `json:"catalogId"`
	ProductId uuid.UUID `json:"productId"`
	Price float64 `json:"price"`
}

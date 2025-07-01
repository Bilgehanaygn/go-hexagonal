package domain

import (
	"github.com/google/uuid"
)

type Catalog struct {
	Id       uuid.UUID
	Name     string
	CatalogProducts []CatalogProduct
} 
package domain

import (
	productdomain "github.com/bilgehanaygn/urun/internal/product/domain"
	"github.com/google/uuid"
)

type Catalog struct {
	Id       uuid.UUID
	Name     string
	Products []productdomain.Product
} 
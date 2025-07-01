package response

import (
	"github.com/bilgehanaygn/urun/internal/catalog/domain"
	"github.com/google/uuid"
)

type CatalogDetailDto struct {
	Id       uuid.UUID   `json:"id"`
	Name     string      `json:"name"`
	Products []uuid.UUID `json:"products"`
}

func NewCatalogDetailDto(catalog *domain.Catalog) CatalogDetailDto {
	return CatalogDetailDto{
		Id:       catalog.Id,
		Name:     catalog.Name,
		// Products: should be mapped from catalog.Products to UUIDs
	}
} 
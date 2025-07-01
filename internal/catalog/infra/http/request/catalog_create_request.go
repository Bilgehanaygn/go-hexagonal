package request

import (
	"github.com/bilgehanaygn/urun/internal/catalog/domain"
	"github.com/google/uuid"
)

type CatalogCreateRequest struct {
	Name     string      `json:"name"`
	Products []uuid.UUID `json:"products"`
}

func (request *CatalogCreateRequest) ToDomainEntity() (*domain.Catalog, error) {
	return &domain.Catalog{
		Id:       uuid.New(),
		Name:     request.Name,
		// Products: will be resolved from UUIDs elsewhere
	}, nil
} 
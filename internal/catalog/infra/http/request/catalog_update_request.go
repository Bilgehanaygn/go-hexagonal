package request

import (
	"github.com/bilgehanaygn/urun/internal/catalog/domain"
	"github.com/google/uuid"
)

type CatalogUpdateRequest struct {
	Id       uuid.UUID   `json:"id"`
	Name     string      `json:"name"`
	Products []uuid.UUID `json:"products"`
}

func (request *CatalogUpdateRequest) ToDomainEntity() (*domain.Catalog, error) {
	return &domain.Catalog{
		Id:   request.Id,
		Name: request.Name,
		// Products: will be resolved from UUIDs elsewhere
	}, nil
}

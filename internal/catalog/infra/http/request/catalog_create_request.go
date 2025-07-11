package request

import (
	"github.com/bilgehanaygn/urun/internal/catalog/domain"
	"github.com/google/uuid"
)

type CatalogCreateRequest struct {
	Name            string                        `json:"name"`
	CatalogProducts []CatalogProductCreateRequest `json:"catalogProducts"`
}

type CatalogProductCreateRequest struct {
	ProductId uuid.UUID `json:"productId"`
	Price     float64   `json:"price"`
}

func (request *CatalogCreateRequest) ToDomainEntity() (*domain.Catalog, error) {
	catalogId := uuid.New()
	cps := make([]domain.CatalogProduct, len(request.CatalogProducts))

	for i, cp := range request.CatalogProducts {
		cps[i] = domain.CatalogProduct{
			Id:        uuid.New(),
			CatalogId: catalogId,
			ProductId: cp.ProductId,
			Price:     cp.Price,
		}
	}

	return &domain.Catalog{
		Id:              catalogId,
		Name:            request.Name,
		CatalogProducts: cps,
	}, nil
}

package request

import (
	"github.com/bilgehanaygn/urun/internal/catalog/domain"
	"github.com/google/uuid"
)

type CatalogCreateRequest struct {
	Name     string      `json:"name"`
	CatalogProducts []CatalogProductCreateRequest `json:"catalogProducts"`
}

type CatalogProductCreateRequest struct {
	ProductId uuid.UUID `json:"productId"`
	Price float64 `json:"price"`
}

func (request *CatalogCreateRequest) ToDomainEntity() (*domain.Catalog, error) {
	catalogId := uuid.New()
	catalogProducts := make([]*domain.CatalogProduct, len(request.CatalogProducts))

	for i, catalogProduct := range request.CatalogProducts {
		catalogProducts[i] = &domain.CatalogProduct{
			Id: uuid.New(),
			CatalogId: catalogId,
			ProductId: catalogProduct.ProductId,
			Price: catalogProduct.Price,
		}
	}

	return &domain.Catalog{
		Id:       catalogId,
		Name:     request.Name,
		CatalogProducts: catalogProducts,
	}, nil
}


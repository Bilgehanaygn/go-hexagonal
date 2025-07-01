package request

import (
	"urun/internal/product/domain"
	"github.com/google/uuid"
)

type ProductUpdateRequest struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Price float64   `json:"price"`
}

func (request *ProductUpdateRequest) ToDomainEntity() (*domain.Product, error) {
	return &domain.Product{
		BaseEntity: domain.BaseEntity{
			ID: request.ID,
		},
		Name:  request.Name,
		Price: request.Price,
	}, nil
} 
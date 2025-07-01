package request

import (
	"github.com/bilgehanaygn/urun/internal/product/domain"
)

type ProductCreateRequest struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func (request *ProductCreateRequest) ToDomainEntity() (*domain.Product, error) {
	return &domain.Product{
		Name:  request.Name,
		Price: request.Price,
	}, nil
} 
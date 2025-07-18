package request

import (
	pkgdomain "github.com/bilgehanaygn/urun/internal/pkg/domain"
	"github.com/bilgehanaygn/urun/internal/product/domain"
	"github.com/google/uuid"
)

type ProductCreateRequest struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func (request *ProductCreateRequest) ToDomainEntity() (*domain.Product, error) {
	return &domain.Product{
		Id:    uuid.New(),
		Name:  request.Name,
		Price: request.Price,
		Status: pkgdomain.ACTIVE,
	}, nil
}

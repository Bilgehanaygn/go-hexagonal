package response

import (
	"github.com/google/uuid"
	"urun/internal/product/domain"
)

type ProductDetailDto struct {
	Id    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Price float64   `json:"price"`
}

func NewProductDetailDto(product *domain.Product) ProductDetailDto {
	return ProductDetailDto{
		Id:    product.ID,
		Name:  product.Name,
		Price: product.Price,
	}
} 
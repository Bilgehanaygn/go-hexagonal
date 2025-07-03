package postgres

import (
	"github.com/bilgehanaygn/urun/internal/pkg/postgres"
	"github.com/bilgehanaygn/urun/internal/product/domain"
	"github.com/bilgehanaygn/urun/internal/product/infra/http/response"
)

func toDbEntity(product *domain.Product) *ProductDbEntity {
	dbProduct := &ProductDbEntity{
		BaseEntity: postgres.BaseEntity{
			Id: product.Id,
		},
		Name:  product.Name,
		Price: product.Price,
	}
	return dbProduct
}

func toDomainEntity(dbProduct *ProductDbEntity) *domain.Product {
	product := domain.Product{
		Name:      dbProduct.Name,
		Price:     dbProduct.Price,
	}
	return &product
}

func toProductDetailDto(dbProduct *ProductDbEntity) *response.ProductDetailDto {
	dto := response.ProductDetailDto{
		Id:    dbProduct.BaseEntity.Id,
		Name:  dbProduct.Name,
		Price: dbProduct.Price,
	}
	return &dto
} 
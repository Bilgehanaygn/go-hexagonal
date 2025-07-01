package postgres

import (
	"urun/internal/product/domain"
	"urun/internal/product/infra/http/response"
	"urun/internal/common/postgres"
)

func toDbEntity(product *domain.Product) *ProductDbEntity {
	dbProduct := &ProductDbEntity{
		BaseEntity: postgres.BaseEntity{
			ID: product.ID,
		},
		Name:  product.Name,
		Price: product.Price,
	}
	return dbProduct
}

func toDomainEntity(dbProduct *ProductDbEntity) *domain.Product {
	product := domain.Product{
		BaseEntity: dbProduct.BaseEntity,
		Name:      dbProduct.Name,
		Price:     dbProduct.Price,
	}
	return &product
}

func toProductDetailDto(dbProduct *ProductDbEntity) *response.ProductDetailDto {
	dto := response.ProductDetailDto{
		Id:    dbProduct.BaseEntity.ID,
		Name:  dbProduct.Name,
		Price: dbProduct.Price,
	}
	return &dto
} 
package postgres

import (
	"context"

	"github.com/bilgehanaygn/urun/internal/product/application/ports"
	"github.com/bilgehanaygn/urun/internal/product/infra/http/response"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductQueryRepository struct {
	db *gorm.DB
}

func NewProductQueryPort(db *gorm.DB) ports.ProductQueryPort {
	return &ProductQueryRepository{db: db}
}

func (repo *ProductQueryRepository) GetDtoById(ctx context.Context, id uuid.UUID) (*response.ProductDetailDto, error) {
	var dbProduct ProductDbEntity
	if err := repo.db.WithContext(ctx).
		First(&dbProduct, "id = ?", id).
		Error; err != nil {
		return nil, err
	}
	return toProductDetailDto(&dbProduct), nil
}

func (repo *ProductQueryRepository) GetDtoList(ctx context.Context) ([]*response.ProductDetailDto, error) {
	var dbProducts []ProductDbEntity
	if err := repo.db.WithContext(ctx).
		Find(&dbProducts).
		Error; err != nil {
		return nil, err
	}
	dtos := make([]*response.ProductDetailDto, len(dbProducts))
	for i, dbProd := range dbProducts {
		dtos[i] = toProductDetailDto(&dbProd)
	}
	return dtos, nil
}

package postgres

import (
	"context"

	"github.com/bilgehanaygn/urun/internal/category/application/ports"
	"github.com/bilgehanaygn/urun/internal/category/infra/http/response"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CategoryQueryRepository struct {
	db *gorm.DB
}

func NewCategoryQueryPort(db *gorm.DB) ports.CategoryQueryPort {
	return &CategoryQueryRepository{db: db}
}

func (repo *CategoryQueryRepository) GetDtoById(ctx context.Context, id uuid.UUID) (*response.CategoryDetailDto, error) {
	var dbCategory CategoryDbEntity
	if err := repo.db.WithContext(ctx).
		First(&dbCategory, "id = ?", id).
		Error; err != nil {
		return nil, err
	}

	return toCategoryDetailDto(&dbCategory), nil
}

func (repo *CategoryQueryRepository) GetDtoList(ctx context.Context) ([]*response.CategoryDetailDto, error) {
	var dbCategories []CategoryDbEntity
	if err := repo.db.WithContext(ctx).
		Find(&dbCategories).
		Error; err != nil {
		return nil, err
	}

	dtos := make([]*response.CategoryDetailDto, len(dbCategories))
	for i, dbCat := range dbCategories {
		dtos[i] = toCategoryDetailDto(&dbCat)
	}

	return dtos, nil
}

package postgres

import (
	"context"

	"github.com/bilgehanaygn/urun/internal/catalog/application/ports"
	"github.com/bilgehanaygn/urun/internal/catalog/infra/http/response"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CatalogQueryRepository struct {
	db *gorm.DB
}

func NewCatalogQueryPort(db *gorm.DB) ports.CatalogQueryPort {
	return &CatalogQueryRepository{db: db}
}

func (repo *CatalogQueryRepository) GetDtoById(ctx context.Context, id uuid.UUID) (*response.CatalogDetailDto, error) {
	var dbCatalog CatalogDbEntity
	if err := repo.db.WithContext(ctx).
		First(&dbCatalog, "id = ?", id).
		Error; err != nil {
		return nil, err
	}
	return toCatalogDetailDto(&dbCatalog), nil
}

func (repo *CatalogQueryRepository) GetDtoList(ctx context.Context) ([]*response.CatalogDetailDto, error) {
	var dbCatalogs []CatalogDbEntity
	if err := repo.db.WithContext(ctx).
		Find(&dbCatalogs).
		Error; err != nil {
		return nil, err
	}
	dtos := make([]*response.CatalogDetailDto, len(dbCatalogs))
	for i, dbCat := range dbCatalogs {
		dtos[i] = toCatalogDetailDto(&dbCat)
	}
	return dtos, nil
} 
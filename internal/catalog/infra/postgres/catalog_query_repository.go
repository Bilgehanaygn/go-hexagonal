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
		Preload("CatalogProducts").
		First(&dbCatalog, "id = ?", id).
		Error; err != nil {
		return nil, err
	}
	return toCatalogDetailDto(&dbCatalog), nil
}

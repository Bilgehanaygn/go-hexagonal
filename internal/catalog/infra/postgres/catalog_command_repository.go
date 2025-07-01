package postgres

import (
	"context"

	"github.com/bilgehanaygn/urun/internal/catalog/application/ports"
	"github.com/bilgehanaygn/urun/internal/catalog/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CatalogCommandRepository struct {
	db *gorm.DB
}

func NewCatalogCommandPort(db *gorm.DB) ports.CatalogCommandPort {
	return &CatalogCommandRepository{db: db}
}

func (r *CatalogCommandRepository) Create(ctx context.Context, cat *domain.Catalog) (*uuid.UUID, error) {
	dbCat := toDbEntity(cat)

	if err := r.db.WithContext(ctx).Create(dbCat).Error; err != nil {
		return nil, err
	}
	return &dbCat.BaseEntity.ID, nil
}


func (repo *CatalogCommandRepository) Update(ctx context.Context, catalog *domain.Catalog) (*uuid.UUID, error) {
	dbCatalog := toDbEntity(catalog)
	result := repo.db.
		WithContext(ctx).
		Model(&CatalogDbEntity{}).
		Where("id = ?", dbCatalog.BaseEntity.ID).
		Updates(dbCatalog)
	if err := result.Error; err != nil {
		return nil, err
	}
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &dbCatalog.BaseEntity.ID, nil
}

func (repo *CatalogCommandRepository) FindById(ctx context.Context, id uuid.UUID) (*domain.Catalog, error) {
	var dbCatalog CatalogDbEntity
	if err := repo.db.WithContext(ctx).
		First(&dbCatalog, "id = ?", id).
		Error; err != nil {
		return nil, err
	}
	catalog := toDomainEntity(&dbCatalog)
	return catalog, nil
} 
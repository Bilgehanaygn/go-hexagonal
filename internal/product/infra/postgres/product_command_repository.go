package postgres

import (
	"context"

	"github.com/bilgehanaygn/urun/internal/catalog/infra/postgres"
	"github.com/bilgehanaygn/urun/internal/product/application/ports"
	"github.com/bilgehanaygn/urun/internal/product/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductCommandRepository struct {
	db *gorm.DB
}

func NewProductCommandPort(db *gorm.DB) ports.ProductCommandPort {
	return &ProductCommandRepository{db: db}
}

func (repo *ProductCommandRepository) Create(ctx context.Context, product *domain.Product) (*uuid.UUID, error) {
	dbProduct := toDbEntity(product)
	if err := repo.db.Create(dbProduct).Error; err != nil {
		return nil, err
	}
	return &dbProduct.BaseEntity.Id, nil
}

func (repo *ProductCommandRepository) Update(ctx context.Context, product *domain.Product) (*uuid.UUID, error) {
	dbProduct := toDbEntity(product)
	result := repo.db.
		WithContext(ctx).
		Model(&ProductDbEntity{}).
		Where("id = ?", dbProduct.BaseEntity.Id).
		Updates(dbProduct)
	if err := result.Error; err != nil {
		return nil, err
	}
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &dbProduct.BaseEntity.Id, nil
}

func (repo *ProductCommandRepository) FindById(ctx context.Context, id uuid.UUID) (*domain.Product, error) {
	var dbProduct ProductDbEntity
	if err := repo.db.WithContext(ctx).
		First(&dbProduct, "id = ?", id).
		Error; err != nil {
		return nil, err
	}
	product := toDomainEntity(&dbProduct)
	return product, nil
}

func (repo *ProductCommandRepository) IsAssociatedWithAnyCatalog(ctx context.Context, id uuid.UUID) (bool, error){
	var count int64

	result := repo.db.
		WithContext(ctx).
		Model(&postgres.CatalogProductDbEntity{}).
		Where("product_id = ?", id).
		Count(&count)

	if result.Error != nil {
		return false, result.Error
	}
		
    return count > 0, nil
}
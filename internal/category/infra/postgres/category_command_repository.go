package postgres

import (
	"context"

	"github.com/bilgehanaygn/urun/internal/category/application/ports"
	"github.com/bilgehanaygn/urun/internal/category/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CategoryCommandRepository struct {
	db *gorm.DB
}

func NewCategoryCommandPort(db *gorm.DB) ports.CategoryCommandPort {
	return &CategoryCommandRepository{db: db}
}

func (repo *CategoryCommandRepository) Create(ctx context.Context, category *domain.Category) (*uuid.UUID, error) {
	dbCategory := toDbEntity(category)
	if err := repo.db.Create(dbCategory).Error; err != nil {
		return nil, err
	}

	return &dbCategory.Id, nil
}

func (repo *CategoryCommandRepository) Update(ctx context.Context, category *domain.Category) (*uuid.UUID, error) {
	dbCategory := toDbEntity(category)

	result := repo.db.
		WithContext(ctx).
		Model(&CategoryDbEntity{}).
		Where("id = ?", dbCategory.Id).
		Updates(dbCategory)

	if err := result.Error; err != nil {
		return nil, err
	}

	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return &dbCategory.Id, nil
}

func (repo *CategoryCommandRepository) FindById(ctx context.Context, id uuid.UUID) (*domain.Category, error) {
	var dbCategory CategoryDbEntity
	if err := repo.db.WithContext(ctx).
		First(&dbCategory, "id = ?", id).
		Error; err != nil {
		return nil, err
	}

	category := toDomainEntity(&dbCategory)
	
	return category, nil
}
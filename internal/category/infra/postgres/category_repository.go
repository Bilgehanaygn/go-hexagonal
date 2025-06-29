package postgres

import (
	"context"

	"github.com/bilgehanaygn/urun/internal/category/application"
	"github.com/bilgehanaygn/urun/internal/category/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GormCategoryRepository struct {
	db *gorm.DB
}

func NewGormCategoryRepository(db *gorm.DB) application.CategoryRepository {
	return &GormCategoryRepository{db: db}
}

func (repo *GormCategoryRepository) Create(ctx context.Context, category *domain.Category) (*domain.Category, error) {
	dbCategory := toDbEntity(category)
	if err := repo.db.Create(dbCategory).Error; err != nil {
		return nil, err
	}

	return repo.FindById(ctx, dbCategory.ID)
}

func (repo *GormCategoryRepository) Update(ctx context.Context, category *domain.Category) (*domain.Category, error) {
	dbCategory := toDbEntity(category)

	result := repo.db.
		WithContext(ctx).
		Model(&CategoryDbEntity{}).
		Where("id = ?", dbCategory.ID).
		Updates(dbCategory)

	if err := result.Error; err != nil {
		return nil, err
	}

	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return repo.FindById(ctx, dbCategory.ID)
}

func (repo *GormCategoryRepository) FindById(ctx context.Context, id uuid.UUID) (*domain.Category, error) {
	var dbCategory CategoryDbEntity
	if err := repo.db.WithContext(ctx).
		First(&dbCategory, "id = ?", id).
		Error; err != nil {
		return nil, err
	}

	return toDomainEntity(&dbCategory), nil
}

func (repo *GormCategoryRepository) List(ctx context.Context) ([]*domain.Category, error) {
	var dbCategories []CategoryDbEntity
	if err := repo.db.WithContext(ctx).
		Find(&dbCategories).
		Error; err != nil {
		return nil, err
	}

	categories := make([]*domain.Category, len(dbCategories))
	for i, dbCat := range dbCategories {
		categories[i] = toDomainEntity(&dbCat)
	}

	return categories, nil
}

package db

import (
	"context"

	"github.com/bilgehanaygn/urun/internal/category/application"
	"github.com/bilgehanaygn/urun/internal/category/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MockCategoryRepository struct {
	db *gorm.DB
}

func NewMockCategoryRepository(db *gorm.DB) application.CategoryRepository {
	return &MockCategoryRepository{db: db}
}

func (repo *MockCategoryRepository) Create(ctx context.Context, category *domain.Category) (*domain.Category, error) {
	return category, nil
}

func (repo *MockCategoryRepository) Update(ctx context.Context, category *domain.Category) (*domain.Category, error) {
	return category, nil
}

func (repo *MockCategoryRepository) FindById(ctx context.Context, id uuid.UUID) (*domain.Category, error) {
	return &domain.Category{
		Id: uuid.New(),
		Name: "Smth",
		Kind: domain.MAIN_CATEGORY,
		ParentCategoryId: uuid.New(),
		Status: domain.ACTIVE,
	}, nil
}

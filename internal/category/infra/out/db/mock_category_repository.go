package db

import (
	"context"
	"fmt"

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

func (repo *MockCategoryRepository) List(ctx context.Context) ([]*domain.Category, error) {
	categories := make([]*domain.Category, 10)

	for i := 0; i < 10; i++ {
		categories[i] = &domain.Category{
			Id:               uuid.New(),
			Name:             fmt.Sprintf("Category %d", i+1),
			Kind:             domain.MAIN_CATEGORY,
			ParentCategoryId: uuid.Nil,                   
			Status:           domain.ACTIVE,
		}
	}

	return categories, nil
}
package repository

import (
	"github.com/bilgehanaygn/urun/internal/category/application"
	"github.com/bilgehanaygn/urun/internal/category/domain"
)

type GormCategoryRepository struct {
	db *gorm.DB
}

func NewGormCategoryRepository(db *gorm.DB) application.CategoryRepository {
	return &GormCategoryRepository{db: db}
}

func (repo *GormCategoryRepository) Create(category *domain.Category) (*domain.Category, error){
	dbCategory := toDbCategory(category)
}
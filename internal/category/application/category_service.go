package application

import (
	"context"

	"github.com/bilgehanaygn/urun/internal/category/domain"
)

type CategoryService struct {
	CategoryRepository CategoryRepository
}

func NewCategoryService(repository CategoryRepository) *CategoryService {
	return &CategoryService{CategoryRepository: repository}
}


func (categoryService *CategoryService) HandleCreate(ctx context.Context, category domain.Category) error {
	_, err := categoryService.CategoryRepository.Create(ctx, &category)

	if err != nil {
		return err
	}

	return nil
}

func (categoryService *CategoryService) HandleUpdate(ctx context.Context, category domain.Category) error {
	_, err := categoryService.CategoryRepository.Update(ctx, &category)

	if err != nil {
		return err
	}

	return nil
}
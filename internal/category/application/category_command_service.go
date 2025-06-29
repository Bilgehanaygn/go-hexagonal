package application

import (
	"context"

	"github.com/bilgehanaygn/urun/internal/category/domain"
)

type CategoryCommandService struct {
	CategoryRepository CategoryRepository
}

func NewCategoryCommandService(repository CategoryRepository) *CategoryCommandService {
	return &CategoryCommandService{CategoryRepository: repository}
}


func (categoryCommandService *CategoryCommandService) HandleCreate(ctx context.Context, category domain.Category) error {
	_, err := categoryCommandService.CategoryRepository.Create(ctx, &category)

	if err != nil {
		return err
	}

	return nil
}
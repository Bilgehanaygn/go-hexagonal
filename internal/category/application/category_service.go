package application

import (
	"context"

	"github.com/bilgehanaygn/urun/internal/category/domain"
	"github.com/google/uuid"
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

func (s *CategoryService) HandleGetById(ctx context.Context, id uuid.UUID) (*domain.Category, error) {
	return s.CategoryRepository.FindById(ctx, id)
}

func (s *CategoryService) HandleList(ctx context.Context) ([]*domain.Category, error) {
	return s.CategoryRepository.List(ctx)
}

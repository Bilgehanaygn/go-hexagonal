package application

import (
	"context"

	"github.com/bilgehanaygn/urun/internal/category/domain"
	"github.com/google/uuid"
)

type CategoryRepository interface {
	FindById(ctx context.Context, id uuid.UUID) (*domain.Category, error)
	List(ctx context.Context) ([]*domain.Category, error)
	Create(ctx context.Context, s *domain.Category) (*domain.Category, error)
	Update(ctx context.Context, s *domain.Category) (*domain.Category, error)
}

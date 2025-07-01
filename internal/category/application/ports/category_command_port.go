package ports

import (
	"context"

	"github.com/bilgehanaygn/urun/internal/category/domain"
	"github.com/google/uuid"
)

type CategoryCommandPort interface {
	Create(ctx context.Context, s *domain.Category) (*uuid.UUID, error)
	Update(ctx context.Context, s *domain.Category) (*uuid.UUID, error)
	FindById(ctx context.Context, id uuid.UUID) (*domain.Category, error)
}

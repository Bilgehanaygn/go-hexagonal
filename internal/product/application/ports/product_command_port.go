package ports

import (
	"context"

	"github.com/bilgehanaygn/urun/internal/product/domain"
	"github.com/google/uuid"
)

type ProductCommandPort interface {
	Create(ctx context.Context, s *domain.Product) (*uuid.UUID, error)
	Update(ctx context.Context, s *domain.Product) (*uuid.UUID, error)
	FindById(ctx context.Context, id uuid.UUID) (*domain.Product, error)
}

package ports

import (
	"context"

	"github.com/bilgehanaygn/urun/internal/catalog/domain"

	"github.com/google/uuid"
)

type CatalogCommandPort interface {
	Create(ctx context.Context, s *domain.Catalog) (*uuid.UUID, error)
	Update(ctx context.Context, s *domain.Catalog) (*uuid.UUID, error)
	FindById(ctx context.Context, id uuid.UUID) (*domain.Catalog, error)
}

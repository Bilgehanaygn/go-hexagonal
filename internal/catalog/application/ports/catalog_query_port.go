package ports

import (
	"context"

	"github.com/bilgehanaygn/urun/internal/catalog/infra/http/response"
	"github.com/google/uuid"
)

type CatalogQueryPort interface {
	GetDtoById(ctx context.Context, id uuid.UUID) (*response.CatalogDetailDto, error)
} 
package ports

import (
	"context"

	"github.com/bilgehanaygn/urun/internal/product/infra/http/response"
	"github.com/google/uuid"
)

type ProductQueryPort interface {
	GetDtoById(ctx context.Context, id uuid.UUID) (*response.ProductDetailDto, error)
} 
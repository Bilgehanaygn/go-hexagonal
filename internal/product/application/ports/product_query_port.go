package ports

import (
	"context"
	"urun/internal/product/infra/http/response"
	"github.com/google/uuid"
)

type ProductQueryPort interface {
	GetDtoById(ctx context.Context, id uuid.UUID) (*response.ProductDetailDto, error)
	GetDtoList(ctx context.Context) ([]*response.ProductDetailDto, error)
} 
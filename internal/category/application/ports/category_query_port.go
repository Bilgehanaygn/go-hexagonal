package ports

import (
	"context"

	"github.com/bilgehanaygn/urun/internal/category/infra/http/response"
	"github.com/google/uuid"
)

type CategoryQueryPort interface {
	GetDtoById(ctx context.Context, id uuid.UUID) (*response.CategoryDetailDto, error)
	GetDtoList(ctx context.Context) ([]*response.CategoryDetailDto, error)
}



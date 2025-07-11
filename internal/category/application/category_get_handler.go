package application

import (
	"context"

	"github.com/bilgehanaygn/urun/internal/category/application/ports"
	"github.com/bilgehanaygn/urun/internal/category/infra/http/response"
	"github.com/google/uuid"
)

type CategoryQueryHandler struct {
	CategoryQPort ports.CategoryQueryPort
}

func (categoryQHandler *CategoryQueryHandler) Handle(ctx context.Context, id *uuid.UUID) (*response.CategoryDetailDto, error) {

	category, err := categoryQHandler.CategoryQPort.GetDtoById(ctx, *id)
	if err != nil {
		return nil, err
	}

	return category, nil
}

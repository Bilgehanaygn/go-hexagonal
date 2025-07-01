package application

import (
	"context"

	"github.com/bilgehanaygn/urun/internal/product/application/ports"
	"github.com/bilgehanaygn/urun/internal/product/infra/http/response"
	"github.com/google/uuid"
)

type ProductGetHandler struct {
	ProductQPort ports.ProductQueryPort
}

func (productQHandler *ProductGetHandler) Handle(ctx context.Context, id *uuid.UUID) (*response.ProductDetailDto, error) {
	product, err := productQHandler.ProductQPort.GetDtoById(ctx, *id)
	if err != nil {
		return nil, err
	}
	return product, nil
} 
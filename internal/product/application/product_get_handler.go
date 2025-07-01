package application

import (
	"context"

	"urun/internal/product/application/ports"
	"urun/internal/product/infra/http/response"
	"github.com/google/uuid"
)

type ProductQueryHandler struct {
	ProductQPort ports.ProductQueryPort
}

func (productQHandler *ProductQueryHandler) Handle(ctx context.Context, id *uuid.UUID) (*response.ProductDetailDto, error) {
	product, err := productQHandler.ProductQPort.GetDtoById(ctx, *id)
	if err != nil {
		return nil, err
	}
	return product, nil
} 
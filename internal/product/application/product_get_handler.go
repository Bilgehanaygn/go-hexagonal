package application

import (
	"context"

	"github.com/bilgehanaygn/urun/internal/product/application/ports"
	"github.com/bilgehanaygn/urun/internal/product/infra/http/request"
	"github.com/bilgehanaygn/urun/internal/product/infra/http/response"
)

type ProductGetHandler struct {
	ProductQPort ports.ProductQueryPort
}

func (productQHandler *ProductGetHandler) Handle(ctx context.Context, req *request.ProductGetRequest) (*response.ProductDetailDto, error) {
	product, err := productQHandler.ProductQPort.GetDtoById(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return product, nil
} 
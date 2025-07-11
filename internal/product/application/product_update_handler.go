package application

import (
	"context"

	"github.com/bilgehanaygn/urun/internal/product/application/ports"
	"github.com/bilgehanaygn/urun/internal/product/infra/http/request"
	"github.com/bilgehanaygn/urun/internal/product/infra/http/response"
)

type ProductUpdateHandler struct {
	ProductCPort ports.ProductCommandPort
}

func (productUpdateHandler *ProductUpdateHandler) Handle(ctx context.Context, req *request.ProductUpdateRequest) (*response.ProductUpdateResponse, error) {
	product, err := req.ToDomainEntity()
	if err != nil {
		return nil, err
	}

	_, err = productUpdateHandler.ProductCPort.Update(ctx, product)
	if err != nil {
		return nil, err
	}

	return &response.ProductUpdateResponse{Id: product.Id}, nil
}

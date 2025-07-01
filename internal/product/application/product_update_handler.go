package application

import (
	"context"

	"urun/internal/product/application/ports"
	"urun/internal/product/infra/http/request"
	"urun/internal/product/infra/http/response"
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

	return &response.ProductUpdateResponse{Id: product.ID}, nil
} 
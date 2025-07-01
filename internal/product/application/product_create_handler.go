package application

import (
	"context"

	"urun/internal/product/application/ports"
	"urun/internal/product/infra/http/request"
	"urun/internal/product/infra/http/response"
)

type ProductCreateHandler struct {
	ProductCPort ports.ProductCommandPort
}

func (productCreateHandler *ProductCreateHandler) Handle(ctx context.Context, req *request.ProductCreateRequest) (*response.ProductCreateResponse, error) {
	product, err := req.ToDomainEntity()
	if err != nil {
		return nil, err
	}

	_, err = productCreateHandler.ProductCPort.Create(ctx, product)
	if err != nil {
		return nil, err
	}

	return &response.ProductCreateResponse{Id: product.ID}, nil
} 
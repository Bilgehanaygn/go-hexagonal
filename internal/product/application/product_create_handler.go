package application

import (
	"context"

	"github.com/bilgehanaygn/urun/internal/product/application/ports"
	"github.com/bilgehanaygn/urun/internal/product/infra/http/request"
	"github.com/bilgehanaygn/urun/internal/product/infra/http/response"
)

type ProductCreateHandler struct {
	ProductCPort ports.ProductCommandPort
}

func (productCreateHandler *ProductCreateHandler) Handle(ctx context.Context, req *request.ProductCreateRequest) (*response.ProductCreateResponse, error) {
	product, err := req.ToDomainEntity()
	if err != nil {
		return nil, err
	}

	id, err := productCreateHandler.ProductCPort.Create(ctx, product)
	if err != nil {
		return nil, err
	}

	return &response.ProductCreateResponse{Id: *id}, nil
} 
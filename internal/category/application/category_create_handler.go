package application

import (
	"context"

	"github.com/bilgehanaygn/urun/internal/category/application/ports"
	"github.com/bilgehanaygn/urun/internal/category/infra/http/request"
	"github.com/bilgehanaygn/urun/internal/category/infra/http/response"
)

type CategoryCreateHandler struct {
	CategoryCPort ports.CategoryCommandPort
}

func (categoryCreateHandler *CategoryCreateHandler) Handle(ctx context.Context, req *request.CategoryCreateRequest) (*response.CategoryCreateResponse, error) {
	category, err := req.ToDomainEntity()
	if err != nil {
		return nil, err
	}

	_, err = categoryCreateHandler.CategoryCPort.Create(ctx, category)

	if err != nil {
		return nil, err
	}

	return &response.CategoryCreateResponse{Id: category.Id}, nil
}

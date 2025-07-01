package application

import (
	"context"

	"github.com/bilgehanaygn/urun/internal/category/application/ports"
	"github.com/bilgehanaygn/urun/internal/category/infra/http/request"
	"github.com/bilgehanaygn/urun/internal/category/infra/http/response"
)

type CategoryUpdateHandler struct {
	CategoryCPort ports.CategoryCommandPort
}

func (categoryUpdateHandler *CategoryUpdateHandler) Handle(ctx context.Context, req *request.CategoryUpdateRequest) (*response.CategoryUpdateResponse, error) {
	category, err := req.ToDomainEntity()
	if err != nil {
		return nil, err
	} 

	_, err = categoryUpdateHandler.CategoryCPort.Update(ctx, category)

	if err != nil {
		return nil, err
	}

	return &response.CategoryUpdateResponse{Id: category.Id}, nil
}
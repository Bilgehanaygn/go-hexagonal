package application

import (
	"context"

	"github.com/bilgehanaygn/urun/internal/catalog/application/ports"
	"github.com/bilgehanaygn/urun/internal/catalog/infra/http/request"
	"github.com/bilgehanaygn/urun/internal/catalog/infra/http/response"
)

type CatalogUpdateHandler struct {
	CatalogCPort ports.CatalogCommandPort
}

func (catalogUpdateHandler *CatalogUpdateHandler) Handle(ctx context.Context, req *request.CatalogUpdateRequest) (*response.CatalogUpdateResponse, error) {
	catalog, err := req.ToDomainEntity()
	if err != nil {
		return nil, err
	}

	_, err = catalogUpdateHandler.CatalogCPort.Update(ctx, catalog)
	if err != nil {
		return nil, err
	}

	return &response.CatalogUpdateResponse{Id: catalog.Id}, nil
} 
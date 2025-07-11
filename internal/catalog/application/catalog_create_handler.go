package application

import (
	"context"

	"github.com/bilgehanaygn/urun/internal/catalog/application/ports"
	"github.com/bilgehanaygn/urun/internal/catalog/infra/http/request"
	"github.com/bilgehanaygn/urun/internal/catalog/infra/http/response"
)

type CatalogCreateHandler struct {
	CatalogCPort ports.CatalogCommandPort
}

func (catalogCreateHandler *CatalogCreateHandler) Handle(ctx context.Context, req *request.CatalogCreateRequest) (*response.CatalogCreateResponse, error) {
	catalog, err := req.ToDomainEntity()
	if err != nil {
		return nil, err
	}

	_, err = catalogCreateHandler.CatalogCPort.Create(ctx, catalog)
	if err != nil {
		return nil, err
	}

	return &response.CatalogCreateResponse{Id: catalog.Id}, nil
}

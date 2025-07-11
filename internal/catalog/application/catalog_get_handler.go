package application

import (
	"context"

	"github.com/bilgehanaygn/urun/internal/catalog/application/ports"
	"github.com/bilgehanaygn/urun/internal/catalog/infra/http/request"
	"github.com/bilgehanaygn/urun/internal/catalog/infra/http/response"
)

type CatalogGetHandler struct {
	CatalogQPort ports.CatalogQueryPort
}

func (catalogQHandler *CatalogGetHandler) Handle(ctx context.Context, req *request.CatalogGetRequest) (*response.CatalogDetailDto, error) {
	catalog, err := catalogQHandler.CatalogQPort.GetDtoById(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return catalog, nil
}

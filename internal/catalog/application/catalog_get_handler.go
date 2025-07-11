package application

import (
	"context"
	"fmt"

	"github.com/bilgehanaygn/urun/internal/catalog/application/ports"
	"github.com/bilgehanaygn/urun/internal/catalog/infra/http/request"
	"github.com/bilgehanaygn/urun/internal/catalog/infra/http/response"
)

type CatalogGetHandler struct {
	CatalogQPort ports.CatalogQueryPort
}

func (catalogQHandler *CatalogGetHandler) Handle(ctx context.Context, req *request.CatalogGetRequest) (*response.CatalogDetailDto, error) {
	fmt.Println("ID is:", req.Id)
	catalog, err := catalogQHandler.CatalogQPort.GetDtoById(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return catalog, nil
}

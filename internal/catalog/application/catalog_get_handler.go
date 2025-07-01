package application

import (
	"context"

	"github.com/bilgehanaygn/urun/internal/catalog/application/ports"
	"github.com/bilgehanaygn/urun/internal/catalog/infra/http/response"

	"github.com/google/uuid"
)

type CatalogGetHandler struct {
	CatalogQPort ports.CatalogQueryPort
}

func (catalogQHandler *CatalogGetHandler) Handle(ctx context.Context, id *uuid.UUID) (*response.CatalogDetailDto, error) {
	catalog, err := catalogQHandler.CatalogQPort.GetDtoById(ctx, *id)
	if err != nil {
		return nil, err
	}
	return catalog, nil
} 
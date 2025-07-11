package application

import (
	"context"

	"github.com/bilgehanaygn/urun/internal/pkg/domain"
	"github.com/bilgehanaygn/urun/internal/product/application/ports"
	"github.com/bilgehanaygn/urun/internal/product/infra/http/request"
	"github.com/bilgehanaygn/urun/internal/product/infra/http/response"
)

type ProductUpdateStatusHandler struct {
	ProductCPort ports.ProductCommandPort
}

func (productUpdateStatusHandler *ProductUpdateStatusHandler) Handle(ctx context.Context, req *request.ProductUpdateStatusRequest) (*response.ProductUpdateStatusResponse, error) {
	product, err := productUpdateStatusHandler.ProductCPort.FindById(ctx, req.ProductId)
	
	if err != nil {
		return nil, err
	}

	if req.Status == domain.PASSIVE {
		isAssociatedWithCatalog, err := productUpdateStatusHandler.ProductCPort.IsAssociatedWithAnyCatalog(ctx, req.ProductId)

		if(err != nil){
			return nil, err
		}

		if isAssociatedWithCatalog {
			return &response.ProductUpdateStatusResponse{Result: "Product is associated with at least one catalog, cannot be updated as Passive."}, nil
		}
	}

	product.UpdateStatus(req.Status)

	_, err = productUpdateStatusHandler.ProductCPort.Update(ctx, product)
	if err != nil {
		return nil, err
	}

	return &response.ProductUpdateStatusResponse{Result: "succeed"}, nil
}

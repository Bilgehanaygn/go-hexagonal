package request

import (
	pkgdomain "github.com/bilgehanaygn/urun/internal/pkg/domain"
	"github.com/google/uuid"
)

type ProductUpdateStatusRequest struct {
	ProductId    uuid.UUID `param:"productId"`
	Status pkgdomain.ActivenessStatus `json:"status"`
}

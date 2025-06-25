package entity

import (
	"github.com/bilgehanaygn/urun/internal/category/domain"
	"github.com/bilgehanaygn/urun/internal/common/infra/out/db"
	"github.com/google/uuid"
)

type CategoryDbEntity struct {
    db.BaseEntity
    Name             string                
    Kind             domain.CategoryKind   
    Status           domain.ActivenessStatus
    ParentCategoryId uuid.UUID
}

package domain

import (
	"github.com/bilgehanaygn/urun/internal/pkg/domain"
	"github.com/google/uuid"
)

type Category struct {
	Id               uuid.UUID
	Name             string
	Kind             CategoryKind
	ParentCategoryId *uuid.UUID
	Status           domain.ActivenessStatus
}

package domain

import (
	"github.com/bilgehanaygn/urun/internal/pkg/domain"
	"github.com/google/uuid"
)

type Product struct {
	Id    uuid.UUID
	Name  string
	Price float64
	Status domain.ActivenessStatus
}

func (p *Product) UpdateStatus(s domain.ActivenessStatus){
	p.Status = s
}
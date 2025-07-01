package domain

import (
	"urun/internal/common/postgres"
)

type Product struct {
	Name  string
	Price float64
}
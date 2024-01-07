package data

import (
	"clean-arch-api/features/product"
	"clean-arch-api/features/user/data"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string
	Description string
	UserID      int
	User        data.User `gorm:"foreignKey:UserID"`
}

func CoreToModel(input product.ProCore) Product {
	return Product{
		Name:        input.Name,
		Description: input.Description,
		UserID:      input.UserID,
	}
}

func (p Product) ModelToCore() product.ProCore {
	return product.ProCore{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
	}
}

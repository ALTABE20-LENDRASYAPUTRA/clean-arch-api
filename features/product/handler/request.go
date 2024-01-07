package handler

import (
	"clean-arch-api/features/product"
)

type ProductRequest struct {
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
}

func RequestToCore(input ProductRequest) product.ProCore {
	return product.ProCore{
		Name:        input.Name,
		Description: input.Description,
	}
}

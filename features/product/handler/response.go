package handler

import (
	"clean-arch-api/features/product"
	"clean-arch-api/features/user/handler"
	"time"
)

type ProductResponse struct {
	ID          uint                 `json:"id" form:"id"`
	Name        string               `json:"name" form:"name"`
	Description string               `json:"description" form:"description"`
	CreatedAt   time.Time            `json:"created_at" form:"created_at"`
	UpdatedAt   time.Time            `json:"updated_at" form:"updated_at"`
	User        handler.UserProResponse `json:"user"`
}

type ProductResponseUser struct {
	ID          uint                 `json:"id" form:"id"`
	Name        string               `json:"name" form:"name"`
	Description string               `json:"description" form:"description"`
	CreatedAt   time.Time            `json:"created_at" form:"created_at"`
	UpdatedAt   time.Time            `json:"updated_at" form:"updated_at"`
}

func CoreToResponse(data product.ProCore) ProductResponse {
	return ProductResponse{
		ID:          data.ID,
		Name:        data.Name,
		Description: data.Description,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
		User: handler.UserProResponse {
			ID:   data.User.ID,
			Name: data.User.Name,
			Role: data.User.Role,
		},
	}
}

func CoreToResponseList(data []product.ProCore) []ProductResponse {
	var results []ProductResponse
	for _, v := range data {
		results = append(results, CoreToResponse(v))
	}
	return results
}

func CoreToResponseUser(data product.ProCore) ProductResponseUser {
	return ProductResponseUser{
		ID:          data.ID,
		Name:        data.Name,
		Description: data.Description,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
	}
}

func CoreToResponseUserList(data []product.ProCore) []ProductResponseUser {
	var results []ProductResponseUser
	for _, v := range data {
		results = append(results, CoreToResponseUser(v))
	}
	return results
}
package service

import (
	"clean-arch-api/features/product"
)

type productService struct {
	productData product.ProductDataInterface
}

// dependency injection
func New(repo product.ProductDataInterface) product.ProductServiceInterface {
	return &productService{
		productData: repo,
	}
}

// Create implements product.ProductServiceInterface.
func (ps *productService) Create(input product.ProCore) error {
	err := ps.productData.Insert(input.UserID, input)
	if err != nil {
		return err
	}
	return nil
}

// GetAll implements product.ProductServiceInterface.
func (ps *productService) GetAll() ([]product.ProCore, error) {
	products, err := ps.productData.SelectAll()
	if err != nil {
		return nil, err
	}
	return products, nil
}

// Update implements product.ProductServiceInterface.
func (ps *productService) Update(UserID int, input product.ProCore) error {
	err := ps.productData.Update(UserID, input)
	if err != nil {
		return err
	}
	return nil
}

// Delete implements product.ProductServiceInterface.
func (ps *productService) Delete(UserID int, ProductID int) error {
	err := ps.productData.Delete(UserID, ProductID)
	if err != nil {
		return err
	}

	return nil
}

// GettProductUser implements product.ProductServiceInterface.
func (ps *productService) GettProductUser(UserID int) ([]product.ProCore, error) {
	products, err := ps.productData.SelectProductUser(UserID)
    if err != nil {
        return nil, err
    }

    return products, nil
}

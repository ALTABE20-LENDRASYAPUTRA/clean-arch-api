package data

import (
	"clean-arch-api/features/product"
	"clean-arch-api/features/user"
	"errors"

	"gorm.io/gorm"
)

type productQuery struct {
	db *gorm.DB
}


func New(db *gorm.DB) product.ProductDataInterface {
	return &productQuery{
		db: db,
	}
}

// Insert implements product.ProductDataInterface.
func (repo *productQuery) Insert(UserID int, input product.ProCore) error {
	productInputGorm := Product{
		UserID:      UserID,
		Name:        input.Name,
		Description: input.Description,
	}

	// simpan ke DB
	tx := repo.db.Create(&productInputGorm)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("insert failed, row affected = 0")
	}
	return nil
}

// SelectAll implements product.ProductDataInterface.
func (repo *productQuery) SelectAll() ([]product.ProCore, error) {
	var products []Product
	if err := repo.db.Preload("User").Find(&products).Error; err != nil {
		return nil, err
	}

	var result []product.ProCore
	for _, p := range products {
		userCore := user.Core{
			ID:   p.User.ID,
			Name: p.User.Name,
			Role: p.User.Role,
		}

		result = append(result, product.ProCore{
			ID:          p.ID,
			Name:        p.Name,
			Description: p.Description,
			CreatedAt:   p.CreatedAt,
			UpdatedAt:   p.UpdatedAt,
			User:        userCore,
		})
	}

	return result, nil
}

// Update implements product.ProductDataInterface.
func (repo *productQuery) Update(UserID int, input product.ProCore) error {
	var products Product

	if err := repo.db.First(&products, input.ID).Error; err != nil {
		return errors.New("comment not found")
	}

	if products.UserID != UserID {
		return errors.New("you are not authorized to update this product")
	}

	if err := repo.db.Model(&products).Updates(Product{
		Name:        input.Name,
		Description: input.Description,
	}).Error; err != nil {
		return err
	}

	return nil
}

// Delete implements product.ProductDataInterface.
func (repo *productQuery) Delete(UserID int, ProductID int) error {
	var products Product

	if err := repo.db.First(&products, ProductID).Error; err != nil {
		return errors.New("comment not found")
	}

	if products.UserID != UserID {
		return errors.New("you are not authorized to delete this product")
	}

	if err := repo.db.Delete(&products).Error; err != nil {
		return err
	}

	return nil
}

// SelectProductUser implements product.ProductDataInterface.
func (repo *productQuery) SelectProductUser(UserID int) ([]product.ProCore, error) {
	var products []Product
    if err := repo.db.Where("user_id = ?", UserID).Find(&products).Error; err != nil {
        return nil, err
    }

	var result []product.ProCore
    for _, p := range products {
        result = append(result, product.ProCore{
            ID:          p.ID,
            Name:        p.Name,
            Description: p.Description,
            CreatedAt:   p.CreatedAt,
            UpdatedAt:   p.UpdatedAt,
        })
    }

    return result, nil
}
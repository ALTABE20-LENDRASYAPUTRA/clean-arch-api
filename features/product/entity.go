package product

import (
	"clean-arch-api/features/user"
	"time"
)

type ProCore struct {
	ID          uint
	Name        string
	Description string
	UserID      int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	User        user.Core
}

// interface untuk Data Layer
type ProductDataInterface interface {
	Insert(UserID int, input ProCore) error
	SelectAll() ([]ProCore, error)
	Update(UserID int, input ProCore) error
	Delete(UserID int, ProductID int) error
	SelectProductUser(UserID int) ([]ProCore, error)
}

// interface untuk Service Layer
type ProductServiceInterface interface {
	Create(input ProCore) error
	GetAll() ([]ProCore, error)
	Update(UserID int, input ProCore) error
	Delete(UserID int, ProductID int)error
	GettProductUser(UserID int) ([]ProCore, error)
}

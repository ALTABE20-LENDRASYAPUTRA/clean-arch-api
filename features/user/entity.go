package user

import (
	"time"
)

type Core struct {
	ID          uint
	Name        string
	Email       string
	Password    string
	Address     string
	PhoneNumber string
	Role        string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// interface untuk Data Layer
type UserDataInterface interface {
	Insert(input Core) error
	SelectAll() ([]Core, error)
	Update(id int, input Core) (Core, error)
	Login(email string) (data *Core, err error)
	Delete(id int) error
}

// interface untuk Service Layer
type UserServiceInterface interface {
	Create(input Core) error
	GetAll() ([]Core, error)
	Update(id int, input Core) (Core, error)
	Login(email, password string) (data *Core, token string, err error)
	Delete(id int)error
}

package repo

import (
	"errors"

	"gorm.io/gorm"
)

var (
	ErrRecordNotFound  = errors.New("no record found")
	ErrDuplicateRecord = errors.New("record duplicate on unique constraint")
)

type Services struct {
	User    UserService
	Product ProductService
}

func New(db *gorm.DB) *Services {
	return &Services{
		User:    UserService{db: db},
		Product: ProductService{db: db},
	}
}

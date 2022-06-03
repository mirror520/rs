package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/mirror520/rs/model/shopping"
)

type shoppingRepository struct {
	db *gorm.DB
}

func NewShoppingRepository() (shopping.Repository, error) {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	repo := new(shoppingRepository)
	repo.db = db
	db.AutoMigrate(
		&shopping.Store{}, &shopping.Coupon{},
		&shopping.Product{}, &shopping.ProductStyle{}, &shopping.ProductImage{},
		&shopping.Customer{}, &shopping.Order{}, &shopping.Item{},
	)

	return repo, nil
}

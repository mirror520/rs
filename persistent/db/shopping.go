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
	db, err := gorm.Open(sqlite.Open("shopping.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	repo := new(shoppingRepository)
	repo.db = db

	err = db.AutoMigrate(
		&shopping.ProductStyle{}, &shopping.ProductImage{}, &shopping.Product{},
		&shopping.Coupon{}, &shopping.Store{},
		&shopping.Item{}, &shopping.Order{},
		&shopping.Customer{},
	)
	if err != nil {
		return nil, err
	}

	return repo, nil
}

func (repo *shoppingRepository) CreateStore(store *shopping.Store) error {
	tx := repo.db.Create(store)
	if err := tx.Error; err != nil {
		return err
	}
	return nil
}

func (repo *shoppingRepository) UpdateStore(store *shopping.Store) error {
	tx := repo.db.Save(store)
	if err := tx.Error; err != nil {
		return err
	}
	return nil
}

func (repo *shoppingRepository) CreateProduct(product *shopping.Product) error {
	tx := repo.db.Create(product)
	if err := tx.Error; err != nil {
		return err
	}
	return nil
}

func (repo *shoppingRepository) UpdateProduct(product *shopping.Product) error {
	tx := repo.db.Save(product)
	if err := tx.Error; err != nil {
		return err
	}
	return nil
}

func (repo *shoppingRepository) RegisterCustomer(customer *shopping.Customer) error {
	tx := repo.db.Create(customer)
	if err := tx.Error; err != nil {
		return err
	}
	return nil
}

func (repo *shoppingRepository) PlaceOrder(order *shopping.Order) error {
	tx := repo.db.Create(order)
	if err := tx.Error; err != nil {
		return err
	}
	return nil
}

func (repo *shoppingRepository) GetOrder(id uint) (*shopping.Order, error) {
	var o *shopping.Order
	result := repo.db.
		Preload("Customer").
		Preload("Store").
		Preload("Items").
		Preload("Items.Product").
		Preload("Items.ProductStyle").
		Preload("Items.Coupon").
		Take(&o, id)

	if err := result.Error; err != nil {
		return nil, err
	}

	return o, nil
}

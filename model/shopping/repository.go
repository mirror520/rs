package shopping

type Repository interface {
	CreateStore(store *Store) error
	UpdateStore(store *Store) error

	CreateProduct(product *Product) error
	UpdateProduct(product *Product) error

	RegisterCustomer(customer *Customer) error

	PlaceOrder(order *Order) error
	GetOrder(id uint) (*Order, error)
}

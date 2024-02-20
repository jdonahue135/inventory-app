package repository

import "github.com/jdonahue135/inventory/internal/models"

type Repository interface {
	GetProductByName(name string) (models.Product, error)
	GetCustomerByName(name string) (models.Customer, error)
	CountProducts() int
	CountCustomers() int
	AllOrders() []models.Order
	AllCustomers() map[string]models.Customer
	SaveProduct(product models.Product) error
	SaveCustomer(customer models.Customer) error
	UpdateProductQuantityByName(name string, quantity int) error
	OrderProduct(customerName, productName string, quantity int) error
}

package repo

import (
	"errors"

	"github.com/jdonahue135/inventory/internal/models"
)

// GetProductByName gets one product by name
func (m *testRepo) GetProductByName(name string) (models.Product, error) {
	var product models.Product
	if name == "socks" {
		return models.Product{
			Name:     name,
			Quantity: 100,
		}, nil
	}
	return product, errors.New("some error")
}

// GetCustomerByName gets one customer by name
func (m *testRepo) GetCustomerByName(name string) (models.Customer, error) {
	var customer models.Customer
	if name == "jake" {
		return customer, nil
	}
	return customer, errors.New("some error")
}

// CountProducts returns the total amount of products
func (m *testRepo) CountProducts() int {
	return 0
}

// CountCustomers returns the total amount of customers
func (m *testRepo) CountCustomers() int {
	return 1
}

// AllOrders returns all the orders
func (m *testRepo) AllOrders() []models.Order {
	orders := []models.Order{
		{CustomerName: "kate", ProductName: "hats", Price: 2050},
		{CustomerName: "kate", ProductName: "socks", Price: 3450},
		{CustomerName: "dan", ProductName: "socks", Price: 12075},
	}
	return orders
}

// AllCustomers returns all the customers
func (m *testRepo) AllCustomers() map[string]models.Customer {
	customers := make(map[string]models.Customer)
	customers["kate"] = models.Customer{Name: "kate"}
	customers["dan"] = models.Customer{Name: "dan"}
	customers["jake"] = models.Customer{Name: "jake"}

	return customers
}

// SaveProduct saves a product
func (m *testRepo) SaveProduct(product models.Product) error {
	if product.Name == "car" {
		return errors.New("some error")
	}
	return nil
}

// SaveCustomer saves a customer
func (m *testRepo) SaveCustomer(customer models.Customer) error {
	if customer.Name == "jake" {
		return errors.New("some error")
	}
	return nil
}

// UpdateProductQuantityByName updates the quantity of a given product
func (m *testRepo) UpdateProductQuantityByName(name string, quantity int) error {
	if name == "car" {
		return errors.New("some error")
	}
	return nil
}

// OrderProduct orders a product for a customer, creating that customer they do not exist
func (m *testRepo) OrderProduct(customerName, productName string, quantity int) error {
	if productName == "hat" || quantity == 100 {
		return errors.New("Some error")
	}
	return nil
}

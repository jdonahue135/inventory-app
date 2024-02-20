package repo

import (
	"errors"

	"github.com/jdonahue135/inventory/internal/models"
)

// GetProductByName gets one product by name
func (m *memoryRepo) GetProductByName(name string) (models.Product, error) {
	var product models.Product
	if _, ok := m.products[name]; !ok {
		return product, errors.New("Product not found: " + name)
	}

	return m.products[name], nil
}

// GetCustomerByName gets one customer by name
func (m *memoryRepo) GetCustomerByName(name string) (models.Customer, error) {
	var customer models.Customer
	if _, ok := m.customers[name]; !ok {
		return customer, errors.New("Customer not found: " + name)
	}

	return m.customers[name], nil
}

// CountProducts returns the total amount of products
func (m *memoryRepo) CountProducts() int {
	return len(m.products)
}

// CountCustomers returns the total amount of customers
func (m *memoryRepo) CountCustomers() int {
	return len(m.customers)
}

// AllOrders returns all the orders
func (m *memoryRepo) AllOrders() []models.Order {
	return m.orders
}

// AllCustomers returns all the customers
func (m *memoryRepo) AllCustomers() map[string]models.Customer {
	return m.customers
}

// SaveProduct saves a product
func (m *memoryRepo) SaveProduct(product models.Product) error {
	m.products[product.Name] = product
	return nil
}

// SaveCustomer saves a customer
func (m *memoryRepo) SaveCustomer(customer models.Customer) error {
	m.customers[customer.Name] = customer
	return nil
}

// UpdateProductQuantityByName updates the quantity of a given product
func (m *memoryRepo) UpdateProductQuantityByName(name string, quantity int) error {
	product := m.products[name]
	product.Quantity += quantity
	m.products[name] = product
	return nil
}

// OrderProduct orders a product for a customer
func (m *memoryRepo) OrderProduct(customerName, productName string, quantity int) error {
	if _, ok := m.customers[customerName]; !ok {
		return errors.New("Customer not found: " + customerName)
	}

	product, ok := m.products[productName]
	if !ok {
		return errors.New("Product not found: " + productName)
	}

	if product.Quantity < quantity {
		return nil
	}

	product.Quantity -= quantity
	m.products[productName] = product

	order := models.Order{
		ID:           len(m.orders),
		CustomerName: customerName,
		ProductName:  productName,
		Quantity:     quantity,
		Price:        product.Price * quantity,
	}
	m.orders = append(m.orders, order)

	return nil
}

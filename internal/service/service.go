package service

import (
	"errors"

	"github.com/jdonahue135/inventory/internal/models"
	"github.com/jdonahue135/inventory/internal/repository"
)

type Service struct {
	Repo repository.Repository
}

// NewService creates a new instance of the Service
func NewService(r repository.Repository) *Service {
	return &Service{Repo: r}
}

// RegisterProduct creates a new product
func (s *Service) RegisterProduct(name string, price int) error {
	// ignore re-registration of products
	if _, err := s.Repo.GetProductByName(name); err == nil {
		return nil
	}

	product := models.Product{
		ID:       s.Repo.CountProducts(),
		Name:     name,
		Price:    price,
		Quantity: 0,
	}

	if err := s.Repo.SaveProduct(product); err != nil {
		return errors.New("failed to register product: " + name)
	}

	return nil
}

// CheckInProduct updates the quantity of the product in our warehouse
func (s *Service) CheckInProduct(name string, quantity int) error {
	if _, err := s.Repo.GetProductByName(name); err != nil {
		return errors.New("Product doesn't exist: " + name)
	}

	if err := s.Repo.UpdateProductQuantityByName(name, quantity); err != nil {
		return errors.New("failed to update product quantity: " + name)
	}

	return nil
}

// OrderProduct processes an order for a product
func (s *Service) OrderProduct(customerName, productName string, quantity int) error {
	// check if customer exists, create if not
	if _, err := s.Repo.GetCustomerByName(customerName); err != nil {
		customer := models.Customer{
			ID:   s.Repo.CountCustomers(),
			Name: customerName,
		}
		s.Repo.SaveCustomer(customer)
	}

	// ignore orders for non-existent products
	product, err := s.Repo.GetProductByName(productName)
	if err != nil {
		return nil
	}

	//ignore orders where we don't have enough inventory
	if product.Quantity < quantity {
		return nil
	}

	if err := s.Repo.OrderProduct(customerName, productName, quantity); err != nil {
		return errors.New("failed to order product: " + productName + " for " + customerName)
	}

	return nil
}

// GetCustomerReports gets customer order history and returns in a DTO
func (s *Service) GetCustomerReports() map[string]CustomerReportDTO {
	orders := s.Repo.AllOrders()

	customerReports := make(map[string]CustomerReportDTO)

	for _, order := range orders {
		dto, ok := customerReports[order.CustomerName]
		if !ok {
			dto = CustomerReportDTO{
				CustomerName: order.CustomerName,
				TotalSpend:   0,
				TotalOrders:  0,
				Products:     make(map[string]int),
			}
		}
		dto.TotalSpend += order.Price
		dto.TotalOrders++
		if _, ok := dto.Products[order.ProductName]; !ok {
			dto.Products[order.ProductName] = order.Price
		} else {
			dto.Products[order.ProductName] += order.Price
		}

		customerReports[order.CustomerName] = dto
	}

	customers := s.Repo.AllCustomers()
	for _, customer := range customers {
		if _, ok := customerReports[customer.Name]; !ok {
			dto := CustomerReportDTO{
				CustomerName: customer.Name,
				TotalSpend:   0,
				TotalOrders:  0,
				Products:     make(map[string]int),
			}
			customerReports[customer.Name] = dto
		}
	}

	return customerReports
}

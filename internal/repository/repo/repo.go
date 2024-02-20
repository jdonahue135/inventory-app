package repo

import (
	"github.com/jdonahue135/inventory/internal/models"
	"github.com/jdonahue135/inventory/internal/repository"
)

type memoryRepo struct {
	products  map[string]models.Product
	customers map[string]models.Customer
	orders    []models.Order
}

type testRepo struct{}

// NewRepo returns a new in memory repository
func NewRepo() repository.Repository {
	return &memoryRepo{
		products:  make(map[string]models.Product),
		customers: make(map[string]models.Customer),
		orders:    make([]models.Order, 0),
	}
}

// NewTestRepo returns a new test repository
func NewTestRepo() repository.Repository {
	return &testRepo{}
}

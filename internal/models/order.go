package models

// Order is the order model
type Order struct {
	ID           int
	CustomerName string
	ProductName  string
	Quantity     int
	Price        int
}

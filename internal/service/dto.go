package service

// CustomerReportDTO represents the report data for a single customer
type CustomerReportDTO struct {
	CustomerName string
	TotalSpend   int
	TotalOrders  int
	Products     map[string]int
}

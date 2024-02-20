package report

import (
	"fmt"

	"github.com/jdonahue135/inventory/internal/service"
)

type Generator struct {
	Service *service.Service
}

// NewGenerator returns a new instance of a Generator
func NewGenerator(s *service.Service) *Generator {
	return &Generator{Service: s}
}

// GenerateReport gets report DTOs and generates a report from each
func (g *Generator) GenerateReport() {
	customerReports := g.Service.GetCustomerReports()
	for _, report := range customerReports {
		g.printCustomerReport(report)
	}
}

func (g *Generator) printCustomerReport(report service.CustomerReportDTO) {
	reportStr := fmt.Sprintf("%s: ", report.CustomerName)

	if report.TotalOrders == 0 {
		reportStr += "n/a"
	} else {
		for product, spend := range report.Products {
			reportStr += product + " - " + formatPrice(spend) + ", "
		}

		// remove trailing comma
		reportStr = reportStr[:len(reportStr)-2]

		averageOrderValue := calculateAverageOrderValue(report)
		reportStr += fmt.Sprintf(" | Average Order Value: $%.2f", averageOrderValue)
	}

	fmt.Println(reportStr)
}

func formatPrice(price int) string {
	dollars := price / 100
	cents := price % 100

	return fmt.Sprintf("$%d.%02d", dollars, cents)
}

func calculateAverageOrderValue(report service.CustomerReportDTO) float64 {
	if report.TotalOrders == 0 {
		return 0
	}
	return float64(report.TotalSpend) / float64(report.TotalOrders) / 100
}

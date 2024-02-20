package service

import (
	"os"
	"testing"

	"github.com/jdonahue135/inventory/internal/repository/repo"
)

var s *Service

func TestMain(m *testing.M) {
	repo := repo.NewTestRepo()
	s = NewService(repo)
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestRegisterProduct(t *testing.T) {
	// register a product that doesn't exist
	err := s.RegisterProduct("hat", 12345)
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	// try to register a product that exists
	err = s.RegisterProduct("socks", 12345)
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	// error registering product
	err = s.RegisterProduct("car", 120987)
	if err == nil {
		t.Errorf("Expected error, got no error")
	}
}

func TestCheckInProduct(t *testing.T) {
	// try to check in non-existent product
	err := s.CheckInProduct("hat", 12345)
	if err == nil {
		t.Errorf("Expected error, got no error")
	}

	// product exists
	err = s.CheckInProduct("socks", 12345)
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	// fail to update product
	err = s.CheckInProduct("car", 12345)
	if err == nil {
		t.Errorf("Expected error, got no error")
	}
}

func TestOrderProduct(t *testing.T) {
	// try to order non-existent product
	err := s.OrderProduct("kate", "car", 10)
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	// product exists
	err = s.OrderProduct("dan", "socks", 1)
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	// not enough quantity
	err = s.OrderProduct("dan", "socks", 1000)
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	// fail to order product
	err = s.OrderProduct("jake", "socks", 100)
	if err == nil {
		t.Errorf("Expected error, got no error")
	}
}

func TestGetCustomerReports(t *testing.T) {
	reports := s.GetCustomerReports()

	// Assert the results
	if len(reports) != 3 {
		t.Errorf("Expected 3 customer reports, got %d", len(reports))
	}

	kateReport, ok := reports["kate"]
	if !ok {
		t.Error("Expected customer report for kate, got none")
	}
	if kateReport.TotalSpend != 5500 {
		t.Errorf("Expected total spend for kate to be 5500, got %d", kateReport.TotalSpend)
	}
	if kateReport.TotalOrders != 2 {
		t.Errorf("Expected total orders for kate to be 2, got %d", kateReport.TotalOrders)
	}
	if len(kateReport.Products) != 2 {
		t.Errorf("Expected 2 products in kate's report, got %d", len(kateReport.Products))
	}
	if kateReport.Products["hats"] != 2050 {
		t.Errorf("Expected spend for hats in kate's report to be 2050, got %d", kateReport.Products["hats"])
	}
	if kateReport.Products["socks"] != 3450 {
		t.Errorf("Expected spend for socks in kate's report to be 3450, got %d", kateReport.Products["socks"])
	}

	danReport, ok := reports["dan"]
	if !ok {
		t.Error("Expected customer report for dan, got none")
	}
	if danReport.TotalSpend != 12075 {
		t.Errorf("Expected total spend for dan to be 12075, got %d", danReport.TotalSpend)
	}
	if danReport.TotalOrders != 1 {
		t.Errorf("Expected total orders for dan to be 1, got %d", danReport.TotalOrders)
	}
	if len(danReport.Products) != 1 {
		t.Errorf("Expected 1 product in dan's report, got %d", len(danReport.Products))
	}
	if danReport.Products["socks"] != 12075 {
		t.Errorf("Expected spend for socks in dan's report to be 12075, got %d", danReport.Products["socks"])
	}

	jakeReport, ok := reports["jake"]
	if !ok {
		t.Error("Expected customer report for jake, got none")
	}
	if jakeReport.TotalSpend != 0 {
		t.Errorf("Expected total spend for jake to be 12075, got %d", jakeReport.TotalSpend)
	}
	if jakeReport.TotalOrders != 0 {
		t.Errorf("Expected total orders for jake to be 1, got %d", jakeReport.TotalOrders)
	}
	if len(jakeReport.Products) != 0 {
		t.Errorf("Expected 1 product in jake's report, got %d", len(jakeReport.Products))
	}
}

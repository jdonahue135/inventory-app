package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jdonahue135/inventory/internal/handler"
	"github.com/jdonahue135/inventory/internal/report"
	"github.com/jdonahue135/inventory/internal/repository/repo"
	"github.com/jdonahue135/inventory/internal/service"
)

// main is the main function
func main() {
	// check if we have path to file as command-line argument
	if len(os.Args) < 2 {
		fmt.Println("Usage: main <path/to/input/file>")
		os.Exit(1)
	}

	// get file
	filePath := os.Args[1]
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error opening input file: %v", err)
	}
	defer file.Close()

	// initialize repository, service and handler
	repository := repo.NewRepo()
	s := service.NewService(repository)
	h := handler.NewHandler(s)

	// parse file
	if err := h.Parse(file); err != nil {
		log.Fatalf("Error parsing input: %v", err)
	}

	// generate report
	rg := report.NewGenerator(s)
	rg.GenerateReport()
}

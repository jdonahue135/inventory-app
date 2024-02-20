package handler

import (
	"log"
	"os"
	"testing"

	"github.com/jdonahue135/inventory/internal/repository/repo"
	"github.com/jdonahue135/inventory/internal/service"
)

var h *Handler

func TestMain(m *testing.M) {
	repo := repo.NewTestRepo()
	s := service.NewService(repo)
	h = NewHandler(s)
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestParse_ValidCommands(t *testing.T) {
	// create a temporary file for testing
	tempFile := getTempFile()
	err := writeToFile("register hats $20.50\norder kate hats 2\ncheckin socks 5\n", tempFile)
	if err != nil {
		t.Fatalf("Failed to write to temporary file: %v", err)
	}

	err = h.Parse(tempFile)
	if err != nil {
		t.Errorf("Parse returned an unexpected error: %v", err)
	}
	if err := tempFile.Close(); err != nil {
		log.Fatalf("Error closing temporary file: %v", err)
	}
}

var testCases = []struct {
	name  string
	input string
}{
	{
		name:  "missing argument - register",
		input: "register $20.50",
	},
	{
		name:  "invalid price - register",
		input: "register socks $20.001",
	},
	{
		name:  "string price - register",
		input: "register socks twenty",
	},
	{
		name:  "service error - register",
		input: "register car $20.50",
	},
	{
		name:  "missing argument - checkin",
		input: "checkin socks",
	},
	{
		name:  "invalid quantity - checkin",
		input: "checkin socks -1",
	},
	{
		name:  "float quantity - checkin",
		input: "checkin socks 1.5",
	},
	{
		name:  "string quantity - checkin",
		input: "checkin socks one",
	},
	{
		name:  "service error - checkin",
		input: "checkin car 1",
	},
	{
		name:  "missing argument - order",
		input: "order kate 2",
	},
	{
		name:  "invalid quantity - order",
		input: "order kate socks -2",
	},
	{
		name:  "float quantity - order",
		input: "order kate socks 2.5",
	},
	{
		name:  "string quantity - order",
		input: "order kate socks two",
	},
	{
		name:  "service error - order",
		input: "order kate socks 100",
	},
	{
		name:  "invalid command",
		input: "orde kate 2",
	},
}

func TestParse_InvalidCommands(t *testing.T) {
	for _, testCase := range testCases {
		tempFile := getTempFile()
		err := writeToFile(testCase.input, tempFile)
		if err != nil {
			t.Fatalf("Failed to write to temporary file: %v", err)
		}

		err = h.Parse(tempFile)
		if err == nil {
			t.Errorf("Parse returned no error for test case: %s", testCase.name)
		}

		if err := tempFile.Close(); err != nil {
			log.Fatalf("Error closing temporary file: %v", err)
		}
	}
}

func getTempFile() *os.File {
	tempFile, err := os.CreateTemp("", "testfile")
	if err != nil {
		log.Fatalf("Error creating temporary file: %v", err)
	}
	return tempFile
}

func writeToFile(str string, tempFile *os.File) error {
	_, err := tempFile.WriteString(str)
	if err != nil {
		return err
	}
	// reset the file cursor to the beginning of the file
	if _, err := tempFile.Seek(0, 0); err != nil {
		return err
	}
	return nil
}

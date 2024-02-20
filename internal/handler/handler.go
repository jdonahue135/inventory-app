package handler

import (
	"bufio"
	"errors"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/jdonahue135/inventory/internal/service"
)

const (
	CommandRegister = "register"
	CommandCheckIn  = "checkin"
	CommandOrder    = "order"
)

const (
	CommandRegisterArgLength = 2
	CommandCheckInArgLength  = 2
	CommandOrderArgLength    = 3
)

type Handler struct {
	Service *service.Service
}

// NewHandler provides a new instance of a Handler
func NewHandler(s *service.Service) *Handler {
	return &Handler{Service: s}
}

// Parse reads commands from the input file, validates, processes, and calls the corresponding Service functions
func (h *Handler) Parse(file *os.File) error {
	scanner := bufio.NewScanner(file)

	// read file line-by-line
	for scanner.Scan() {
		line := scanner.Text()
		args := strings.Fields(line)

		command := args[0]
		args = args[1:]

		switch command {
		case CommandRegister:
			if err := h.register(args); err != nil {
				return err
			}
		case CommandCheckIn:
			if err := h.checkIn(args); err != nil {
				return err
			}
		case CommandOrder:
			if err := h.order(args); err != nil {
				return err
			}
		default:
			return errors.New("unknown command: " + command)
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func (h *Handler) register(args []string) error {
	if len(args) != CommandRegisterArgLength {
		return errors.New("incorrect number of arguments provided, wanted 3. Usage: register product_name price")
	}

	if isValidPrice(args[1]) == false {
		return errors.New("invalid price: " + args[1])
	}

	name, price := args[0], parsePrice(args[1])

	err := h.Service.RegisterProduct(name, price)
	if err != nil {
		return err
	}

	return nil
}

func (h *Handler) checkIn(args []string) error {
	if len(args) != CommandCheckInArgLength {
		return errors.New("incorrect number of arguments provided, wanted 3. Usage: checkin product_name quantity")
	}

	quantity, err := strconv.Atoi(args[1])
	if err != nil || !isValidInt(quantity) {
		return errors.New("quantity must be a valid integer. Got " + args[1])
	}

	name := args[0]

	err = h.Service.CheckInProduct(name, quantity)
	if err != nil {
		return err
	}

	return nil
}

func (h *Handler) order(args []string) error {
	if len(args) != CommandOrderArgLength {
		return errors.New("incorrect number of arguments provided, wanted 4. Usage: order customer_name product_name quantity")
	}

	quantity, err := strconv.Atoi(args[2])
	if err != nil || !isValidInt(quantity) {
		return errors.New("quantity must be a valid integer. Got " + args[1])
	}

	customerName, productName := args[0], args[1]

	err = h.Service.OrderProduct(customerName, productName, quantity)
	if err != nil {
		return err
	}

	return nil
}

func isValidPrice(price string) bool {
	if price[0] != '$' {
		return false
	}

	parts := strings.Split(price[1:], ".")

	// check that there is exactly one decimal
	if len(parts) != 2 {
		return false
	}

	//check for valid decimal places
	if len(parts[1]) != 2 {
		return false
	}

	dollars, err := strconv.Atoi(parts[0])
	if err != nil || dollars < 0 {
		return false
	}

	cents, err := strconv.Atoi(parts[1])
	if err != nil || cents < 0 || cents >= 100 {
		return false
	}

	return true
}

func isValidInt(i int) bool {
	return i > 0 && i < math.MaxInt32
}

func parsePrice(str string) int {
	priceArr := strings.Split(str[1:], ".")

	dollars, _ := strconv.Atoi(priceArr[0])
	cents, _ := strconv.Atoi(priceArr[1])

	totalCents := dollars*100 + cents
	return totalCents
}

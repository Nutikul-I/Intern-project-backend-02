package handler

import (
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

type CustomerHandler interface {
	GetCustomerData(c *fiber.Ctx) error
	CreateCustomerData(c *fiber.Ctx) error
	DeleteCustomerData(c *fiber.Ctx) error
}

type customerHandler struct{}

func NewCustomerHandler() CustomerHandler {
	return &customerHandler{}
}

func (ch *customerHandler) GetCustomerData(c *fiber.Ctx) error {
	log.Println("GetCustomerData handler called")
	// Implement logic for fetching customer data
	return c.SendString("Fetched customer data from handler")
}

func (ch *customerHandler) CreateCustomerData(c *fiber.Ctx) error {
	log.Println("CreateCustomerData handler called")
	// Implement logic for creating customer data
	return c.SendString("Created customer data in handler")
}

func (ch *customerHandler) DeleteCustomerData(c *fiber.Ctx) error {
	log.Println("DeleteCustomerData handler called")
	// Implement logic for deleting customer data
	return c.SendString("Deleted customer data in handler")
}

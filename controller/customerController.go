package controller

import (
	"payso-internal-api/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

type CustomerController interface {
	GetCustomer(c *fiber.Ctx) error
}

type customerController struct {
	customerService service.CustomerService
}

func NewCustomerController(customerService service.CustomerService) CustomerController {
	return &customerController{customerService: customerService}
}

func (ctl *customerController) GetCustomer(c *fiber.Ctx) error {
	log.Infof("==-- GetCustomer --==")

	var Page int = 1
	var Row int = 50

	RequestMID := c.Query("MID")
	RequestPage, err := strconv.Atoi(c.Query("Page"))
	RequestRow, err := strconv.Atoi(c.Query("Row"))

	if RequestPage > 0 {
		Page = RequestPage
	}

	if RequestRow > 0 {
		Row = RequestRow
	}

	res, err := ctl.customerService.GetcustomerService(RequestMID, Page, Row)
	if err != nil {
		log.Errorf("GetCustomer Error from service GetCustomer: %v", err)
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "API Failed.",
			"data":    err,
		})
	}

	totalPages := res.TotalPages / Row
	if res.TotalPages%Row != 0 {
		totalPages++
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":       200,
		"message":      "GetCustomer",
		"data":         res.CustomerList,
		"totalPages":   totalPages,
		"nextPage":     Page + 1,
		"previousPage": Page - 1,
		"currentPage":  Page,
	})
}

func (cc *customerController) UpdateCustomer(c *fiber.Ctx) error {
	log.Info("UpdateCustomer called")
	// Implement logic for updating a customer
	return c.SendString("UpdateCustomer")
}

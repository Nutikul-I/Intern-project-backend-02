package controller

import (
	"payso-internal-api/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

type EmployeesController interface {
	GetEmployees(c *fiber.Ctx) error
}

type employeesController struct {
	employeesService service.EmployeesService
}

func NewEmployeesController(employeesService service.EmployeesService) EmployeesController {
	return &employeesController{employeesService: employeesService}
}

func (ctl *employeesController) GetEmployees(c *fiber.Ctx) error {
	log.Infof("==-- GetEmployees --==")

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

	res, err := ctl.employeesService.GetEmployeesService(RequestMID, Page, Row)
	if err != nil {
		log.Errorf("GetEmployees Error from service GetEmployees: %v", err)
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
		"message":      "GetEmployees",
		"data":         res.EmployeesList,
		"totalPages":   totalPages,
		"nextPage":     Page + 1,
		"previousPage": Page - 1,
		"currentPage":  Page,
	})
}

// func (cc *employeesController) UpdateEmployees(c *fiber.Ctx) error {
// 	log.Info("UpdateEmployees called")
// 	// Implement logic for updating a employees
// 	return c.SendString("UpdateEmployees")
// }

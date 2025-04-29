package controller

import (
	"payso-internal-api/model"
	"payso-internal-api/service"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

type EmployeesController interface {
	GetEmployees(c *fiber.Ctx) error
	CreateEmployees(c *fiber.Ctx) error
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

func (ctl *employeesController) CreateEmployees(c *fiber.Ctx) error {
	log.Infof("==-- CreateCustomer --==")

	var payload model.CreateEmployeesPayload

	if err := c.BodyParser(&payload); err != nil {
		log.Errorf("Create Connection Type Error parsing")
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "API Failed.",
			"data":    err,
		})
	}

	res, err := ctl.employeesService.CreateEmployeesService(payload, c.IP())
	if err != nil {
		log.Errorf("CreateCustomer Error from service CreateCustomer: %v", err)
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "API Failed.",
			"data":    err,
		})
	}

	if res.StatusCode == 400 {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  400,
			"message": res.Message,
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  200,
		"message": "CreateCustomer",
		"data":    nil,
	})

}

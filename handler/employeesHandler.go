package handler

import (
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

type EmployeesHandler interface {
	GetEmployeesData(c *fiber.Ctx) error
	CreateEmployeesData(c *fiber.Ctx) error
	DeleteEmployeesData(c *fiber.Ctx) error
}

type employeesHandler struct{}

func NewEmployeesHandler() EmployeesHandler {
	return &employeesHandler{}
}

func (ch *employeesHandler) GetEmployeesData(c *fiber.Ctx) error {
	log.Println("GetEmployeesData handler called")
	return c.SendString("Fetched employees data from handler")
}

func (ch *employeesHandler) CreateEmployeesData(c *fiber.Ctx) error {
	log.Println("CreateEmployeesData handler called")
	return c.SendString("Created Employees data in handler")
}

func (ch *employeesHandler) DeleteEmployeesData(c *fiber.Ctx) error {
	log.Println("DeleteEmployeesData handler called")
	return c.SendString("Deleted Employees data in handler")
}

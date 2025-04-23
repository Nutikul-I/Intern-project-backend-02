package router

import (
	"payso-internal-api/controller"
	"payso-internal-api/handler"
	"payso-internal-api/service"
	"strings"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

func SetupRoutes(app *fiber.App) {

	merchantController := controller.NewMerchantController(service.NewMerchantService(handler.NewMerchantHandler()))
	customerController := controller.NewCustomerController(service.NewcustomerService(handler.NewCustomerHandler()))
	employeesController := controller.NewEmployeesController(service.NewEmployeesService(handler.NewEmployeesHandler()))

	api := app.Group("/", func(c *fiber.Ctx) error {
		if !strings.Contains(c.Request().URI().String(), "/ping") {
			log.Infof("all : %v", c.Request().URI().String())
		}
		return c.Next()
	})

	merchant := api.Group("/api/merchant")
	merchant.Get("/merchant", merchantController.GetMerchant)
	merchant.Post("/create-merchant", merchantController.CreateMerchant)
	merchant.Delete("/delete-merchant", merchantController.DeleteMerchant)

	customer := api.Group("/api/customer")
	customer.Get("/customer", customerController.GetCustomer)
	// customer.Post("/create-customer", customerController.CreateCustomer)
	// customer.Delete("/delete-customer", customerController.DeleteCustomer)

	employees := api.Group("/api/employees")
	employees.Get("/employees", employeesController.GetEmployees)
	// employees.Post("/create-employees", employeesController.CreateEmployees)
	// employees.Delete("/delete-employees", employeesController.DeleteEmployees)
}

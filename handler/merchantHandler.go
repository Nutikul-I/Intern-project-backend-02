package handler

import (
	"github.com/gofiber/fiber/v2"
)

type MerchantHandler interface {
	GetMerchant(c *fiber.Ctx) error
	CreateMerchant(c *fiber.Ctx) error
	DeleteMerchant(c *fiber.Ctx) error
}

type merchantHandler struct{}

func NewMerchantHandler() MerchantHandler {
	return &merchantHandler{}
}

func (h *merchantHandler) GetMerchant(c *fiber.Ctx) error {
	data := fiber.Map{
		"merchant_id":   "12345",
		"merchant_name": "Example Store",
	}

	return c.JSON(data)
}

func (h *merchantHandler) CreateMerchant(c *fiber.Ctx) error {
	var body struct {
		Name string `json:"name"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid request"})
	}

	return c.JSON(fiber.Map{
		"message": "Merchant created successfully",
		"name":    body.Name,
	})
}

func (h *merchantHandler) DeleteMerchant(c *fiber.Ctx) error {
	id := c.Query("id")

	if id == "" {
		return c.Status(400).JSON(fiber.Map{"message": "Merchant ID is required"})
	}

	return c.JSON(fiber.Map{
		"message":     "Merchant deleted successfully",
		"merchant_id": id,
	})
}

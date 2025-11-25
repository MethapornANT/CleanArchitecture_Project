package handler

import (
	"Structure-Project/internal/core/domain"
	"Structure-Project/internal/core/service"
	"database/sql"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// GetCustomers
func GetCustomers(c *fiber.Ctx) error {
	customers, err := service.GetCustomers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(customers)
}

// GetCustomerByID
func GetCustomerByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("customer_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Customer ID format."})
	}

	cust, err := service.GetCustomerByID(id)
	if err == sql.ErrNoRows {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Customer not found."})
	} else if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(cust)
}

// CreateCustomer
func CreateCustomer(c *fiber.Ctx) error {
	var cust domain.CustomerModel
	if err := c.BodyParser(&cust); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body."})
	}

	if err := service.CreateCustomer(&cust); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusCreated)
}

// UpdateCustomer
func UpdateCustomer(c *fiber.Ctx) error {
	var cust domain.CustomerModel
	if err := c.BodyParser(&cust); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body."})
	}

	found, err := service.UpdateCustomer(&cust)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	if !found {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Customer ID not found for update."})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// DeleteCustomer
func DeleteCustomer(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("customer_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Customer ID format."})
	}

	found, err := service.DeleteCustomer(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	if !found {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Customer ID not found for deletion."})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// GetVehiclesByCustomerID
func GetVehiclesByCustomerID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("customer_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Customer ID format."})
	}

	vehicles, err := service.GetVehiclesByCustomerID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(vehicles)
}

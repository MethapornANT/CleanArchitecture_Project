package handler

import (
	"Structure-Project/internal/core/domain"
	"Structure-Project/internal/core/service"
	"database/sql"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// GetVehicles
func GetVehicles(c *fiber.Ctx) error {
	vehicles, err := service.GetVehicles()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(vehicles)
}

// GetVehicleByID
func GetVehicleByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("vehicle_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	vehicle, err := service.GetVehicleByID(id)
	if err == sql.ErrNoRows {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Vehicle not found"})
	} else if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(vehicle)
}

// GetCustomersByVehicleID
func GetCustomersByVehicleID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("vehicle_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	customers, err := service.GetCustomersByVehicleID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if len(customers) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Owner not found"})
	}

	return c.JSON(customers)
}

// GetVehicleAndCustomerByVehicleID
func GetVehicleAndCustomerByVehicleID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("vehicle_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	data, err := service.GetVehicleAndCustomerByVehicleID(id)
	if err == sql.ErrNoRows {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Data not found"})
	} else if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(data)
}

// CreateVehicle
func CreateVehicle(c *fiber.Ctx) error {
	var v domain.VehicleModel
	if err := c.BodyParser(&v); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	if err := service.CreateVehicle(&v); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Vehicle created successfully"})
}

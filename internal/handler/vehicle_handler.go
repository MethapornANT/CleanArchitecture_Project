package handler

import (
	"Structure-Project/internal/core/domain"
	"Structure-Project/internal/core/service"
	"database/sql"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type VehicleHandler struct {
	service *service.VehicleService
}

func NewVehicleHandler(service *service.VehicleService) *VehicleHandler {
	return &VehicleHandler{service: service}
}

// 1. GetVehicles
func (h *VehicleHandler) GetVehicles(c *fiber.Ctx) error {
	vehicles, err := h.service.GetVehicles()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(vehicles)
}

// 2. GetVehicleByID
func (h *VehicleHandler) GetVehicleByID(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("vehicle_id"))

	vehicle, err := h.service.GetVehicleByID(id)
	if err == sql.ErrNoRows {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Vehicle not found"})
	} else if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(vehicle)
}

// 3. GetCustomersByVehicleID
func (h *VehicleHandler) GetCustomersByVehicleID(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("vehicle_id"))

	customers, err := h.service.GetCustomersByVehicleID(id)
	if err != nil {
		if err.Error() == "data_not_found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Owner not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(customers)
}

// 4. GetVehicleAndCustomerByVehicleID
func (h *VehicleHandler) GetVehicleAndCustomerByVehicleID(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("vehicle_id"))

	data, err := h.service.GetVehicleAndCustomerByVehicleID(id)
	if err == sql.ErrNoRows {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Data not found"})
	} else if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(data)
}

// 5. CreateVehicle
func (h *VehicleHandler) CreateVehicle(c *fiber.Ctx) error {
	var v domain.VehicleModel
	if err := c.BodyParser(&v); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	if err := h.service.CreateVehicle(&v); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusCreated)
}

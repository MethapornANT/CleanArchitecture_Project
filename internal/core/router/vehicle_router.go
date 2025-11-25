package router

import (
	"Structure-Project/internal/handler"
	"github.com/gofiber/fiber/v2"
)

// เพิ่ม h เข้ามาใน Parameter
func RVehicle(app *fiber.App, h *handler.VehicleHandler) {
	api := app.Group("/api/vehicles")

	// เรียกใช้ Method ผ่านตัวแปร h
	api.Get("/", h.GetVehicles)
	api.Get("/:vehicle_id", h.GetVehicleByID)
	api.Get("/:vehicle_id/customers", h.GetCustomersByVehicleID)
	api.Get("/:vehicle_id/vehicleandcustomer", h.GetVehicleAndCustomerByVehicleID)
	
	api.Post("/", h.CreateVehicle)
}
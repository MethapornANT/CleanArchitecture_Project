package router

import (
	"Structure-Project/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func RVehicle(app *fiber.App) {
	api := app.Group("/api/vehicles")
	api.Get("/", handler.GetVehicles)
	api.Get("/:vehicle_id", handler.GetVehicleByID)
	api.Get("/:vehicle_id/customers", handler.GetCustomersByVehicleID)
	api.Get("/:vehicle_id/vehicleandcustomer", handler.GetVehicleAndCustomerByVehicleID)
	app.Post("/api/vehicles", handler.CreateVehicle)
}

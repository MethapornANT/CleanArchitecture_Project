package router

import (
	"Structure-Project/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func RCustomer(app *fiber.App) {
	customer := app.Group("/api/customers")
	customer.Get("/", handler.GetCustomers)
	customer.Get("/:customer_id", handler.GetCustomerByID)
	customer.Get("/:customer_id/vehicles", handler.GetVehiclesByCustomerID)
	customer.Post("/", handler.CreateCustomer)
	customer.Put("/:customer_id", handler.UpdateCustomer)
	customer.Delete("/:customer_id", handler.DeleteCustomer)
}

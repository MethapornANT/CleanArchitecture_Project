package main

import (
	"Structure-Project/internal/core/router"
	"Structure-Project/internal/core/service"
	"Structure-Project/internal/handler"
	"Structure-Project/internal/repository"
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

// สร้าง Func นี้มาจัดการความรก
func SetupRoutes(app *fiber.App, db *sql.DB) {

	// --- 1. Group: Vehicle ---
	vRepo := repository.NewVehicleRepository(db)
	vSvc := service.NewVehicleService(vRepo)
	vHandler := handler.NewVehicleHandler(vSvc)
	router.RVehicle(app, vHandler) // ส่ง handler เข้าไป

	// // --- 2. Group: Customer ---
	// cRepo := repository.NewCustomerRepository(db)
	// cSvc := service.NewCustomerService(cRepo)
	// cHandler := handler.NewCustomerHandler(cSvc)
	// router.RCustomer(app, cHandler)
}

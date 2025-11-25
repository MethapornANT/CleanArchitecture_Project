package main

import (
	"Structure-Project/internal/core/router"
	"Structure-Project/pkg/database"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// 1. สร้าง Fiber App
	app := fiber.New()

	// 2. เรียก SetupRoutes (จากไฟล์ setup.go)
	// ส่ง database.DB (Global Variable ที่ถูก init อัตโนมัติแล้ว) เข้าไป
	SetupRoutes(app, database.DB)
	router.RCustomer(app)

	// 3. กำหนด Port
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	// 4. Start Server
	log.Printf("Server is starting on port %s...", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatal(err)
	}

	// ปิด DB เมื่อจบการทำงาน
	defer database.DB.Close()
}

package main

import (
	"Structure-Project/internal/core/router"
	"Structure-Project/pkg/database"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	// นำเข้าแพ็กเกจ database เพื่อให้ init() ในนั้นทำงาน
)

func main() {

	app := fiber.New()

	router.RCustomer(app)
	router.RVehicle(app)

	port := "8080" // หรือจะอ่านจาก Environment Variable ด้วย os.Getenv ก็ได้
	log.Printf("Server is starting on port %s...", port)
	if err := app.Listen(":8080"); err != nil {
		log.Fatal(err)
	}
	defer database.DB.Close()

}

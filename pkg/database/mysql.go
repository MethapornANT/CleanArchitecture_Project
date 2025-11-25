package database

import (
	"database/sql"
	"fmt"
	"log"
	"Structure-Project/pkg/config"

	_ "github.com/go-sql-driver/mysql"
)

// DB เป็น Global Variable สำหรับ Database Connection Pool
var DB *sql.DB

// InitializeDatabase จะใช้ค่าตั้งค่าที่โหลดแล้วเพื่อเชื่อมต่อฐานข้อมูล
func InitializeDatabase() {
	// 1. ดึงค่าตั้งค่า
	config := config.GlobalConfig

	// 2. สร้าง Database Source Name (dbConfig)
	dbConfig := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		config.DatabaseUser, config.DatabasePassword, config.DatabaseHost, config.DatabasePort, config.DatabaseName)

	// 3. เปิด Connection
	db, err := sql.Open("mysql", dbConfig)
	if err != nil {
		log.Fatalf("Error setting up database connection: %v", err)
	}

	// 4. ตรวจสอบการเชื่อมต่อ
	if err = db.Ping(); err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	// 5. ตั้งค่า Connection Pool
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)

	// กำหนดค่า Connection Pool ให้ Global Variable ชื่อ DB
	DB = db
	log.Println("Successfully connected to MySQL database!")
}

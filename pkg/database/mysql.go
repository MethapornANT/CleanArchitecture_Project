package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/joho/godotenv"
)

// DB เป็น Global Variable สำหรับ Connection Pool ที่เราจะใช้ทั่วทั้งแอปพลิเคชัน
var DB *sql.DB

// init() function จะถูกเรียกโดยอัตโนมัติก่อน main()
func init() {
	// 1. โหลดไฟล์ .env
	// ถ้าไฟล์ .env อยู่ใน Root Directory, godotenv.Load() จะหาเจออัตโนมัติ
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: Could not load .env file. Falling back to system environment variables.")
		// ไม่ต้อง panic เพราะอาจจะรันใน Production ที่ใช้ Env Var จากระบบโดยตรง
	}

	// 2. อ่านค่าตัวแปร Environment
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// 3. สร้าง Connection String (DSN - Data Source Name)
	// รูปแบบ: username:password@tcp(host:port)/database_name?param=value
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		user, password, host, port, dbName)

	// 4. เปิด Database Connection (ยังไม่ได้เชื่อมต่อจริง)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		// หากมีปัญหาในการตั้งค่า (เช่น DSN ผิดรูปแบบ), ควรหยุดโปรแกรม
		log.Fatalf("Error setting up database connection: %v", err)
	}

	// 5. ตรวจสอบการเชื่อมต่อจริง (Ping)
	if err = db.Ping(); err != nil {
		// หาก DB Server ไม่พร้อมใช้งาน, ควรหยุดโปรแกรม
		log.Fatalf("Error connecting to the database: %v", err)
	}

	// 6. ตั้งค่า Connection Pool (แนะนำสำหรับ Production)
	db.SetMaxOpenConns(25) // จำนวน Connection สูงสุดที่เปิดพร้อมกัน
	db.SetMaxIdleConns(25) // จำนวน Connection ที่พร้อมใช้งานเมื่อว่าง

	DB = db // กำหนดค่า Connection Pool ให้ Global Variable
	log.Println("Successfully connected to MySQL database!")
}

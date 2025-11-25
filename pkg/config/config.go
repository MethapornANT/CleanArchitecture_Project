package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Configuration เป็น Struct สำหรับเก็บค่าตั้งค่าทั้งหมดของแอปพลิเคชัน
type Configuration struct {
	// การตั้งค่า Server
	ServerPort int
	JWTSecret  string
	APIVersion string

	// การตั้งค่า Database Connection
	DatabaseUser     string
	DatabasePassword string
	DatabaseHost     string
	DatabasePort     string
	DatabaseName     string
}

// GlobalConfig เป็นตัวแปร Global ที่เก็บค่าตั้งค่าที่โหลดแล้ว
var GlobalConfig Configuration

// getEnvOrDefault อ่านค่าจาก Environment Variable หรือใช้ค่าเริ่มต้นหากไม่พบ
func getEnvOrDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// LoadConfiguration จะโหลดค่าตั้งค่าทั้งหมดจาก .env และ Environment Variables
func LoadConfiguration() {
	// โหลดไฟล์ .env (ทำเพียงครั้งเดียว)
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: Could not load .env file. Using system environment variables or default values.")
	}

	// 1. กำหนดค่า Server
	serverPortStr := getEnvOrDefault("SERVER_PORT", "8080")
	serverPort, _ := strconv.Atoi(serverPortStr)

	// 2. กำหนดค่า Database
	dbUser := getEnvOrDefault("DB_USER", "root")
	dbPassword := getEnvOrDefault("DB_PASSWORD", "1234")
	dbHost := getEnvOrDefault("DB_HOST", "localhost")
	dbPort := getEnvOrDefault("DB_PORT", "3306")
	dbName := getEnvOrDefault("DB_NAME", "usercardb")

	// 3. ใส่ค่าทั้งหมดลงใน GlobalConfig
	GlobalConfig = Configuration{
		ServerPort: serverPort,
		JWTSecret:  os.Getenv("JWT_SECRET"),
		APIVersion: os.Getenv("API_VERSION"),

		DatabaseUser:     dbUser,
		DatabasePassword: dbPassword,
		DatabaseHost:     dbHost,
		DatabasePort:     dbPort,
		DatabaseName:     dbName,
	}
}

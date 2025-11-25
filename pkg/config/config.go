package config

import (
    "log"
    "os"
    "strconv"

    "github.com/joho/godotenv"
)

// Struct สำหรับเก็บค่า Config ทั้งหมด
type AppConfig struct {
    ServerPort int
    JWTSecret  string
    APIVersion string
}

var Config AppConfig

func LoadConfig() {
    // โหลด .env
    if err := godotenv.Load(); err != nil {
        log.Println("Warning: Could not load .env file.")
    }

    // กำหนดค่า Default
    portStr := os.Getenv("SERVER_PORT")
    if portStr == "" {
        portStr = "8080"
    }
    port, _ := strconv.Atoi(portStr)

    // ใส่ค่าลง Struct
    Config = AppConfig{
        ServerPort: port,
        JWTSecret:  os.Getenv("JWT_SECRET"),
        APIVersion: os.Getenv("API_VERSION"),
    }
}
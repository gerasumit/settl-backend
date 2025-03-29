package config

import (
    "log"
    "os"

    "github.com/joho/godotenv"
)

type Config struct {
    Port       string
    DBHost     string
    DBPort     string
    DBUser     string
    DBPassword string
    DBName     string
    GoogleClientID string
    GoogleClientSecret string
}

var AppConfig Config

func LoadConfig() {
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found, using system environment variables")
    }

    AppConfig = Config{
        Port:       getEnv("PORT", "8080"),
        DBHost:     getEnv("DB_HOST", "localhost"),
        DBPort:     getEnv("DB_PORT", "5432"),
        DBUser:     getEnv("DB_USER", "postgres"),
        DBPassword: getEnv("DB_PASSWORD", ""),
        DBName:     getEnv("DB_NAME", "settl"),
        GoogleClientID: getEnv("GOOGLE_CLIENT_ID", ""),
        GoogleClientSecret: getEnv("GOOGLE_CLIENT_SECRET", ""),
    }
}

func getEnv(key, fallback string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return fallback
}
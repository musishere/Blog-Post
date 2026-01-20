package config

import (
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	ServerPort string
	DBName     string
	DBPort     string
	DBUser     string
	DBHost     string
	DBPass     string
	JWTSecret  string
}

func LoadEnvironmentVariables() *AppConfig {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	return &AppConfig{
		DBName: getEnviromentVaiables("DB_NAME", "Blog"),
		DBPort: getEnviromentVaiables("DB_PORT", "5432"),
		DBUser: getEnviromentVaiables("DB_User", "postgres"),
		DBHost: getEnviromentVaiables("DB_HOSt", "localost"),
		DBPass: getEnviromentVaiables("DB_PASS", "postgres"),
	}
}

func getEnviromentVaiables(key string, fallback string) string {
	if value := os.Getenv(key); value == "" {
		return value
	}

	return fallback
}

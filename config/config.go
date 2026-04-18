package config

import (
	"log"
	"os"
	"strconv"
	"time"
)

// DatabaseConfig holds PostgreSQL database connection settings.
type DatabaseConfig struct {
	Host               string
	Port               string
	User               string
	Password           string
	Name               string
	DbUrl              string
	SSLMode            string
	MaxOpenConns       int
	MaxIdleConns       int
	ConnMaxLifetimeSec int
}

// AppConfig is the centralized configuration struct for the application.
// It is populated from environment variables. Sensitive values are not hardcoded.
type AppConfig struct {
	AppHost  string
	Port     string
	Timezone string
	Location *time.Location
	Database DatabaseConfig
}

// Load reads configuration from environment variables.
// A .env file can be auto-loaded by importing: _ "github.com/joho/godotenv/autoload"
// This function does not require a .env file; real environment variables take precedence.
func Load() *AppConfig {
	tzStr := getEnv("APP_TZ", "Asia/Jakarta")
	loc, err := time.LoadLocation(tzStr)
	if err != nil {
		log.Printf("warning: invalid APP_TZ %q, falling back to Asia/Jakarta: %v", tzStr, err)
		// Fallback to Asia/Jakarta (UTC+7)
		loc = time.FixedZone("Asia/Jakarta", 7*60*60)
	}

	return &AppConfig{
		AppHost:  getEnv("APP_HOST", "localhost:8080"),
		Port:     getEnv("PORT", "8080"), // default only for non-sensitive value
		Timezone: tzStr,
		Location: loc,
		Database: DatabaseConfig{
			Host:               getEnv("DB_HOST", ""),
			Port:               getEnv("DB_PORT", "5432"),
			User:               getEnv("DB_USER", ""),
			Password:           getEnv("DB_PASSWORD", ""),
			Name:               getEnv("DB_NAME", ""),
			DbUrl:              getEnv("DB_URL", ""),
			SSLMode:            getEnv("DB_SSLMODE", "disable"),
			MaxOpenConns:       getEnvInt("DB_MAX_OPEN_CONNS", 10),
			MaxIdleConns:       getEnvInt("DB_MAX_IDLE_CONNS", 5),
			ConnMaxLifetimeSec: getEnvInt("DB_CONN_MAX_LIFETIME_SEC", 300),
		},
	}
}

func getEnv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

func getEnvBool(key string, def bool) bool {
	if v := os.Getenv(key); v != "" {
		b, err := strconv.ParseBool(v)
		if err == nil {
			return b
		}
	}
	return def
}

func getEnvInt(key string, def int) int {
	if v := os.Getenv(key); v != "" {
		i, err := strconv.Atoi(v)
		if err == nil {
			return i
		}
	}
	return def
}

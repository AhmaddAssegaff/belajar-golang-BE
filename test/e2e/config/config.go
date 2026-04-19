package e2e

import "os"

var (
	BaseURL   = getEnv("TEST_BASE_URL", "http://localhost:8081")
	TestDB    = getEnv("TEST_DB_URL", "")
	HealthURL = BaseURL + "/health"
)

func getEnv(key, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	return val
}

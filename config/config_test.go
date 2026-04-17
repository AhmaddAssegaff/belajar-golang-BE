package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	t.Setenv("DB_HOST", "test-host")
	t.Setenv("DB_MAX_OPEN_CONNS", "20")
	t.Setenv("MINIO_USE_SSL", "true")

	cfg := Load()

	assert.Equal(t, "test-host", cfg.Database.Host)
	assert.Equal(t, 20, cfg.Database.MaxOpenConns)
}

func TestGetEnv(t *testing.T) {
	key := "TEST_ENV_VAR"
	os.Setenv(key, "value")
	defer os.Unsetenv(key)

	assert.Equal(t, "value", getEnv(key, "default"))
	assert.Equal(t, "default", getEnv("NON_EXISTENT", "default"))
}

func TestGetEnvBool(t *testing.T) {
	key := "TEST_BOOL_VAR"

	os.Setenv(key, "true")
	assert.True(t, getEnvBool(key, false))

	os.Setenv(key, "false")
	assert.False(t, getEnvBool(key, true))

	os.Setenv(key, "invalid")
	assert.True(t, getEnvBool(key, true))

	os.Unsetenv(key)
	assert.True(t, getEnvBool(key, true))
}

func TestGetEnvInt(t *testing.T) {
	key := "TEST_INT_VAR"

	os.Setenv(key, "123")
	assert.Equal(t, 123, getEnvInt(key, 0))

	os.Setenv(key, "invalid")
	assert.Equal(t, 10, getEnvInt(key, 10))

	os.Unsetenv(key)
	assert.Equal(t, 10, getEnvInt(key, 10))
}

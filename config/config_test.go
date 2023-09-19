package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoadConfig(t *testing.T) {
	// Test that LoadConfig returns a non-nil *config
	c := LoadConfig()
	require.NotNil(t, c, "LoadConfig should return a non-nil *config")

	// Test that each field in the config struct is set correctly
	expected := &config{
		JWT_KEY:        "test_jwt_key",
		DB1_USERNAME:   "test_db1_username",
		DB1_PASSWORD:   "test_db1_password",
		DB1_HOST:       "test_db1_host",
		DB1_PORT:       "test_db1_port",
		DB1_DATABASE:   "test_db1_database",
		REDIS_USERNAME: "test_redis_username",
		REDIS_PASSWORD: "test_redis_password",
		REDIS_HOST:     "test_redis_host",
		REDIS_PORT:     "test_redis_port",
		REDIS_DATABASE: "test_redis_database",
		DEBUG:          1,
	}
	require.Equal(t, expected, c, "config should be %v, but got %v", expected, c)
}

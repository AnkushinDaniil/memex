package config

import (
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

// Config holds all configuration for the Memex MCP server
type Config struct {
	// Database path for SQLite storage
	DatabasePath string

	// Server mode: "local" or "cloud"
	Mode string

	// Project detection: "git" or "cwd"
	ProjectDetection string

	// Log level: "debug", "info", "warn", "error"
	LogLevel string
}

// Load reads configuration from environment variables with sensible defaults
func Load() *Config {
	// Load .env file if it exists (ignore errors, file is optional)
	//nolint:errcheck // .env file is optional, error can be safely ignored
	godotenv.Load()

	cfg := &Config{
		DatabasePath:     getEnv("MEMEX_DATABASE_PATH", defaultDatabasePath()),
		Mode:             getEnv("MEMEX_MODE", "local"),
		ProjectDetection: getEnv("MEMEX_PROJECT_DETECTION", "git"),
		LogLevel:         getEnv("MEMEX_LOG_LEVEL", "info"),
	}

	return cfg
}

// getEnv retrieves an environment variable or returns the default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// defaultDatabasePath returns the default SQLite database path
func defaultDatabasePath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return ".memex/memex.db"
	}
	return filepath.Join(home, ".memex", "memex.db")
}

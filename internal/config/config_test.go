package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestConfigDefaults(t *testing.T) {
	// Clear any existing env vars
	if err := os.Unsetenv("MEMEX_DATABASE_PATH"); err != nil {
		t.Fatalf("Failed to unset MEMEX_DATABASE_PATH: %v", err)
	}
	if err := os.Unsetenv("MEMEX_MODE"); err != nil {
		t.Fatalf("Failed to unset MEMEX_MODE: %v", err)
	}
	if err := os.Unsetenv("MEMEX_PROJECT_DETECTION"); err != nil {
		t.Fatalf("Failed to unset MEMEX_PROJECT_DETECTION: %v", err)
	}
	if err := os.Unsetenv("MEMEX_LOG_LEVEL"); err != nil {
		t.Fatalf("Failed to unset MEMEX_LOG_LEVEL: %v", err)
	}

	cfg := Load()

	// Check defaults
	if cfg.Mode != "local" {
		t.Errorf("Expected Mode='local', got '%s'", cfg.Mode)
	}

	if cfg.ProjectDetection != "git" {
		t.Errorf("Expected ProjectDetection='git', got '%s'", cfg.ProjectDetection)
	}

	if cfg.LogLevel != "info" {
		t.Errorf("Expected LogLevel='info', got '%s'", cfg.LogLevel)
	}

	// DatabasePath should be ~/.memex/memex.db
	home, err := os.UserHomeDir()
	if err != nil {
		t.Fatalf("Failed to get user home dir: %v", err)
	}
	expectedPath := filepath.Join(home, ".memex", "memex.db")
	if cfg.DatabasePath != expectedPath {
		t.Errorf("Expected DatabasePath='%s', got '%s'", expectedPath, cfg.DatabasePath)
	}
}

func TestConfigFromEnv(t *testing.T) {
	// Set env vars
	if err := os.Setenv("MEMEX_DATABASE_PATH", "/tmp/test.db"); err != nil {
		t.Fatalf("Failed to set MEMEX_DATABASE_PATH: %v", err)
	}
	if err := os.Setenv("MEMEX_MODE", "cloud"); err != nil {
		t.Fatalf("Failed to set MEMEX_MODE: %v", err)
	}
	if err := os.Setenv("MEMEX_PROJECT_DETECTION", "cwd"); err != nil {
		t.Fatalf("Failed to set MEMEX_PROJECT_DETECTION: %v", err)
	}
	if err := os.Setenv("MEMEX_LOG_LEVEL", "debug"); err != nil {
		t.Fatalf("Failed to set MEMEX_LOG_LEVEL: %v", err)
	}

	cfg := Load()

	// Check env overrides
	if cfg.DatabasePath != "/tmp/test.db" {
		t.Errorf("Expected DatabasePath='/tmp/test.db', got '%s'", cfg.DatabasePath)
	}

	if cfg.Mode != "cloud" {
		t.Errorf("Expected Mode='cloud', got '%s'", cfg.Mode)
	}

	if cfg.ProjectDetection != "cwd" {
		t.Errorf("Expected ProjectDetection='cwd', got '%s'", cfg.ProjectDetection)
	}

	if cfg.LogLevel != "debug" {
		t.Errorf("Expected LogLevel='debug', got '%s'", cfg.LogLevel)
	}

	// Clean up
	if err := os.Unsetenv("MEMEX_DATABASE_PATH"); err != nil {
		t.Errorf("Failed to unset MEMEX_DATABASE_PATH: %v", err)
	}
	if err := os.Unsetenv("MEMEX_MODE"); err != nil {
		t.Errorf("Failed to unset MEMEX_MODE: %v", err)
	}
	if err := os.Unsetenv("MEMEX_PROJECT_DETECTION"); err != nil {
		t.Errorf("Failed to unset MEMEX_PROJECT_DETECTION: %v", err)
	}
	if err := os.Unsetenv("MEMEX_LOG_LEVEL"); err != nil {
		t.Errorf("Failed to unset MEMEX_LOG_LEVEL: %v", err)
	}
}

# Subtask 05: Configuration & CLI

**Business Value**: Users can configure the server and run it with sensible defaults, no manual setup required.

**Dependencies**: 01, 02, 03, 04

---

## Implementation Plan

### Files to Create

| File | Purpose |
|------|---------|
| internal/config/config.go | Configuration loading |
| internal/config/defaults.go | Default values |

### Files to Modify

| File | Changes | Why |
|------|---------|-----|
| cmd/memex/main.go | Add CLI flags, config loading | User configuration |
| Makefile | Add install target | Easy installation |

### Steps

1. Create internal/config/config.go:
   - Config struct with all settings
   - Load() function reads env vars and defaults
   - Validate() checks required settings

2. Create internal/config/defaults.go:
   - Default database path: ~/.memex/memex.db
   - Default log level: info
   - Default project detection: git

3. Update cmd/memex/main.go:
   - Parse CLI flags (--db-path, --log-level, --version)
   - Load config from env + flags
   - Initialize logging
   - Print startup info to stderr (stdout reserved for MCP)

4. Update Makefile:
   - `build-memex`: Build binary
   - `install-memex`: Install to ~/.local/bin
   - `test-memex`: Run tests

5. Create default directories on first run:
   - ~/.memex/ for database and config
   - Proper permissions (700)

### Patterns & Hints

```go
// Config struct
type Config struct {
    DatabasePath string `env:"MEMEX_DB_PATH"`
    LogLevel     string `env:"MEMEX_LOG_LEVEL"`
    ProjectMode  string `env:"MEMEX_PROJECT_MODE"` // git, cwd, manual
}

func Load() (*Config, error) {
    cfg := &Config{
        DatabasePath: defaultDBPath(),
        LogLevel:     "info",
        ProjectMode:  "git",
    }

    // Override from environment
    if v := os.Getenv("MEMEX_DB_PATH"); v != "" {
        cfg.DatabasePath = v
    }

    return cfg, cfg.Validate()
}

func defaultDBPath() string {
    home, _ := os.UserHomeDir()
    return filepath.Join(home, ".memex", "memex.db")
}
```

```go
// Main with CLI flags
func main() {
    var (
        dbPath   = flag.String("db", "", "Database path")
        logLevel = flag.String("log-level", "info", "Log level")
        version  = flag.Bool("version", false, "Print version")
    )
    flag.Parse()

    if *version {
        fmt.Println("memex v0.1.0")
        os.Exit(0)
    }

    cfg, err := config.Load()
    if err != nil {
        log.Fatal(err)
    }

    // CLI flags override config
    if *dbPath != "" {
        cfg.DatabasePath = *dbPath
    }

    // Log to stderr (stdout is for MCP)
    log.SetOutput(os.Stderr)

    // ... start server
}
```

```makefile
# Makefile additions
.PHONY: build-memex install-memex test-memex

build-memex:
	go build -o bin/memex ./cmd/memex

install-memex: build-memex
	mkdir -p ~/.local/bin
	cp bin/memex ~/.local/bin/
	@echo "Installed to ~/.local/bin/memex"

test-memex:
	go test -v ./cmd/memex/... ./internal/...
```

---

## Testing

**Test File**: `internal/config/config_test.go`

| Test Name | Description | Expected |
|-----------|-------------|----------|
| TestLoadDefaults | Load with no env | Default values |
| TestLoadFromEnv | Set MEMEX_DB_PATH | Env value used |
| TestValidateEmpty | Empty database path | Error |
| TestDefaultDBPath | Check default path | ~/.memex/memex.db |

**Test File**: `cmd/memex/main_test.go`

| Test Name | Description | Expected |
|-----------|-------------|----------|
| TestVersionFlag | Run with --version | Prints version, exits 0 |
| TestDBPathFlag | Run with --db /tmp/test.db | Uses specified path |

**Mocking Strategy**: Use environment variable manipulation for config tests.

---

## Acceptance Criteria

- [ ] Server runs with zero configuration
- [ ] Database created at ~/.memex/memex.db by default
- [ ] MEMEX_DB_PATH env var overrides default
- [ ] --db flag overrides env and default
- [ ] --version prints version and exits
- [ ] Logs go to stderr, MCP to stdout
- [ ] `make install-memex` installs to ~/.local/bin

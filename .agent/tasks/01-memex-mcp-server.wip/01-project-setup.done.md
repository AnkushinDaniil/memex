# Subtask 01: Project Setup

**Business Value**: Establish the Go project foundation with correct module structure, dependencies, and build tooling.

**Dependencies**: None

---

## Implementation Plan

### Files to Create

| File | Purpose |
|------|---------|
| cmd/memex/main.go | MCP server entry point |
| internal/mcp/server.go | MCP protocol handling |
| internal/memory/service.go | Memory business logic |
| internal/storage/sqlite.go | SQLite storage implementation |
| internal/config/config.go | Configuration management |
| go.mod | Go module definition |
| go.sum | Dependency checksums |
| Makefile | Build commands (update existing) |

### Directory Structure

```
cmd/
  memex/
    main.go           # Entry point
internal/
  mcp/
    server.go         # MCP protocol
    tools.go          # Tool definitions
    types.go          # MCP types
  memory/
    service.go        # Memory CRUD
    types.go          # Memory types
  storage/
    sqlite.go         # SQLite + FTS5
    interface.go      # Storage interface
  config/
    config.go         # Configuration
```

### Steps

1. Update go.mod with required dependencies:
   - `github.com/mattn/go-sqlite3` (SQLite with FTS5)
   - `github.com/google/uuid` (UUID generation)
   - `github.com/joho/godotenv` (env loading)

2. Create cmd/memex/main.go with minimal MCP server startup

3. Create internal/mcp/types.go with MCP JSON-RPC types

4. Create internal/memory/types.go with Memory struct

5. Create internal/storage/interface.go with Storage interface

6. Create internal/config/config.go with Config struct

7. Update Makefile with memex build target

### Patterns & Hints

```go
// MCP servers communicate via stdio JSON-RPC
// Entry point reads from stdin, writes to stdout

func main() {
    cfg := config.Load()
    storage := sqlite.New(cfg.DatabasePath)
    memoryService := memory.NewService(storage)
    mcpServer := mcp.NewServer(memoryService)

    // MCP uses stdio
    mcpServer.ServeStdio()
}
```

```go
// Memory type
type Memory struct {
    ID          string    `json:"id"`
    ProjectID   string    `json:"project_id"`
    Content     string    `json:"content"`
    Tags        []string  `json:"tags"`
    Priority    string    `json:"priority"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}
```

---

## Testing

**Test File**: `internal/config/config_test.go`

| Test Name | Description | Expected |
|-----------|-------------|----------|
| TestConfigDefaults | Load config with no env vars | Default values applied |
| TestConfigFromEnv | Load config from environment | Env values override defaults |

**Mocking Strategy**: No mocking needed - unit tests for config parsing.

**Build Verification**:
```bash
go build ./cmd/memex
# Should produce memex binary without errors
```

---

## Acceptance Criteria

- [ ] `go build ./cmd/memex` succeeds
- [ ] `go test ./...` passes
- [ ] Directory structure matches plan
- [ ] All type definitions compile
- [ ] Makefile has `build-memex` target

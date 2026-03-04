# Memex MCP Server

## Project Overview

This is a **Model Context Protocol (MCP) server** that provides persistent memory storage for AI assistants. Memex enables long-term memory capabilities through SQLite with full-text search (FTS5).

**Product Type**: MCP Server / Developer Tool
**Language**: Go 1.26
**Database**: SQLite with FTS5
**Protocol**: MCP 2024-11-05 (JSON-RPC 2.0 over stdio)
**Repository**: https://github.com/AnkushinDaniil/memex

## Core Architecture

```
Claude/AI Assistant
       │
       ▼
   MCP Client
       │
       ▼
  Memex Server (stdio)
       │
       ├─▶ Memory Service
       │      ├─ Create
       │      ├─ Get
       │      ├─ Search
       │      ├─ List
       │      └─ Delete
       │
       └─▶ SQLite Storage (FTS5)
```

## Technology Stack

- **Go 1.26**: Core implementation
- **SQLite + FTS5**: Full-text search storage
- **MCP Protocol**: Model Context Protocol (2024-11-05)
- **JSON-RPC 2.0**: Communication protocol over stdio
- **golangci-lint v2.10.1**: Code quality (25+ linters)
- **govulncheck**: Vulnerability scanning
- **Trivy**: Container security

## Development Workflow

### Local Development
```bash
# Install dependencies
go mod download

# Build
make build

# Run
make run

# Run with hot reload
make dev
```

### Code Quality
```bash
# Run linter
make lint

# Format code
make fmt

# Run tests
make test

# Run tests with coverage
make test-coverage
```

### Security
```bash
# Vulnerability scan
make vuln

# Trivy scan
make trivy

# Full security audit
make security
```

### CI Pipeline
```bash
# Full CI (lint, test, coverage, security)
make ci

# Quick CI (no security scans)
make ci-quick
```

## MCP Tools (Planned)

The server will provide these MCP tools:

- **memory_create**: Store a new memory with content, tags, and project ID
- **memory_get**: Retrieve a specific memory by ID
- **memory_search**: Full-text search across all memories
- **memory_list**: List recent memories with optional filtering
- **memory_delete**: Remove a memory by ID

## Configuration

Environment variables:
```bash
MEMEX_DB_PATH    # Database path (default: ./memex.db)
MEMEX_MODE       # Mode: development or production
```

## CI/CD

GitHub Actions workflows:

### CI Workflow (.github/workflows/ci.yml)
- Triggers: Push to main/develop, PRs
- Linting with golangci-lint v2.10.1
- Tests with race detector
- Coverage threshold (80% warning)
- Security scanning (govulncheck + Trivy)
- Builds memex binary

### Release Workflow (.github/workflows/release.yml)
- Triggers: Push to main branch
- Quality gate (lint, test, security)
- Semantic versioning (conventional commits)
- Docker image build and push to ghcr.io
- Automated changelog generation

### Security Workflow (.github/workflows/security.yml)
- Triggers: Daily at 2am UTC + manual
- Vulnerability scanning
- Container image scanning
- Dependency review

## Commit Convention

```
<type>(<scope>): <description>

Types: feat, fix, perf, refactor, docs, test, chore
Breaking: Add "!" after type (e.g., feat!:)
```

Version bumps:
- `feat` → MINOR
- `fix`, `perf` → PATCH
- Breaking change (`!`) → MAJOR

## Project Structure

```
.
├── cmd/
│   └── memex/          # MCP server entry point
├── internal/
│   ├── config/         # Configuration
│   ├── mcp/            # MCP protocol implementation
│   ├── memory/         # Memory domain models
│   └── storage/        # SQLite storage layer
├── .github/workflows/  # CI/CD workflows
├── Dockerfile          # Docker build
├── Makefile            # Build automation
└── go.mod              # Go dependencies
```

## Implementation Status

Current status: **Foundation Complete**
- ✅ Project structure
- ✅ Configuration system
- ✅ MCP protocol scaffolding
- ✅ Storage interface
- ✅ CI/CD pipelines
- ⏳ Storage implementation (stub)
- ⏳ MCP tools (stub)
- ⏳ Full-text search (stub)

## Development Guidelines

### Code Quality
- All code must pass golangci-lint (25+ linters)
- Coverage threshold: 80% (warning, not blocking)
- No HIGH/CRITICAL security vulnerabilities
- All errors must be handled

### Security
- Never expose credentials
- Never log sensitive data
- Parameterize all SQL queries
- Validate all user input

### Testing
- Unit tests for all business logic
- Race detector enabled
- Coverage reports generated
- Integration tests for storage

### Git Workflow
- Feature branches from main
- Conventional commits required
- All commits must pass CI
- Semantic versioning automated

## Quick Reference

### Build Commands
- `make build` - Build binary
- `make run` - Run server
- `make test` - Run tests
- `make lint` - Run linter
- `make ci` - Full CI pipeline

### Docker
- `make docker-build` - Build image
- `make docker-run` - Run container

### Files
- Entry point: `cmd/memex/main.go`
- MCP server: `internal/mcp/server.go`
- Storage: `internal/storage/sqlite.go`
- CI: `.github/workflows/ci.yml`
- Release: `.github/workflows/release.yml`

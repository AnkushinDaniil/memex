# Memex MCP Server

A Model Context Protocol (MCP) server for persistent memory storage and retrieval. Memex provides AI assistants with long-term memory capabilities through SQLite with full-text search.

## Features

- **MCP Protocol**: Implements the Model Context Protocol (2024-11-05) for seamless integration with Claude and other AI assistants
- **Full-Text Search**: SQLite with FTS5 for fast, efficient memory retrieval
- **Persistent Storage**: Durable memory storage across sessions
- **Project Isolation**: Organize memories by project
- **Tag System**: Categorize and filter memories with tags
- **Stdio Communication**: Standard JSON-RPC 2.0 protocol over stdin/stdout

## Architecture

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

## Quick Start

```bash
# Install dependencies
go mod download

# Build the server
make build

# Run the server
./bin/memex

# Or use during development
make run
```

## Configuration

The server can be configured via environment variables:

```bash
# Database path (default: ./memex.db)
MEMEX_DB_PATH=/path/to/database.db

# Mode: development or production (default: development)
MEMEX_MODE=production
```

Or use a `.env` file:

```env
MEMEX_DB_PATH=./data/memex.db
MEMEX_MODE=development
```

## MCP Tools

Memex provides the following MCP tools:

- **memory_create**: Store a new memory
- **memory_get**: Retrieve a specific memory by ID
- **memory_search**: Full-text search across all memories
- **memory_list**: List recent memories with optional filtering
- **memory_delete**: Remove a memory

## Development

```bash
# Run tests
make test

# Run linter
make lint

# Format code
make fmt

# Run all checks
make ci
```

## Technology Stack

- **Language**: Go 1.26
- **Database**: SQLite with FTS5
- **Protocol**: MCP (Model Context Protocol)
- **Communication**: JSON-RPC 2.0 over stdio

## CI/CD

The project uses GitHub Actions for continuous integration:

- **Linting**: golangci-lint v2.10.1
- **Testing**: Unit tests with race detector
- **Security**: govulncheck + Trivy scanning
- **Coverage**: 80% threshold (warning)
- **Release**: Automated semantic versioning

## License

MIT

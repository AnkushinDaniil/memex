# Task 01: Memex MCP Server

## Source References

| Source | Path | Purpose |
|--------|------|---------|
| Product Spec | specs/features/agent-memory-infrastructure.yaml | Features, pricing, metrics |
| System Architecture | specs/architectural/memex-system.yaml | Component design, deployment |
| MCP Server Spec | specs/architectural/memex-mcp-server.yaml | MCP tools, search strategy |
| API Contracts | specs/architectural/memex-api.yaml | REST API endpoints |
| Data Models | specs/architectural/memex-data-models.yaml | Database schema |

## Objective

Build the Memex MCP Server - a persistent memory system for Claude Code that stores and retrieves developer memories using full-text search. The server provides MCP tools that Claude Code can call directly, with no external LLM dependencies.

The system supports two modes: local (SQLite + FTS5, free forever) and cloud (PostgreSQL + tsvector, paid tiers). Both modes expose identical MCP tools, allowing seamless upgrade paths.

Claude Code's native intelligence handles semantic understanding - the MCP server only provides storage and full-text search. This eliminates embedding costs, API dependencies, and privacy concerns.

## Architecture

### Mental Model

Memex is a memory layer that sits between Claude Code and persistent storage. When a developer says "remember this", the MCP server stores it. When Claude needs context, it calls `memex_recall` to get candidate memories, then uses its own understanding to pick relevant ones.

The architecture is intentionally simple: Go backend with Chi router, SQLite/PostgreSQL storage with full-text search, and MCP protocol for Claude Code integration. No embeddings, no external AI, no complexity.

### Component Responsibilities

| Component | Role | Subtask |
|-----------|------|---------|
| MCP Server (Go) | Handle MCP protocol, expose tools | 01, 02 |
| Memory Service | CRUD operations for memories | 03 |
| Search Service | Full-text search (FTS5/tsvector) | 04 |
| Storage Layer | SQLite (local) / PostgreSQL (cloud) | 03, 04 |
| CLI/Config | Configuration and local setup | 05 |

### Core Use Cases

1. **Remember**: Developer tells Claude to remember a decision → stored via `memex_remember`
2. **Recall**: Claude searches for relevant context → `memex_recall` returns candidates
3. **Forget**: Developer removes outdated memories → `memex_forget` deletes
4. **List**: Developer reviews stored memories → `memex_list` returns recent items

### Data Flow

```
Developer: "Remember we use JWT for auth"
    ↓
Claude Code calls: memex_remember(content="We use JWT for auth", tags=["auth"])
    ↓
MCP Server receives tool call via stdio
    ↓
Memory Service creates record with FTS index
    ↓
Returns: {memory_id: "mem_xxx", created_at: "..."}

Later...

Developer: "How do we handle authentication?"
    ↓
Claude Code calls: memex_recall(query="authentication")
    ↓
Search Service runs: SELECT * FROM memories WHERE search MATCH 'authentication'
    ↓
Returns 5 candidate memories
    ↓
Claude reads them, picks relevant one: "We use JWT for auth"
    ↓
Claude incorporates in response naturally
```

## Subtasks

### Dependency Graph

```
01 → 02 → 03 → 04 → 05 → 06
```

All subtasks are sequential - each builds on the previous.

### Index

| # | Subtask | Status | Business Value | File |
|---|---------|--------|----------------|------|
| 01 | Project Setup | done | Establish Go project with correct structure and dependencies | [01-project-setup.done.md](01-project-setup.done.md) |
| 02 | MCP Protocol Handler | todo | Claude Code can communicate with server via stdio | [02-mcp-protocol.todo.md](02-mcp-protocol.todo.md) |
| 03 | Memory Storage | todo | Memories persist to SQLite with full-text indexing | [03-memory-storage.todo.md](03-memory-storage.todo.md) |
| 04 | MCP Tools Implementation | todo | All 6 MCP tools work end-to-end | [04-mcp-tools.todo.md](04-mcp-tools.todo.md) |
| 05 | Configuration & CLI | todo | Users can configure and run the server | [05-config-cli.todo.md](05-config-cli.todo.md) |
| 06 | Claude Code Integration | todo | Server works as Claude Code MCP server | [06-claude-code-integration.todo.md](06-claude-code-integration.todo.md) |

## Verification

After all subtasks complete:

1. Add MCP server to Claude Code settings
2. Start Claude Code session
3. Run: "Remember that we use PostgreSQL for the database"
4. Start new session
5. Ask: "What database do we use?"
6. Claude should retrieve and use the memory

## Constraints & Gotchas

- **No external LLM**: Search is full-text only. Claude handles semantic understanding.
- **MCP stdio**: Communication via stdin/stdout JSON-RPC, not HTTP
- **SQLite first**: Local mode is the default, cloud is upgrade path
- **FTS5 required**: SQLite must be compiled with FTS5 extension
- **Project detection**: Use git root or cwd as project identifier

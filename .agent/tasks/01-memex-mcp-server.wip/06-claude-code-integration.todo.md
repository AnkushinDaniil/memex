# Subtask 06: Claude Code Integration

**Business Value**: Server works as a Claude Code MCP server, enabling persistent memory in real Claude Code sessions.

**Dependencies**: 01, 02, 03, 04, 05

---

## Implementation Plan

### Files to Create

| File | Purpose |
|------|---------|
| .claude/settings.local.json.example | Example Claude Code config |
| README.md | Installation and usage docs |

### Files to Modify

| File | Changes | Why |
|------|---------|-----|
| cmd/memex/main.go | Ensure proper MCP compliance | Claude Code compatibility |

### Steps

1. Verify MCP protocol compliance:
   - Test with Claude Code's MCP inspector
   - Ensure proper JSON-RPC framing
   - Handle all required MCP methods

2. Create example Claude Code settings:
   - mcpServers configuration
   - Point to memex binary
   - Document environment variables

3. Create README.md with:
   - Quick start guide
   - Installation instructions
   - Configuration options
   - Usage examples

4. Test end-to-end with Claude Code:
   - Add to settings
   - Start session
   - Use remember/recall
   - Verify persistence

5. Handle edge cases:
   - Server crash recovery
   - Database locked errors
   - Large memory content

### Patterns & Hints

```json
// .claude/settings.local.json.example
{
  "mcpServers": {
    "memex": {
      "command": "memex",
      "args": [],
      "env": {
        "MEMEX_DB_PATH": "~/.memex/memex.db",
        "MEMEX_LOG_LEVEL": "info"
      }
    }
  }
}
```

```markdown
// README.md structure
# Memex - Persistent Memory for Claude Code

## Quick Start
1. Install: `go install github.com/your-org/memex/cmd/memex@latest`
2. Add to Claude Code settings
3. Start using!

## Usage
- "Remember that we use PostgreSQL" → stores memory
- "What database do we use?" → Claude recalls memory

## Configuration
| Env Var | Default | Description |
|---------|---------|-------------|
| MEMEX_DB_PATH | ~/.memex/memex.db | Database location |
```

```go
// Ensure MCP compliance in main.go
func main() {
    // ... setup ...

    // Log startup to stderr (not stdout!)
    log.SetOutput(os.Stderr)
    log.Printf("memex v%s starting", version)
    log.Printf("database: %s", cfg.DatabasePath)

    // Serve MCP on stdio
    if err := server.ServeStdio(); err != nil {
        log.Fatalf("server error: %v", err)
    }
}
```

---

## Testing

**Manual Test Script**: `scripts/test-integration.sh`

```bash
#!/bin/bash
# Test MCP server manually

# Build
go build -o bin/memex ./cmd/memex

# Test initialize
echo '{"jsonrpc":"2.0","id":1,"method":"initialize","params":{}}' | ./bin/memex

# Test tools/list
echo '{"jsonrpc":"2.0","id":2,"method":"tools/list"}' | ./bin/memex

# Test remember
echo '{"jsonrpc":"2.0","id":3,"method":"tools/call","params":{"name":"memex_remember","arguments":{"content":"Test memory"}}}' | ./bin/memex
```

| Test | Description | Expected |
|------|-------------|----------|
| MCP Initialize | Send initialize request | Valid capabilities response |
| Tools List | Request tool list | 6 tools with schemas |
| Remember Flow | Store then recall | Memory persists |
| Cross-Session | Stop server, restart, recall | Memory still there |
| Claude Code E2E | Real Claude Code session | Works naturally |

**Mocking Strategy**: No mocking - this is integration/E2E testing.

---

## Acceptance Criteria

- [ ] Server listed in Claude Code MCP servers
- [ ] Initialize handshake completes
- [ ] All tools visible to Claude Code
- [ ] "Remember X" stores memory
- [ ] "What was X?" retrieves memory
- [ ] Memory persists across sessions
- [ ] README has clear setup instructions
- [ ] Example settings file works as-is

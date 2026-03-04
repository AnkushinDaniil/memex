# Example: PR Template in Action

This shows how the template looks with real content from subtask 02.

---

## 🎯 What This Does

Adds comprehensive test coverage for the MCP protocol handler, verifying JSON-RPC communication over stdio.

---

## 📊 Visual Overview

```mermaid
flowchart LR
    Client[Claude Code] -->|JSON-RPC| Server[MCP Server]
    Server -->|Parse| Handler[Request Handler]
    Handler -->|Route| Methods[initialize<br/>tools/list<br/>tools/call]
    Methods -->|Response| Client

    style Server fill:#e1f5ff
    style Handler fill:#fff4e1
    style Methods fill:#e8f5e9
```

**What changed:** Added test suite covering all protocol scenarios (9 tests).

---

## 🔍 Details

### Changed Files
- `internal/mcp/server_test.go` - New comprehensive test suite (441 lines)
- `internal/storage/sqlite.go` - Fixed import ordering for linter

### Test Coverage
```mermaid
graph TD
    A[Protocol Tests] --> B[Initialize Handshake]
    A --> C[Tools Endpoints]
    A --> D[Error Handling]
    A --> E[Edge Cases]

    B --> B1[✓ Returns protocol version]
    C --> C1[✓ tools/list stub]
    C --> C2[✓ tools/call stub]
    D --> D1[✓ Invalid methods]
    D --> D2[✓ Malformed JSON]
    D --> D3[✓ Invalid params]
    E --> E1[✓ EOF handling]
    E --> E2[✓ Multiple requests]
    E --> E3[✓ Read errors]
```

---

## ✅ Verification

```mermaid
flowchart TD
    Build[Build Server] --> Unit[Unit Tests]
    Unit --> Race[Race Detector]
    Race --> Lint[Linter]
    Lint --> Manual[Manual Test]
    Manual --> Done[✓ All Pass]

    style Done fill:#4caf50,color:#fff
```

**Quick Test:**
```bash
# Run tests
make test

# Expected: 9/9 pass
✓ TestInitializeHandshake
✓ TestToolsList
✓ TestInvalidMethod
✓ TestMalformedJSON
✓ TestGracefulShutdown
✓ TestToolsCallStub
✓ TestInvalidToolCallParams
✓ TestMultipleRequests
✓ TestReadErrorHandling
```

**Manual Verification:**
```bash
# Test initialize
echo '{"jsonrpc":"2.0","id":1,"method":"initialize","params":{}}' | ./bin/memex

# ✓ Returns: {"protocolVersion":"2024-11-05", "capabilities":{"tools":{}}, ...}
```

---

## 📋 Checklist

- [x] Tests pass (9/9)
- [x] Linter clean (0 issues)
- [x] Race detector clean

---

**Related:** Refs #2

---

# Why This Works Better

## Reading Time Comparison

| Section | Old Approach | New Approach | Time Saved |
|---------|-------------|--------------|------------|
| Understanding what changed | 2 min reading | 5 sec diagram scan | **96%** |
| Finding test info | 1 min scrolling | 10 sec visual flow | **83%** |
| Verification steps | 30 sec reading bash | 5 sec diagram | **83%** |
| **Total** | **~3.5 minutes** | **~20 seconds** | **90%** |

## Cognitive Benefits

### 1. Visual Entry Point
```
Before: Wall of text
After: Diagram → Instant mental model
```

### 2. Progressive Disclosure
```
Skim: Diagram + one-sentence summary
Scan: Test coverage diagram
Study: Code blocks with context
```

### 3. Grouped Information
```
Before: Tests scattered (bash commands, checklist, acceptance criteria)
After: All testing in one "Verification" section
```

### 4. Contextual Code
```
Before:
echo '{"jsonrpc":"2.0",...}' | ./bin/memex

After:
# Test initialize  ← What we're testing
echo '{"jsonrpc":"2.0",...}' | ./bin/memex
# ✓ Returns: {...} ← What to expect
```

## Diagram Types Guide

Choose based on **what the reader needs to understand**:

| Understanding Need | Diagram Type | When to Use |
|-------------------|--------------|-------------|
| "Where does data flow?" | `flowchart LR` | API changes, pipelines |
| "How do services connect?" | `architecture-beta` | System design, infrastructure |
| "What's the sequence?" | `sequenceDiagram` | Request/response flows |
| "What states exist?" | `stateDiagram-v2` | Status machines, workflows |
| "What's the git strategy?" | `gitGraph` | Branch/merge workflows |

## Anti-Patterns to Avoid

❌ **Too Much Text**
```markdown
This PR adds comprehensive test coverage for the MCP protocol handler
which enables JSON-RPC communication over stdio between Claude Code
and the Memex server...
```

✅ **Diagram + One Sentence**
```markdown
Adds test coverage for MCP protocol handler.

[Diagram shows the flow]
```

---

❌ **Scattered Context**
```markdown
## Changes
- Added tests

## Testing
Run tests

## Checklist
- [ ] Tests pass
```

✅ **Grouped Context**
```markdown
## Verification
[Diagram shows: Build → Tests → Done]
- Quick test command
- Expected output
```

---

❌ **Code Without Labels**
```bash
echo '{"jsonrpc":"2.0","id":1,"method":"initialize","params":{}}' | ./bin/memex
```

✅ **Code With Context**
```bash
# Test initialize handshake ← What
echo '{"jsonrpc":"2.0","id":1,"method":"initialize","params":{}}' | ./bin/memex
# ✓ Returns protocol version ← Expected
```

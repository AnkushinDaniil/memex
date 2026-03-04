# Subtask 04: MCP Tools Implementation

**Business Value**: All 6 MCP tools work end-to-end, enabling Claude Code to store, search, and manage memories.

**Dependencies**: 01, 02, 03

---

## Implementation Plan

### Files to Create

| File | Purpose |
|------|---------|
| internal/mcp/tools.go | Tool definitions and schemas |
| internal/mcp/tools_impl.go | Tool execution logic |

### Files to Modify

| File | Changes | Why |
|------|---------|-----|
| internal/mcp/handler.go | Wire tools to handler | Route tool calls |

### Steps

1. Create internal/mcp/tools.go with tool definitions:
   - Define all 6 tools with JSON schemas
   - memex_remember, memex_recall, memex_forget
   - memex_list, memex_update, memex_stats

2. Create internal/mcp/tools_impl.go with implementations:
   - Each tool function takes params, returns result
   - Validate inputs, call memory service, format output

3. Update handler.go to dispatch tool calls:
   - Parse tool name and arguments
   - Call appropriate implementation
   - Return result or error

4. Implement each tool:

**memex_remember**: Store a new memory
- Required: content
- Optional: tags, priority, project
- Returns: memory_id, created_at

**memex_recall**: Search memories
- Required: query
- Optional: limit, tags, project, since
- Returns: array of memories with relevance

**memex_forget**: Delete a memory
- Required: memory_id
- Returns: deleted (bool)

**memex_list**: List recent memories
- Optional: limit, project, tags
- Returns: array of memories

**memex_update**: Update existing memory
- Required: memory_id
- Optional: content, tags, priority
- Returns: updated memory

**memex_stats**: Get statistics
- Optional: project
- Returns: counts, storage size

### Patterns & Hints

```go
// Tool definition for MCP
type Tool struct {
    Name        string          `json:"name"`
    Description string          `json:"description"`
    InputSchema json.RawMessage `json:"inputSchema"`
}

var tools = []Tool{
    {
        Name:        "memex_remember",
        Description: "Store a memory for later recall",
        InputSchema: json.RawMessage(`{
            "type": "object",
            "properties": {
                "content": {
                    "type": "string",
                    "description": "The content to remember"
                },
                "tags": {
                    "type": "array",
                    "items": {"type": "string"},
                    "description": "Tags for categorization"
                },
                "priority": {
                    "type": "string",
                    "enum": ["low", "normal", "high"],
                    "default": "normal"
                }
            },
            "required": ["content"]
        }`),
    },
    // ... other tools
}
```

```go
// Tool implementation
func (s *Server) executeRemember(ctx context.Context, params map[string]interface{}) (interface{}, error) {
    content, ok := params["content"].(string)
    if !ok || content == "" {
        return nil, fmt.Errorf("content is required")
    }

    mem := &memory.Memory{
        Content:  content,
        Tags:     extractTags(params),
        Priority: extractPriority(params),
    }

    if err := s.memoryService.Create(ctx, mem); err != nil {
        return nil, err
    }

    return map[string]interface{}{
        "memory_id":  mem.ID,
        "created_at": mem.CreatedAt,
    }, nil
}
```

---

## Testing

**Test File**: `internal/mcp/tools_test.go`

| Test Name | Description | Expected |
|-----------|-------------|----------|
| TestRememberBasic | Store simple memory | Memory created, ID returned |
| TestRememberWithTags | Store with tags | Tags persisted |
| TestRecallFound | Search existing | Memories returned |
| TestRecallNotFound | Search non-existent | Empty array |
| TestForgetExisting | Delete memory | deleted: true |
| TestForgetMissing | Delete non-existent | Error or deleted: false |
| TestListEmpty | List empty project | Empty array |
| TestListWithMemories | List after creates | Correct count |
| TestUpdateContent | Change content | Updated content returned |
| TestStats | Get statistics | Correct counts |

**Integration Test File**: `internal/mcp/integration_test.go`

| Test Name | Description | Expected |
|-----------|-------------|----------|
| TestRememberThenRecall | Create then search | Memory found |
| TestRememberUpdateRecall | Create, update, search | Updated content found |
| TestRememberForgetRecall | Create, delete, search | Memory not found |

**Mocking Strategy**: Use real SQLite in-memory for integration tests.

---

## Acceptance Criteria

- [ ] All 6 tools appear in tools/list response
- [ ] memex_remember stores and returns ID
- [ ] memex_recall searches and returns ranked results
- [ ] memex_forget deletes and confirms
- [ ] memex_list returns recent memories
- [ ] memex_update modifies existing memory
- [ ] memex_stats returns accurate counts
- [ ] Invalid inputs return proper errors

# Subtask 02: MCP Protocol Handler

**Business Value**: Claude Code can communicate with the Memex server using standard MCP protocol over stdio.

**Dependencies**: 01

---

## Implementation Plan

### Files to Create

| File | Purpose |
|------|---------|
| internal/mcp/server.go | Main MCP server implementation |
| internal/mcp/types.go | MCP JSON-RPC message types |
| internal/mcp/handler.go | Request routing and handling |

### Files to Modify

| File | Changes | Why |
|------|---------|-----|
| cmd/memex/main.go | Initialize and start MCP server | Wire up components |

### Steps

1. Create internal/mcp/types.go with MCP protocol types:
   - JSONRPCRequest, JSONRPCResponse
   - InitializeParams, InitializeResult
   - ToolCall, ToolResult
   - ServerCapabilities

2. Create internal/mcp/server.go with Server struct:
   - NewServer(memoryService) constructor
   - ServeStdio() - main loop reading stdin, writing stdout
   - handleRequest() - route to appropriate handler

3. Create internal/mcp/handler.go with handlers:
   - handleInitialize() - return server info and capabilities
   - handleToolsList() - return available tools
   - handleToolsCall() - dispatch to tool implementations

4. Implement the MCP handshake flow:
   - Client sends `initialize` with capabilities
   - Server responds with name, version, capabilities
   - Client sends `initialized` notification
   - Ready for tool calls

### Patterns & Hints

```go
// MCP JSON-RPC message structure
type JSONRPCRequest struct {
    JSONRPC string          `json:"jsonrpc"`
    ID      interface{}     `json:"id,omitempty"`
    Method  string          `json:"method"`
    Params  json.RawMessage `json:"params,omitempty"`
}

type JSONRPCResponse struct {
    JSONRPC string      `json:"jsonrpc"`
    ID      interface{} `json:"id,omitempty"`
    Result  interface{} `json:"result,omitempty"`
    Error   *RPCError   `json:"error,omitempty"`
}

// Server capabilities
type ServerCapabilities struct {
    Tools *ToolsCapability `json:"tools,omitempty"`
}

type ToolsCapability struct {
    ListChanged bool `json:"listChanged,omitempty"`
}
```

```go
// Main serve loop
func (s *Server) ServeStdio() error {
    reader := bufio.NewReader(os.Stdin)
    encoder := json.NewEncoder(os.Stdout)

    for {
        line, err := reader.ReadBytes('\n')
        if err == io.EOF {
            return nil
        }
        if err != nil {
            return err
        }

        var req JSONRPCRequest
        if err := json.Unmarshal(line, &req); err != nil {
            continue
        }

        resp := s.handleRequest(&req)
        if resp != nil {
            encoder.Encode(resp)
        }
    }
}
```

---

## Testing

**Test File**: `internal/mcp/server_test.go`

| Test Name | Description | Expected |
|-----------|-------------|----------|
| TestInitializeHandshake | Send initialize, receive capabilities | Valid response with tools capability |
| TestToolsList | Request tools/list | Returns 6 memex tools |
| TestInvalidMethod | Send unknown method | Error response with code -32601 |
| TestMalformedJSON | Send invalid JSON | Error response with code -32700 |

**Mocking Strategy**: Mock the memory service for isolated MCP protocol testing.

```go
type mockMemoryService struct{}

func (m *mockMemoryService) Create(ctx context.Context, mem *Memory) error {
    return nil
}
// ... other methods
```

---

## Acceptance Criteria

- [ ] Server starts and waits for stdin input
- [ ] Initialize handshake completes successfully
- [ ] tools/list returns all 6 Memex tools with schemas
- [ ] Invalid requests return proper JSON-RPC errors
- [ ] Server handles EOF gracefully (clean shutdown)

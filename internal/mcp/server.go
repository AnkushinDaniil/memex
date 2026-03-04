package mcp

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/AnkushinDaniil/memex/internal/memory"
)

// Server represents the MCP server
type Server struct {
	memoryService *memory.Service
	reader        *bufio.Reader
	writer        io.Writer
}

// NewServer creates a new MCP server instance
func NewServer(memoryService *memory.Service) *Server {
	return &Server{
		memoryService: memoryService,
		reader:        bufio.NewReader(os.Stdin),
		writer:        os.Stdout,
	}
}

// ServeStdio starts the MCP server, reading from stdin and writing to stdout
func (s *Server) ServeStdio() error {
	for {
		// Read line from stdin
		line, err := s.reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return fmt.Errorf("read error: %w", err)
		}

		// Parse JSON-RPC request
		var req JSONRPCRequest
		if err := json.Unmarshal(line, &req); err != nil {
			s.sendError(nil, ParseError, "parse error", err.Error())
			continue
		}

		// Handle request
		s.handleRequest(&req)
	}
}

// handleRequest processes a JSON-RPC request and sends a response
func (s *Server) handleRequest(req *JSONRPCRequest) {
	switch req.Method {
	case "initialize":
		s.handleInitialize(req)
	case "tools/list":
		s.handleListTools(req)
	case "tools/call":
		s.handleCallTool(req)
	default:
		s.sendError(req.ID, MethodNotFound, "method not found", req.Method)
	}
}

// handleInitialize handles the initialize method
func (s *Server) handleInitialize(req *JSONRPCRequest) {
	result := InitializeResult{
		ProtocolVersion: "2024-11-05",
		Capabilities: ServerCapabilities{
			Tools: &ToolsCapability{},
		},
		ServerInfo: ServerInfo{
			Name:    "memex",
			Version: "1.0.0",
		},
	}

	s.sendResult(req.ID, result)
}

// handleListTools handles the tools/list method
func (s *Server) handleListTools(req *JSONRPCRequest) {
	// Stub - will be implemented in subtask 04
	result := ListToolsResult{
		Tools: []Tool{},
	}

	s.sendResult(req.ID, result)
}

// handleCallTool handles the tools/call method
func (s *Server) handleCallTool(req *JSONRPCRequest) {
	// Stub - will be implemented in subtask 04
	var params CallToolParams
	if err := json.Unmarshal(req.Params, &params); err != nil {
		s.sendError(req.ID, InvalidParams, "invalid params", err.Error())
		return
	}

	result := CallToolResult{
		Content: []Content{
			{Type: "text", Text: "Tool not yet implemented"},
		},
	}

	s.sendResult(req.ID, result)
}

// sendResult sends a successful JSON-RPC response
func (s *Server) sendResult(id, result interface{}) {
	resp := JSONRPCResponse{
		JSONRPC: "2.0",
		ID:      id,
		Result:  result,
	}

	data, err := json.Marshal(resp)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to marshal response: %v\n", err)
		return
	}
	if _, err := fmt.Fprintf(s.writer, "%s\n", data); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to write response: %v\n", err)
	}
}

// sendError sends an error JSON-RPC response
func (s *Server) sendError(id interface{}, code int, message string, data interface{}) {
	resp := JSONRPCResponse{
		JSONRPC: "2.0",
		ID:      id,
		Error: &RPCError{
			Code:    code,
			Message: message,
			Data:    data,
		},
	}

	jsonData, err := json.Marshal(resp)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to marshal error response: %v\n", err)
		return
	}
	if _, err := fmt.Fprintf(s.writer, "%s\n", jsonData); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to write error response: %v\n", err)
	}
}

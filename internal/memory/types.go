package memory

import "time"

// Memory represents a stored memory
type Memory struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ID        string    `json:"id"`
	ProjectID string    `json:"project_id"`
	Content   string    `json:"content"`
	Priority  string    `json:"priority"`
	Tags      []string  `json:"tags"`
}

// SearchQuery represents parameters for searching memories
type SearchQuery struct {
	Since     *time.Time
	Query     string
	ProjectID string
	Tags      []string
	Limit     int
}

// SearchResult represents a memory with relevance score
type SearchResult struct {
	Memory Memory  `json:"memory"`
	Score  float64 `json:"score"`
}

// Storage defines the interface for memory storage backends
type Storage interface {
	// Initialize sets up the storage (create tables, etc.)
	Initialize() error

	// Close closes the storage connection
	Close() error

	// Create stores a new memory
	Create(mem *Memory) error

	// Get retrieves a memory by ID
	Get(id string) (*Memory, error)

	// Delete removes a memory by ID
	Delete(id string) error

	// Search performs full-text search on memories
	Search(query SearchQuery) ([]SearchResult, error)

	// List returns recent memories for a project
	List(projectID string, limit int, tags []string) ([]Memory, error)
}

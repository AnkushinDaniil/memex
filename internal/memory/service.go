package memory

import "errors"

// ErrNotImplemented is returned when a method is not yet implemented
var ErrNotImplemented = errors.New("not yet implemented")

// Service handles memory business logic
type Service struct {
	storage Storage
}

// NewService creates a new memory service
func NewService(storage Storage) *Service {
	return &Service{
		storage: storage,
	}
}

// Remember stores a new memory (stub)
func (s *Service) Remember(content string, tags []string, priority, projectID string) (*Memory, error) {
	// Stub - will be implemented in subtask 03
	return nil, ErrNotImplemented
}

// Recall searches for memories (stub)
func (s *Service) Recall(query SearchQuery) ([]SearchResult, error) {
	// Stub - will be implemented in subtask 04
	return nil, ErrNotImplemented
}

// Forget deletes a memory (stub)
func (s *Service) Forget(memoryID string) error {
	// Stub - will be implemented in subtask 04
	return nil
}

// List returns recent memories (stub)
func (s *Service) List(projectID string, limit int, tags []string) ([]Memory, error) {
	// Stub - will be implemented in subtask 04
	return nil, ErrNotImplemented
}

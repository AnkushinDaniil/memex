# Subtask 03: Memory Storage

**Business Value**: Memories persist to SQLite with full-text search indexing, enabling fast retrieval without external dependencies.

**Dependencies**: 01, 02

---

## Implementation Plan

### Files to Create

| File | Purpose |
|------|---------|
| internal/storage/interface.go | Storage interface definition |
| internal/storage/sqlite.go | SQLite + FTS5 implementation |
| internal/memory/service.go | Memory business logic |
| internal/memory/types.go | Memory domain types |

### Steps

1. Create internal/storage/interface.go:
   - Define Storage interface with CRUD + Search methods
   - Keep interface minimal and implementation-agnostic

2. Create internal/memory/types.go:
   - Memory struct with all fields
   - CreateMemoryInput, UpdateMemoryInput
   - SearchOptions, SearchResult

3. Create internal/storage/sqlite.go:
   - NewSQLiteStorage(dbPath) constructor
   - Initialize database with schema and FTS5 table
   - Implement all Storage interface methods

4. Create internal/memory/service.go:
   - NewService(storage) constructor
   - Business logic layer (validation, defaults)
   - Project ID detection from cwd/git

5. Implement FTS5 search:
   - Create FTS5 virtual table
   - Sync triggers for insert/update/delete
   - BM25 ranking for relevance

### Patterns & Hints

```go
// Storage interface
type Storage interface {
    Create(ctx context.Context, mem *Memory) error
    Get(ctx context.Context, id string) (*Memory, error)
    Update(ctx context.Context, mem *Memory) error
    Delete(ctx context.Context, id string) error
    Search(ctx context.Context, projectID, query string, limit int) ([]*Memory, error)
    List(ctx context.Context, projectID string, limit int) ([]*Memory, error)
}
```

```go
// SQLite schema with FTS5
const schema = `
CREATE TABLE IF NOT EXISTS memories (
    id TEXT PRIMARY KEY,
    project_id TEXT NOT NULL,
    content TEXT NOT NULL,
    tags TEXT,
    priority TEXT DEFAULT 'normal',
    retrieval_count INTEGER DEFAULT 0,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE VIRTUAL TABLE IF NOT EXISTS memories_fts USING fts5(
    content,
    tags,
    content='memories',
    content_rowid='rowid',
    tokenize='porter unicode61'
);

-- Triggers to keep FTS in sync
CREATE TRIGGER IF NOT EXISTS memories_ai AFTER INSERT ON memories BEGIN
    INSERT INTO memories_fts(rowid, content, tags)
    VALUES (new.rowid, new.content, new.tags);
END;

CREATE TRIGGER IF NOT EXISTS memories_ad AFTER DELETE ON memories BEGIN
    INSERT INTO memories_fts(memories_fts, rowid, content, tags)
    VALUES('delete', old.rowid, old.content, old.tags);
END;
`
```

```go
// Search with FTS5 and BM25 ranking
func (s *SQLiteStorage) Search(ctx context.Context, projectID, query string, limit int) ([]*Memory, error) {
    rows, err := s.db.QueryContext(ctx, `
        SELECT m.id, m.project_id, m.content, m.tags, m.priority,
               m.retrieval_count, m.created_at, m.updated_at,
               bm25(memories_fts) as rank
        FROM memories m
        JOIN memories_fts ON m.rowid = memories_fts.rowid
        WHERE memories_fts MATCH ?
          AND m.project_id = ?
        ORDER BY rank
        LIMIT ?
    `, query, projectID, limit)
    // ...
}
```

---

## Testing

**Test File**: `internal/storage/sqlite_test.go`

| Test Name | Description | Expected |
|-----------|-------------|----------|
| TestCreateMemory | Insert new memory | Memory stored, ID generated |
| TestGetMemory | Retrieve by ID | Correct memory returned |
| TestUpdateMemory | Modify existing | Changes persisted |
| TestDeleteMemory | Remove memory | Memory gone, FTS updated |
| TestSearchMemory | Full-text search | Matching memories ranked |
| TestSearchNoResults | Search with no matches | Empty slice, no error |
| TestFTSSync | Insert then search | FTS index updated |

**Test File**: `internal/memory/service_test.go`

| Test Name | Description | Expected |
|-----------|-------------|----------|
| TestCreateWithDefaults | Create without priority | Default priority applied |
| TestProjectDetection | Create in git repo | Project ID from git root |
| TestValidation | Create empty content | Error returned |

**Mocking Strategy**: Use in-memory SQLite (`:memory:`) for storage tests. Mock storage for service tests.

---

## Acceptance Criteria

- [ ] SQLite database created at configured path
- [ ] FTS5 table and triggers created automatically
- [ ] Memories persist across server restarts
- [ ] Search returns relevant results ranked by BM25
- [ ] CRUD operations work correctly
- [ ] Project ID detected from git root or cwd

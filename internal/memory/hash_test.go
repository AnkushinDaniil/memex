package memory

import (
	"testing"
	"time"
)

func TestHashCode(t *testing.T) {
	tests := []struct {
		name     string
		code     string
		compare  string
		wantSame bool
	}{
		{
			name:     "identical code produces same hash",
			code:     "func main() {\n\tfmt.Println(\"hello\")\n}",
			wantSame: true,
			compare:  "func main() {\n\tfmt.Println(\"hello\")\n}",
		},
		{
			name:     "different code produces different hash",
			code:     "func main() {\n\tfmt.Println(\"hello\")\n}",
			wantSame: false,
			compare:  "func main() {\n\tfmt.Println(\"goodbye\")\n}",
		},
		{
			name:     "empty code produces consistent hash",
			code:     "",
			wantSame: true,
			compare:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hash1 := HashCode(tt.code)
			hash2 := HashCode(tt.compare)

			// Verify hash is non-empty
			if hash1 == "" {
				t.Error("HashCode returned empty string")
			}

			// Verify hash length (SHA256 = 64 hex chars)
			if len(hash1) != 64 {
				t.Errorf("HashCode length = %d, want 64", len(hash1))
			}

			// Verify comparison
			if tt.wantSame && hash1 != hash2 {
				t.Errorf("Hashes should be equal but differ:\n  got:  %s\n  want: %s", hash1, hash2)
			}
			if !tt.wantSame && hash1 == hash2 {
				t.Errorf("Hashes should differ but are equal: %s", hash1)
			}
		})
	}
}

func TestHashIgnoresWhitespace(t *testing.T) {
	tests := []struct {
		name  string
		code1 string
		code2 string
	}{
		{
			name: "different indentation",
			code1: `func main() {
	fmt.Println("hello")
}`,
			code2: `func main() {
    fmt.Println("hello")
}`,
		},
		{
			name: "extra blank lines",
			code1: `func main() {
	fmt.Println("hello")
}`,
			code2: `func main() {

	fmt.Println("hello")


}`,
		},
		{
			name: "trailing whitespace",
			code1: `func main() {
	fmt.Println("hello")
}`,
			code2: `func main() {
	fmt.Println("hello")
}   `,
		},
		{
			name:  "CRLF vs LF",
			code1: "func main() {\n\tfmt.Println(\"hello\")\n}",
			code2: "func main() {\r\n\tfmt.Println(\"hello\")\r\n}",
		},
		{
			name: "mixed whitespace",
			code1: `func test() {
	x := 42
	return x
}`,
			code2: `func test() {
    x := 42
    return x
}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hash1 := HashCode(tt.code1)
			hash2 := HashCode(tt.code2)

			if hash1 != hash2 {
				t.Errorf("Hashes should be equal despite whitespace differences:\n  code1 hash: %s\n  code2 hash: %s", hash1, hash2)
			}
		})
	}
}

func TestDetectChanges(t *testing.T) {
	originalCode := `func calculate(x int) int {
	return x * 2
}`
	originalHash := HashCode(originalCode)

	tests := []struct {
		name        string
		memory      *Memory
		currentCode string
		wantChanged bool
	}{
		{
			name: "no changes detected",
			memory: &Memory{
				ID:       "mem-1",
				CodeHash: originalHash,
			},
			currentCode: originalCode,
			wantChanged: false,
		},
		{
			name: "whitespace-only changes not detected",
			memory: &Memory{
				ID:       "mem-2",
				CodeHash: originalHash,
			},
			currentCode: `func calculate(x int) int {
    return x * 2
}`,
			wantChanged: false,
		},
		{
			name: "code changes detected",
			memory: &Memory{
				ID:       "mem-3",
				CodeHash: originalHash,
			},
			currentCode: `func calculate(x int) int {
	return x * 3
}`,
			wantChanged: true,
		},
		{
			name: "empty hash means no change detection",
			memory: &Memory{
				ID:       "mem-4",
				CodeHash: "",
			},
			currentCode: "completely different code",
			wantChanged: false,
		},
		{
			name: "major refactor detected",
			memory: &Memory{
				ID:       "mem-5",
				CodeHash: originalHash,
			},
			currentCode: `func compute(value int) int {
	return value << 1
}`,
			wantChanged: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			changed := DetectChanges(tt.memory, tt.currentCode)

			if changed != tt.wantChanged {
				t.Errorf("DetectChanges() = %v, want %v", changed, tt.wantChanged)
			}
		})
	}
}

func TestComputeCodeHash(t *testing.T) {
	tests := []struct {
		name string
		want string
		ctx  CodeContext
	}{
		{
			name: "compute hash from context",
			ctx: CodeContext{
				File:      "test.go",
				Function:  "TestFunc",
				StartLine: 10,
				EndLine:   20,
				Code:      "func TestFunc() {\n\treturn 42\n}",
			},
			want: HashCode("func TestFunc() {\n\treturn 42\n}"),
		},
		{
			name: "empty code context",
			ctx: CodeContext{
				File: "test.go",
				Code: "",
			},
			want: HashCode(""),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ComputeCodeHash(tt.ctx)

			if got != tt.want {
				t.Errorf("ComputeCodeHash() = %v, want %v", got, tt.want)
			}

			// Verify it's a valid SHA256 hash
			if len(got) != 64 {
				t.Errorf("Hash length = %d, want 64", len(got))
			}
		})
	}
}

func TestNormalizeCode(t *testing.T) {
	tests := []struct {
		name string
		code string
		want string
	}{
		{
			name: "removes leading/trailing whitespace",
			code: "  func test()  ",
			want: "func test()",
		},
		{
			name: "removes empty lines",
			code: "func test() {\n\n\treturn 42\n\n}",
			want: "func test() {\nreturn 42\n}",
		},
		{
			name: "normalizes line endings",
			code: "line1\r\nline2\r\nline3",
			want: "line1\nline2\nline3",
		},
		{
			name: "trims each line",
			code: "  func test()  \n   return 42  \n  }  ",
			want: "func test()\nreturn 42\n}",
		},
		{
			name: "handles empty input",
			code: "",
			want: "",
		},
		{
			name: "handles whitespace-only input",
			code: "   \n\n   \n   ",
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := normalizeCode(tt.code)

			if got != tt.want {
				t.Errorf("normalizeCode() =\n%q\nwant:\n%q", got, tt.want)
			}
		})
	}
}

func TestHashCodeConsistency(t *testing.T) {
	code := `package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
`

	// Hash the same code multiple times
	hash1 := HashCode(code)
	hash2 := HashCode(code)
	hash3 := HashCode(code)

	if hash1 != hash2 || hash2 != hash3 {
		t.Errorf("HashCode should be consistent:\n  hash1: %s\n  hash2: %s\n  hash3: %s", hash1, hash2, hash3)
	}
}

func TestDetectChangesWithMemory(t *testing.T) {
	// Integration-style test with full Memory struct
	code1 := `func authenticate(token string) bool {
	return validateToken(token)
}`

	mem := &Memory{
		ID:        "mem-test",
		ProjectID: "test-project",
		Content:   "Auth bug fix",
		Type:      TypeBugFix,
		CodeHash:  HashCode(code1),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Same code with different whitespace
	code2 := `func authenticate(token string) bool {
    return validateToken(token)
}`

	if DetectChanges(mem, code2) {
		t.Error("Should not detect changes for whitespace-only differences")
	}

	// Changed logic
	code3 := `func authenticate(token string) bool {
	return validateToken(token) && checkPermissions(token)
}`

	if !DetectChanges(mem, code3) {
		t.Error("Should detect changes for logic differences")
	}
}

func TestHashCodeWithSpecialCharacters(t *testing.T) {
	tests := []struct {
		name string
		code string
	}{
		{
			name: "unicode characters",
			code: "func test() { fmt.Println(\"こんにちは\") }",
		},
		{
			name: "emoji",
			code: "// 🚀 Launch function\nfunc launch() {}",
		},
		{
			name: "special symbols",
			code: "x := []int{1, 2, 3}\ny := map[string]int{\"a\": 1}",
		},
		{
			name: "multiline strings",
			code: "s := `line1\nline2\nline3`",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hash := HashCode(tt.code)

			// Should produce consistent hash
			hash2 := HashCode(tt.code)
			if hash != hash2 {
				t.Error("Hash should be consistent for special characters")
			}

			// Should be valid SHA256
			if len(hash) != 64 {
				t.Errorf("Hash length = %d, want 64", len(hash))
			}
		})
	}
}

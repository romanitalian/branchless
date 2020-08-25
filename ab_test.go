package branchless

import (
	"testing"
)

// `go test -bench=. -benchmem`

func BenchmarkMinBranchless(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MinBranchless(int64(i), int64(i - 1))
	}
}

func BenchmarkMin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Min(int64(i), int64(i - 1))
	}
}

// `go test .`

func TestMinBranchless(t *testing.T) {
	var val int64
	if val = MinBranchless(1, 3); val != 1 {
		t.Errorf("Expected: 1. Got: %+v", val)
	}
	if val = MinBranchless(3, 1); val != 1 {
		t.Errorf("Expected: 1. Got: %+v", val)
	}
	if val = MinBranchless(2, 21); val != 2 {
		t.Errorf("Expected: 2. Got: %+v", val)
	}
	if val = MinBranchless(21, 2); val != 2 {
		t.Errorf("Expected: 2. Got: %+v", val)
	}

	if val = Abs(-12); val != 12 {
		t.Errorf("Expected: 12. Go: %+v", val)
	}
	if val = Abs(13); val != 13 {
		t.Errorf("Expected: 13. Go: %+v", val)
	}
	if val = Abs(0); val != 0 {
		t.Errorf("Expected: 0. Go: %+v", val)
	}
}

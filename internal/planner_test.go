package internal

import "testing"

func TestAllocateBudget(t *testing.T) {
	tests := []struct {
		total     int
		interests []string
	}{
		{100000, []string{}},
		{100000, []string{"beach"}},
		{100000, []string{"food"}},
		{100000, []string{"wildlife", "food"}},
	}
	for _, tt := range tests {
		b := AllocateBudget(tt.total, tt.interests)
		if b.Total != tt.total {
			t.Fatalf("expected total %d, got %d", tt.total, b.Total)
		}
		sum := b.Accommodation + b.Transport + b.Food + b.Activities
		if sum != tt.total {
			t.Fatalf("split does not add up: %d", sum)
		}
	}
}

package main

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		x    int
		y    int
		want int
	}{
		{1, 2, 3},
		{10, 20, 30},
		{-1, -2, -3},
		{-10, -20, -30},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := add(tt.x, tt.y)
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

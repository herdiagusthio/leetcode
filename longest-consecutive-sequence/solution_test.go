package main

import (
	"testing"
)

func TestLongestConsecutive(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want int
	}{
		{"empty", []int{}, 0},
		{"single", []int{100}, 1},
		{"example", []int{100, 4, 200, 1, 3, 2}, 4},
		{"negatives", []int{0, -1}, 2},
		{"duplicates", []int{1, 2, 2, 3}, 3},
		{"mixed", []int{9, 1, 4, 7, 3, -1, 0, 5, 8, -1, 6}, 7},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := longestConsecutive(tc.in); got != tc.want {
				t.Fatalf("longestConsecutive(%v) = %d, want %d", tc.in, got, tc.want)
			}
		})
	}
}

func BenchmarkLongestConsecutive(b *testing.B) {
	// build a large input with a long consecutive run
	n := 100000
	arr := make([]int, 0, n)
	for i := 0; i < n; i++ {
		arr = append(arr, i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = longestConsecutive(arr)
	}
}

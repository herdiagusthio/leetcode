package main

import "testing"

func TestMaxOperations(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{"example1", []int{1, 2, 3, 4}, 5, 2},
		{"example2", []int{3, 1, 3, 4, 3}, 6, 1},
		{"empty", []int{}, 5, 0},
		{"single", []int{1}, 2, 0},
		{"all_pairable", []int{1, 1, 2, 2}, 3, 2},
		{"no_pairs", []int{1, 2, 3}, 100, 0},
		{"same_numbers_even", []int{2, 2, 2, 2}, 4, 2},
		{"same_numbers_odd", []int{2, 2, 2, 2, 2}, 4, 2},
		{"zeros", []int{0, 0, 0, 0}, 0, 2},
		{"mixed_large", []int{1, 3, 2, 2, 3, 1}, 4, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// copy slice to prevent accidental modification
			numsCopy := append([]int(nil), tt.nums...)
			got := maxOperations(numsCopy, tt.k)
			if got != tt.want {
				t.Fatalf("maxOperations(%v, %d) = %d; want %d", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}

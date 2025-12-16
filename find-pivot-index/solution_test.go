package main

import "testing"

func TestPivotIndex(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"example1", []int{1, 7, 3, 6, 5, 6}, 3},
		{"example2", []int{1, 2, 3}, -1},
		{"example3", []int{2, 1, -1}, 0},
		{"single_element", []int{5}, 0},
		{"two_elements_equal", []int{1, 1}, -1},
		{"pivot_at_start", []int{0, 1, -1}, 0},
		{"pivot_at_end", []int{-1, 1, 0}, 2},
		{"all_zeros", []int{0, 0, 0}, 0},
		{"negative_numbers", []int{-1, -1, -1, 1, 1, 1}, -1},
		{"pivot_middle", []int{1, 2, 3, 4, 3, 2, 1}, 3},
		{"no_pivot_descending", []int{5, 4, 3, 2, 1}, -1},
		{"all_same_nonzero", []int{2, 2, 2}, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := pivotIndex(tt.nums)
			if got != tt.want {
				t.Fatalf("pivotIndex(%v) = %d; want %d", tt.nums, got, tt.want)
			}
		})
	}
}

package main

import "testing"

func TestLongestSubarray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"example1", []int{1, 1, 0, 1}, 3},
		{"example2", []int{0, 1, 1, 1, 0, 1, 1, 0, 1}, 5},
		{"example3", []int{1, 1, 1}, 2},
		{"all_zeros", []int{0, 0, 0, 0}, 0},
		{"single_one", []int{1}, 0},
		{"single_zero", []int{0}, 0},
		{"two_ones", []int{1, 1}, 1},
		{"alternating", []int{1, 0, 1, 0, 1}, 2},
		{"one_zero_at_start", []int{0, 1, 1, 1}, 3},
		{"one_zero_at_end", []int{1, 1, 1, 0}, 3},
		{"multiple_zeros", []int{1, 1, 0, 0, 1, 1, 1}, 3},
		{"long_sequence", []int{1, 1, 1, 1, 1, 0, 1, 1, 1}, 8},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestSubarray(tt.nums)
			if got != tt.want {
				t.Fatalf("longestSubarray(%v) = %d; want %d", tt.nums, got, tt.want)
			}
		})
	}
}

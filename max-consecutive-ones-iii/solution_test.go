package main

import "testing"

func TestLongestOnes(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{"example1", []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2, 6},
		{"example2", []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3, 10},
		{"all_ones", []int{1, 1, 1, 1, 1}, 0, 5},
		{"all_zeros", []int{0, 0, 0, 0, 0}, 2, 2},
		{"k_zero", []int{1, 0, 1, 0, 1}, 0, 1},
		{"k_equals_length", []int{0, 0, 0, 0}, 4, 4},
		{"single_element_one", []int{1}, 0, 1},
		{"single_element_zero", []int{0}, 1, 1},
		{"no_flip_needed", []int{1, 1, 1, 1}, 2, 4},
		{"alternating", []int{1, 0, 1, 0, 1, 0, 1}, 2, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestOnes(tt.nums, tt.k)
			if got != tt.want {
				t.Fatalf("longestOnes(%v, %d) = %d; want %d", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}

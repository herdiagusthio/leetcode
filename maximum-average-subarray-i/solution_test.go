package main

import (
	"math"
	"testing"
)

func TestFindMaxAverage(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want float64
	}{
		{"example1", []int{1, 12, -5, -6, 50, 3}, 4, 12.75},
		{"example2", []int{5}, 1, 5.0},
		{"all_same", []int{3, 3, 3, 3}, 2, 3.0},
		{"negative_numbers", []int{-1, -2, -3, -4}, 2, -1.5},
		{"mixed_values", []int{0, 1, 1, 3, 3}, 4, 2.0},
		{"single_negative", []int{-1}, 1, -1.0},
		{"two_elements", []int{1, 2}, 1, 2.0},
		{"large_window", []int{1, 2, 3, 4, 5}, 5, 3.0},
	}

	const epsilon = 1e-5

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findMaxAverage(tt.nums, tt.k)
			if math.Abs(got-tt.want) > epsilon {
				t.Fatalf("findMaxAverage(%v, %d) = %.5f; want %.5f", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}

package main

import "testing"

func TestMaxArea(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		expect int
	}{
		{"example1", []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{"two ones", []int{1, 1}, 1},
		{"zeros", []int{0, 0, 0}, 0},
		{"single", []int{5}, 0},
		{"increasing", []int{1, 2, 3, 4, 5}, 6},
		{"decreasing", []int{5, 4, 3, 2, 1}, 6},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := maxArea(tc.input)
			if got != tc.expect {
				t.Fatalf("maxArea(%v) = %d; want %d", tc.input, got, tc.expect)
			}
		})
	}
}

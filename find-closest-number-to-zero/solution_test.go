package main

import "testing"

func TestFindClosestNumber(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"example", []int{-4, -1, 1, 4, 8}, 1},
		{"empty", []int{}, 0},
		{"zero_present", []int{5, 0, -2}, 0},
		{"tie_positive", []int{-2, 2}, 2},
		{"tie_negative_and_positive", []int{-3, 3, -1}, -1},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := findClosestNumber(tc.nums)
			if got != tc.want {
				t.Fatalf("%s: for nums=%v want=%d got=%d", tc.name, tc.nums, tc.want, got)
			}
		})
	}
}

package main

import "testing"

func TestLargestAltitude(t *testing.T) {
	tests := []struct {
		name string
		gain []int
		want int
	}{
		{"example1", []int{-5, 1, 5, 0, -7}, 1},
		{"example2", []int{-4, -3, -2, -1, 4, 3, 2}, 0},
		{"all_positive", []int{1, 2, 3, 4}, 10},
		{"all_negative", []int{-1, -2, -3}, 0},
		{"single_positive", []int{5}, 5},
		{"single_negative", []int{-5}, 0},
		{"peak_in_middle", []int{5, -3, 2, -4, 1}, 5},
		{"zero_gains", []int{0, 0, 0}, 0},
		{"mixed_with_peak_start", []int{10, -5, -3, 2}, 10},
		{"up_and_down", []int{52, -91, 72}, 52},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := largestAltitude(tt.gain)
			if got != tt.want {
				t.Fatalf("largestAltitude(%v) = %d; want %d", tt.gain, got, tt.want)
			}
		})
	}
}

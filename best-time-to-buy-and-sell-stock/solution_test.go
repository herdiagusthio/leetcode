package main

import "testing"

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   int
	}{
		{"example", []int{7, 1, 5, 3, 6, 4}, 5},
		{"descending", []int{7, 6, 4, 3, 1}, 0},
		{"small_gain", []int{2, 4, 1}, 2},
		{"empty", []int{}, 0},
		{"single", []int{5}, 0},
		{"two_days_profit", []int{1, 5}, 4},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := maxProfit(tc.prices)
			if got != tc.want {
				t.Fatalf("%s: maxProfit(%v) = %d, want %d", tc.name, tc.prices, got, tc.want)
			}
		})
	}
}

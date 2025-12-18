package main

import "testing"

func TestUniqueOccurrences(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want bool
	}{
		{"example1", []int{1, 2, 2, 1, 1, 3}, true},
		{"example2", []int{1, 2}, false},
		{"example3", []int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}, true},
		{"single_element", []int{1}, true},
		{"all_same", []int{5, 5, 5, 5}, true},
		{"all_unique", []int{1, 2, 3, 4, 5}, false},
		{"two_pairs", []int{1, 1, 2, 2}, false},
		{"three_groups_unique", []int{1, 2, 2, 3, 3, 3}, true},
		{"negative_numbers", []int{-1, -1, -2, -2, -2}, true},
		{"mixed_positive_negative", []int{1, -1, 1, -1, 2}, false},
		{"large_counts", []int{1, 1, 1, 1, 1, 2, 2, 2, 3, 3}, true},
		{"zeros", []int{0, 0, 0, 1, 1}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := uniqueOccurrences(tt.arr)
			if got != tt.want {
				t.Fatalf("uniqueOccurrences(%v) = %v; want %v", tt.arr, got, tt.want)
			}
		})
	}
}

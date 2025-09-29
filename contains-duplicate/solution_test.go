package main

import "testing"

func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want bool
	}{
		{"empty", []int{}, false},
		{"single", []int{1}, false},
		{"noDup", []int{1, 2, 3}, false},
		{"dup", []int{1, 2, 1}, true},
		{"neg", []int{-1, 0, -1}, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := containsDuplicate(tc.in); got != tc.want {
				t.Fatalf("containsDuplicate(%v) = %v, want %v", tc.in, got, tc.want)
			}
		})
	}
}

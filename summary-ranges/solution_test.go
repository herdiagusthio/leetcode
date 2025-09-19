package main

import "testing"

func equalSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestSummaryRanges(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want []string
	}{
		{"empty", []int{}, nil},
		{"single", []int{1}, []string{"1"}},
		{"example", []int{0, 1, 2, 4, 5, 7}, []string{"0->2", "4->5", "7"}},
		{"negatives", []int{-3, -2, -1, 1}, []string{"-3->-1", "1"}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := summaryRanges(tc.in)
			if !equalSlices(got, tc.want) {
				t.Fatalf("got %v want %v", got, tc.want)
			}
		})
	}
}

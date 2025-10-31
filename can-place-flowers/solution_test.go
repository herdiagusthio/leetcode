package main

import "testing"

func TestCanPlaceFlowers(t *testing.T) {
	tests := []struct {
		bed  []int
		n    int
		want bool
	}{
		{[]int{1, 0, 0, 0, 1}, 1, true},
		{[]int{1, 0, 0, 0, 1}, 2, false},
		{[]int{0}, 1, true},
		{[]int{1}, 1, false},
		{[]int{0, 0, 0, 0, 0}, 3, true},
		{[]int{0, 0, 1, 0, 0}, 2, true},
		{[]int{0, 1, 0, 1, 0}, 1, false},
	}

	for _, tc := range tests {
		// copy bed since canPlaceFlowers may modify it
		bed := append([]int(nil), tc.bed...)
		got := canPlaceFlowers(bed, tc.n)
		if got != tc.want {
			t.Fatalf("bed=%v n=%d got=%v want=%v", tc.bed, tc.n, got, tc.want)
		}
	}
}

package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestFindDifference(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		nums2 []int
		want  [][]int
	}{
		{
			"example1",
			[]int{1, 2, 3},
			[]int{2, 4, 6},
			[][]int{{1, 3}, {4, 6}},
		},
		{
			"example2",
			[]int{1, 2, 3, 3},
			[]int{1, 1, 2, 2},
			[][]int{{3}, {}},
		},
		{
			"duplicates_in_nums2",
			[]int{-80, -15, -81, -28, -61, 63, 14, -45, -35, -10},
			[]int{-1, -40, -44, 41, 10, -43, 69, 10, 2},
			[][]int{{-80, -15, -81, -28, -61, 63, 14, -45, -35, -10}, {-1, -40, -44, 41, 10, -43, 69, 2}},
		},
		{
			"empty_nums1",
			[]int{},
			[]int{1, 2, 3},
			[][]int{{}, {1, 2, 3}},
		},
		{
			"empty_nums2",
			[]int{1, 2, 3},
			[]int{},
			[][]int{{1, 2, 3}, {}},
		},
		{
			"identical_arrays",
			[]int{1, 2, 3},
			[]int{1, 2, 3},
			[][]int{{}, {}},
		},
		{
			"no_overlap",
			[]int{1, 2, 3},
			[]int{4, 5, 6},
			[][]int{{1, 2, 3}, {4, 5, 6}},
		},
		{
			"negative_numbers",
			[]int{-1, -2, -3},
			[]int{-2, -4, -5},
			[][]int{{-1, -3}, {-4, -5}},
		},
		{
			"single_elements",
			[]int{1},
			[]int{2},
			[][]int{{1}, {2}},
		},
		{
			"duplicates_in_both",
			[]int{1, 1, 2, 2, 3, 3},
			[]int{2, 2, 3, 3, 4, 4},
			[][]int{{1}, {4}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findDifference(tt.nums1, tt.nums2)

			// Sort both slices for comparison since order doesn't matter
			sort.Ints(got[0])
			sort.Ints(got[1])
			sort.Ints(tt.want[0])
			sort.Ints(tt.want[1])

			// Helper function to compare slices treating nil and empty as equal
			equalSlices := func(a, b []int) bool {
				if len(a) == 0 && len(b) == 0 {
					return true
				}
				return reflect.DeepEqual(a, b)
			}

			if !equalSlices(got[0], tt.want[0]) || !equalSlices(got[1], tt.want[1]) {
				t.Fatalf("findDifference(%v, %v) = %v; want %v", tt.nums1, tt.nums2, got, tt.want)
			}
		})
	}
}

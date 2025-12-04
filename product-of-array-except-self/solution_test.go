package main

import (
	"reflect"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want []int
	}{
		{"example1", []int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
		{"example2", []int{-1, 1, 0, -3, 3}, []int{0, 0, 9, 0, 0}},
		{"singleZero", []int{0, 1, 2, 3}, []int{6, 0, 0, 0}},
		{"multipleZeros", []int{0, 0, 2}, []int{0, 0, 0}},
		{"twoElements", []int{2, 3}, []int{3, 2}},
		{"negatives", []int{-2, 3, 4}, []int{12, -8, -6}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := productExceptSelf(append([]int(nil), tc.in...)) // pass copy to avoid accidental mutation
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("productExceptSelf(%v) = %v, want %v", tc.in, got, tc.want)
			}
		})
	}
}

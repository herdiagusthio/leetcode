package main

import (
	"reflect"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want []int
	}{
		{"example1", []int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
		{"singleZero", []int{0}, []int{0}},
		{"noZeroes", []int{1, 2, 3}, []int{1, 2, 3}},
		{"leadingZeroes", []int{0, 0, 1}, []int{1, 0, 0}},
		{"allZeroes", []int{0, 0, 0}, []int{0, 0, 0}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			input := make([]int, len(tc.in))
			copy(input, tc.in)
			moveZeroes(input)
			if !reflect.DeepEqual(input, tc.want) {
				t.Fatalf("moveZeroes(%v) = %v, want %v", tc.in, input, tc.want)
			}
		})
	}
}

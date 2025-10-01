package main

import (
	"reflect"
	"testing"
)

func TestSortColorsImplementations(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want []int
	}{
		{"empty", []int{}, []int{}},
		{"single color", []int{0, 0, 0}, []int{0, 0, 0}},
		{"already sorted", []int{0, 0, 1, 1, 2, 2}, []int{0, 0, 1, 1, 2, 2}},
		{"reverse", []int{2, 2, 1, 1, 0, 0}, []int{0, 0, 1, 1, 2, 2}},
		{"mixed", []int{2, 0, 2, 1, 1, 0}, []int{0, 0, 1, 1, 2, 2}},
		{"two colors", []int{2, 0, 0, 2, 0}, []int{0, 0, 0, 2, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name+"/DNF", func(t *testing.T) {
			in := append([]int(nil), tt.in...)
			sortColors(in)
			if !reflect.DeepEqual(normalize(in), normalize(tt.want)) {
				t.Fatalf("sortColors DNF: got %v want %v", in, tt.want)
			}
		})

		t.Run(tt.name+"/Count", func(t *testing.T) {
			in := append([]int(nil), tt.in...)
			sortColorsCount(in)
			if !reflect.DeepEqual(normalize(in), normalize(tt.want)) {
				t.Fatalf("sortColorsCount: got %v want %v", in, tt.want)
			}
		})
	}
}

func normalize(s []int) []int {
	if len(s) == 0 {
		return nil
	}
	return s
}

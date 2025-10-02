package main

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name    string
		numbers []int
		target  int
		want    []int
	}{
		{"example1", []int{2, 7, 11, 15}, 9, []int{1, 2}},
		{"example2", []int{2, 3, 4}, 6, []int{1, 3}},
		{"example3", []int{-1, 0}, -1, []int{1, 2}},
		{"duplicate values", []int{1, 2, 3, 3, 4}, 6, []int{3, 4}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := twoSum(append([]int(nil), tt.numbers...), tt.target)
			if tt.name == "duplicate values" {
				// There are multiple valid pairs; validate returned indices are correct
				if got == nil || len(got) != 2 {
					t.Fatalf("twoSum(%v, %d) returned %v, want two indices", tt.numbers, tt.target, got)
				}
				i, j := got[0], got[1]
				if i < 1 || j < 1 || i >= j || j > len(tt.numbers) {
					t.Fatalf("invalid indices returned: %v", got)
				}
				if tt.numbers[i-1]+tt.numbers[j-1] != tt.target {
					t.Fatalf("indices do not sum to target: nums[%d]=%d + nums[%d]=%d != %d", i-1, tt.numbers[i-1], j-1, tt.numbers[j-1], tt.target)
				}
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("twoSum(%v, %d) = %v; want %v", tt.numbers, tt.target, got, tt.want)
			}
		})
	}
}

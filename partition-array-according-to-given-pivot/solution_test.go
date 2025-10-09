package main

import (
    "reflect"
    "sort"
    "testing"
)

func equalMultiset(a, b []int) bool {
    if len(a) != len(b) {
        return false
    }
    a2 := append([]int(nil), a...)
    b2 := append([]int(nil), b...)
    sort.Ints(a2)
    sort.Ints(b2)
    return reflect.DeepEqual(a2, b2)
}

func TestPivotStableImplementations(t *testing.T) {
    tests := []struct{
        name string
        nums []int
        pivot int
        want []int
    }{
        {"example1", []int{9,12,5,10,14,3,10}, 10, []int{9,5,3,10,10,12,14}},
        {"example2", []int{-3,4,3,2}, 2, []int{-3,2,4,3}},
        {"all less", []int{1,1,0}, 5, []int{1,1,0}},
        {"all greater", []int{5,6,7}, 5, []int{5,6,7}},
    {"duplicates", []int{2,2,1,3,2}, 2, []int{1,2,2,2,3}},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // stable implementations must equal expected
            got1 := pivotArray(append([]int(nil), tt.nums...), tt.pivot)
            if !reflect.DeepEqual(got1, tt.want) {
                t.Fatalf("pivotArray: got %v want %v", got1, tt.want)
            }

            got2 := pivotArrayPrealloc(append([]int(nil), tt.nums...), tt.pivot)
            if !reflect.DeepEqual(got2, tt.want) {
                t.Fatalf("pivotArrayPrealloc: got %v want %v", got2, tt.want)
            }

            got3 := pivotArrayCount(append([]int(nil), tt.nums...), tt.pivot)
            if !reflect.DeepEqual(got3, tt.want) {
                t.Fatalf("pivotArrayCount: got %v want %v", got3, tt.want)
            }
        })
    }
}

func TestPivotDNFPartitionProperty(t *testing.T) {
    // DNF implementation removed; this test no longer applies.
}

package main

import (
    "reflect"
    "testing"
)

func TestPartition(t *testing.T) {
    tests := []struct{
        name string
        in []int
        x int
        want []int
    }{
        {name: "nil list", in: nil, x: 1, want: []int{}},
        {name: "single less", in: []int{1}, x: 2, want: []int{1}},
        {name: "single greater", in: []int{3}, x: 2, want: []int{3}},
        {name: "example1", in: []int{1,4,3,2,5,2}, x: 3, want: []int{1,2,2,4,3,5}},
        {name: "example2", in: []int{2,1}, x: 2, want: []int{1,2}},
        {name: "all less", in: []int{1,0,-1}, x: 5, want: []int{1,0,-1}},
        {name: "all greater", in: []int{5,6,7}, x: 5, want: []int{5,6,7}},
    {name: "duplicates", in: []int{2,2,1,3,2}, x: 2, want: []int{1,2,2,3,2}},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            head := buildList(tt.in)
            got := partition(head, tt.x)
            gotSlice := toSlice(got)
            // Normalize nil vs empty slice result
            if gotSlice == nil {
                gotSlice = []int{}
            }
            if !reflect.DeepEqual(gotSlice, tt.want) {
                t.Fatalf("%s: partition(%v, %d) = %v; want %v", tt.name, tt.in, tt.x, gotSlice, tt.want)
            }

            // Also ensure the alternative implementation matches the same output
            head2 := buildList(tt.in)
            got2 := partitionArray(head2, tt.x)
            gotSlice2 := toSlice(got2)
            if gotSlice2 == nil {
                gotSlice2 = []int{}
            }
            if !reflect.DeepEqual(gotSlice2, gotSlice) {
                t.Fatalf("%s: partitionArray(%v, %d) = %v; partition returned %v", tt.name, tt.in, tt.x, gotSlice2, gotSlice)
            }
        })
    }
}

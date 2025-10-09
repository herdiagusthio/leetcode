package main

import "fmt"

// pivotArray simple append-based stable solution (three slices, then concat).
func pivotArray(nums []int, pivot int) []int {
    left, mid, right := []int{}, []int{}, []int{}

    for _, num := range nums {
        if num < pivot {
            left = append(left, num)
        } else if num == pivot {
            mid = append(mid, num)
        } else {
            right = append(right, num)
        }
    }

    left = append(left, mid...)
    left = append(left, right...)
    return left
}

// pivotArrayPrealloc keeps the append-based style but reduces reallocations by
// preallocating capacity for left/mid/right slices before appending.
func pivotArrayPrealloc(nums []int, pivot int) []int {
    if len(nums) == 0 {
        return nums
    }

    var cntLeft, cntMid, cntRight int
    for _, v := range nums {
        if v < pivot {
            cntLeft++
        } else if v == pivot {
            cntMid++
        } else {
            cntRight++
        }
    }

    left := make([]int, 0, cntLeft)
    mid := make([]int, 0, cntMid)
    right := make([]int, 0, cntRight)

    for _, v := range nums {
        if v < pivot {
            left = append(left, v)
        } else if v == pivot {
            mid = append(mid, v)
        } else {
            right = append(right, v)
        }
    }

    left = append(left, mid...)
    left = append(left, right...)
    return left
}

// pivotArrayCount reduces allocations by preallocating the final result slice
// and filling it in a second pass; stable and efficient.
func pivotArrayCount(nums []int, pivot int) []int {
    if len(nums) == 0 {
        return nums
    }

    // first pass: count sizes
    var cntLeft, cntMid, cntRight int
    for _, v := range nums {
        if v < pivot {
            cntLeft++
        } else if v == pivot {
            cntMid++
        } else {
            cntRight++
        }
    }

    res := make([]int, cntLeft+cntMid+cntRight)
    iLeft := 0
    iMid := cntLeft
    iRight := cntLeft + cntMid

    for _, v := range nums {
        if v < pivot {
            res[iLeft] = v
            iLeft++
        } else if v == pivot {
            res[iMid] = v
            iMid++
        } else {
            res[iRight] = v
            iRight++
        }
    }

    return res
}

func main() {
    nums := []int{9, 12, 5, 10, 14, 3, 10}
    pivot := 10

    a := pivotArray(append([]int(nil), nums...), pivot)
    b := pivotArrayPrealloc(append([]int(nil), nums...), pivot)
    c := pivotArrayCount(append([]int(nil), nums...), pivot)

    fmt.Println("input:", nums)
    fmt.Println("pivot:", pivot)
    fmt.Println("pivotArray:", a)
    fmt.Println("pivotArrayPrealloc:", b)
    fmt.Println("pivotArrayCount:", c)
}

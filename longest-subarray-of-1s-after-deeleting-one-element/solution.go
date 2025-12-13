package main

import "fmt"

// longestSubarray returns the size of the longest non-empty subarray containing only 1's after deleting exactly one element.
func longestSubarray(nums []int) int {
	var result, left int

	zeroCount := 0
	potentialLength := 0
	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			zeroCount++
		}

		if zeroCount > 1 {
			if nums[left] == 0 {
				zeroCount--
			}
			left++
		}

		if zeroCount < 2 {
			potentialLength = right - left
			if potentialLength > result {
				result = potentialLength
			}
		}
	}

	return result
}

func main() {
	examples := []struct {
		nums []int
	}{
		{[]int{1, 1, 0, 1}},
		{[]int{0, 1, 1, 1, 0, 1, 1, 0, 1}},
		{[]int{1, 1, 1}},
	}

	for _, ex := range examples {
		fmt.Printf("longestSubarray(%v) = %d\n", ex.nums, longestSubarray(ex.nums))
	}
}

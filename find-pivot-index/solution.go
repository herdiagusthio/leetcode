package main

import "fmt"

// pivotIndex returns the leftmost pivot index where the sum of elements to the left equals the sum to the right.
func pivotIndex(nums []int) int {
	var left, right, totalSum int

	for _, num := range nums {
		totalSum += num
	}

	for i, num := range nums {
		right = totalSum - left - num

		if left == right {
			return i
		}
		left += num
	}

	return -1
}

func main() {
	examples := []struct {
		nums []int
	}{
		{[]int{1, 7, 3, 6, 5, 6}},
		{[]int{1, 2, 3}},
		{[]int{2, 1, -1}},
	}

	for _, ex := range examples {
		fmt.Printf("pivotIndex(%v) = %d\n", ex.nums, pivotIndex(ex.nums))
	}
}

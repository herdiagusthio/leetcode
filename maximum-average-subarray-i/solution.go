package main

import "fmt"

// findMaxAverage returns the maximum average of a contiguous subarray of length k.
func findMaxAverage(nums []int, k int) float64 {
	// Calculate initial sum of the first k elements
	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}

	maxSum := sum

	// Slide the window from k element to the end
	for i := k; i < len(nums); i++ {
		// Update the sum by adding the new element and removing the element going out of the window
		sum += nums[i] - nums[i-k]
		if sum > maxSum {
			maxSum = sum
		}
	}

	return float64(maxSum) / float64(k)
}

func main() {
	examples := []struct {
		nums []int
		k    int
	}{
		{[]int{1, 12, -5, -6, 50, 3}, 4},
		{[]int{5}, 1},
		{[]int{0, 1, 1, 3, 3}, 4},
		{[]int{-1}, 1},
	}

	for _, ex := range examples {
		fmt.Printf("nums=%v k=%d -> %.5f\n", ex.nums, ex.k, findMaxAverage(ex.nums, ex.k))
	}
}

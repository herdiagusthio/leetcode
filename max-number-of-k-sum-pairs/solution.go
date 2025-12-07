package main

import "fmt"

// maxOperations returns the maximum number of disjoint pairs that sum to k.
func maxOperations(nums []int, k int) int {
	counts := make(map[int]int, len(nums))
	pairs := 0

	for _, num := range nums {
		diff := k - num

		if counts[diff] > 0 {
			pairs++
			counts[diff]--
		} else {
			counts[num]++
		}
	}

	return pairs
}

func main() {
	examples := []struct {
		nums []int
		k    int
	}{
		{[]int{1, 2, 3, 4}, 5},
		{[]int{3, 1, 3, 4, 3}, 6},
		{[]int{0, 0, 0, 0}, 0},
	}

	for _, ex := range examples {
		fmt.Printf("nums=%v k=%d -> %d\n", ex.nums, ex.k, maxOperations(ex.nums, ex.k))
	}
}

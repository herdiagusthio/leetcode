package main

import "fmt"

// containsDuplicate returns true if any value appears at least twice in the slice.
// Time: O(n), Space: O(n)
func containsDuplicate(nums []int) bool {
	if len(nums) < 2 {
		return false
	}

	seen := make(map[int]struct{}, len(nums))
	for _, num := range nums {
		if _, exists := seen[num]; exists {
			return true
		}
		seen[num] = struct{}{}
	}

	return false
}

func main() {
	examples := [][]int{
		{},
		{1},
		{1, 2, 3},
		{1, 2, 1},
		{-1, 0, -1},
	}

	for _, ex := range examples {
		fmt.Printf("input: %v -> containsDuplicate: %v\n", ex, containsDuplicate(ex))
	}
}

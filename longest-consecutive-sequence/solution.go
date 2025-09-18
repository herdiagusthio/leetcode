package main

import "fmt"

// longestConsecutive returns the length of the longest consecutive elements sequence.
//
// Complexity:
// - Time: O(n) on average. Each number is inserted into the set and visited at most
//   twice: once when placed into the set and once when expanding a sequence.
// - Space: O(n) for the set.
func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// use struct{} as value to keep map as a set with minimal allocation
	set := make(map[int]struct{}, len(nums))
	for _, n := range nums {
		set[n] = struct{}{}
	}

	longest := 0
	for n := range set {
		// only start counting at the beginning of a sequence
		if _, hasPrev := set[n-1]; hasPrev {
			continue
		}

		streak := 1
		curNum := n
		for {
			if _, ok := set[curNum+1]; !ok {
				break
			}
			curNum++
			streak++
		}

		if streak > longest {
			longest = streak
		}
	}

	return longest
}

// example main to demonstrate usage
func main() {
	cases := [][]int{
		{},
		{100},
		{100, 4, 200, 1, 3, 2},
		{1, 2, 2, 3},
		{9, 1, 4, 7, 3, -1, 0, 5, 8, -1, 6},
	}

	for _, c := range cases {
		fmt.Println("input:", c, "-> longest consecutive:", longestConsecutive(c))
	}
}

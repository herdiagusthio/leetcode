package main

import "fmt"

// sortColors sorts nums in-place so that all 0s come first, then 1s, then 2s.
// The canonical LeetCode entrypoint is `sortColors`. Below we provide two
// implementations:
//  - sortColors (Dutch National Flag) : single-pass, in-place, O(n) time O(1) extra.
//  - sortColorsCount                 : two-pass using a fixed-size count array, O(n) time O(1) extra.

// sortColors uses the Dutch National Flag algorithm (three-way partition).
func sortColors(nums []int) {
	low, mid, high := 0, 0, len(nums)-1
	for mid <= high {
		switch nums[mid] {
		case 0:
			nums[low], nums[mid] = nums[mid], nums[low]
			low++
			mid++
		case 1:
			mid++
		case 2:
			nums[mid], nums[high] = nums[high], nums[mid]
			high--
		default:
			// For inputs outside 0..2 (not expected in LeetCode), advance mid.
			mid++
		}
	}
}

// sortColorsCount is an alternative implementation that counts occurrences
// then overwrites the slice. It is simple and uses constant extra space.
func sortColorsCount(nums []int) {
	if len(nums) == 0 {
		return
	}
	counts := [3]int{}
	for _, v := range nums {
		if v >= 0 && v < 3 {
			counts[v]++
		}
	}

	idx := 0
	for color := 0; color < 3; color++ {
		for k := 0; k < counts[color]; k++ {
			nums[idx] = color
			idx++
		}
	}
}

func main() {
	samples := [][]int{
		{},
		{0},
		{2, 0, 2, 1, 1, 0},
		{2, 0, 0, 2, 0},
	}

	fmt.Println("Demonstrating sortColors (DNF) and sortColorsCount:")
	for i, s := range samples {
		a := append([]int(nil), s...)
		b := append([]int(nil), s...)

		fmt.Printf("\nSample %d before: %v\n", i+1, s)

		sortColors(a)
		fmt.Printf("sortColors (DNF)  -> %v\n", a)

		sortColorsCount(b)
		fmt.Printf("sortColorsCount   -> %v\n", b)
	}
}

package main

import "fmt"

// findDifference returns two lists: distinct integers in nums1 not in nums2, and distinct integers in nums2 not in nums1.
func findDifference(nums1 []int, nums2 []int) [][]int {
	var answer [][]int
	numsMap := make(map[int]int, len(nums1)+len(nums2))

	for _, num := range nums1 {
		numsMap[num] = 0
	}

	for _, num := range nums2 {
		if _, exists := numsMap[num]; !exists {
			numsMap[num] = 1
		} else if numsMap[num] == 0 {
			numsMap[num] = 2
		}
	}

	var answer0, answer1 []int

	for k, v := range numsMap {
		switch v {
		case 0:
			answer0 = append(answer0, k)
		case 1:
			answer1 = append(answer1, k)
		}
	}

	answer = append(answer, answer0, answer1)

	return answer
}

func main() {
	examples := []struct {
		nums1 []int
		nums2 []int
	}{
		{[]int{1, 2, 3}, []int{2, 4, 6}},
		{[]int{1, 2, 3, 3}, []int{1, 1, 2, 2}},
		{[]int{-80, -15, -81, -28, -61, 63, 14, -45, -35, -10}, []int{-1, -40, -44, 41, 10, -43, 69, 10, 2}},
	}

	for i, ex := range examples {
		result := findDifference(ex.nums1, ex.nums2)
		fmt.Printf("Example %d:\n", i+1)
		fmt.Printf("  nums1 = %v\n", ex.nums1)
		fmt.Printf("  nums2 = %v\n", ex.nums2)
		fmt.Printf("  result = %v\n\n", result)
	}
}

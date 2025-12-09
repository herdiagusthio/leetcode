package main

import "fmt"

// longestOnes returns the maximum number of consecutive 1's in the array if you can flip at most k 0's.
func longestOnes(nums []int, k int) int {
	var maxNumber, number int
	currentZero := 0

	left := 0
	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			currentZero++
		}

		if currentZero > k {
			if nums[left] == 0 {
				currentZero--
			}
			left++
		}

		number = right - left + 1
		if number > maxNumber {
			maxNumber = number
		}
	}

	return maxNumber
}

func main() {
	examples := []struct {
		nums []int
		k    int
	}{
		{[]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2},
		{[]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3},
	}

	for _, ex := range examples {
		fmt.Printf("longestOnes(%v, %d) = %d\n", ex.nums, ex.k, longestOnes(ex.nums, ex.k))
	}
}

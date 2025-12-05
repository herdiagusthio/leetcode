package main

import "fmt"

func moveZeroes(nums []int) {
	n := len(nums)
	pointer := 0
	for i := 0; i < n; i++ {
		if nums[i] != 0 {
			nums[pointer], nums[i] = nums[i], nums[pointer]
			pointer++
		}
	}
}

func main() {
	examples := [][]int{
		{0, 1, 0, 3, 12},
		{0},
		{1, 2, 3},
		{0, 0, 1},
	}

	for _, ex := range examples {
		input := make([]int, len(ex))
		copy(input, ex)
		moveZeroes(input)
		fmt.Printf("input: %v -> output: %v\n", ex, input)
	}
}

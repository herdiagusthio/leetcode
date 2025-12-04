package main

import "fmt"

func productExceptSelf(nums []int) []int {
	n := len(nums)
	res := make([]int, n)

	// prefix products (product of elements left of i)
	leftProd := 1
	for i := 0; i < n; i++ {
		res[i] = leftProd
		leftProd *= nums[i]
	}

	// suffix products (product of elements right of i)
	rightProd := 1
	for i := n - 1; i >= 0; i-- {
		res[i] *= rightProd
		rightProd *= nums[i]
	}

	return res
}

func main() {
	examples := [][]int{
		{1, 2, 3, 4},
		{-1, 1, 0, -3, 3},
		{0, 1, 2, 3},
		{0, 0, 2},
		{2, 3},
		{-2, 3, 4},
	}

	for _, ex := range examples {
		fmt.Printf("input: %v -> output: %v\n", ex, productExceptSelf(ex))
	}
}

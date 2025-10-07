package main

import "fmt"

func maxArea(height []int) int {
	result, area, i, j := 0, 0, 0, len(height)-1
	for i < j {
		if height[i] <= height[j] {
			area = height[i] * (j - i)
			i++
		} else {
			area = height[j] * (j - i)
			j--
		}
		if area > result {
			result = area
		}
	}
	return result
}

func main() {
	examples := [][]int{
		{1, 8, 6, 2, 5, 4, 8, 3, 7},
		{1, 1},
		{0, 0, 0},
	}
	for _, ex := range examples {
		fmt.Printf("height=%v -> maxArea=%d\n", ex, maxArea(ex))
	}
}

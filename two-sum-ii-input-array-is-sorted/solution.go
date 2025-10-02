package main

import "fmt"

// twoSum finds two numbers in a sorted slice that add up to target.
// It returns 1-based indices [i, j] where i < j.
func twoSum(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1
	for i < j {
		sum := numbers[i] + numbers[j]
		if sum == target {
			return []int{i + 1, j + 1}
		}
		
		if sum > target {
			j--
		} else {
			i++
		}
	}
	return nil
}

func main() {
	examples := []struct {
		numbers []int
		target  int
	}{
		{[]int{2, 7, 11, 15}, 9},
		{[]int{2, 3, 4}, 6},
		{[]int{-1, 0}, -1},
	}

	for _, ex := range examples {
		fmt.Printf("numbers=%v target=%d -> %v\n", ex.numbers, ex.target, twoSum(ex.numbers, ex.target))
	}
}

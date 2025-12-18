package main

import "fmt"

// uniqueOccurrences returns true if the number of occurrences of each value in the array is unique.
func uniqueOccurrences(arr []int) bool {
	occurrences := make(map[int]int)
	for _, num := range arr {
		occurrences[num]++
	}

	seen := make(map[int]bool)
	for _, count := range occurrences {
		if seen[count] {
			return false
		}
		seen[count] = true
	}

	return true
}

func main() {
	examples := []struct {
		arr []int
	}{
		{[]int{1, 2, 2, 1, 1, 3}},
		{[]int{1, 2}},
		{[]int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}},
	}

	for i, ex := range examples {
		fmt.Printf("Example %d: uniqueOccurrences(%v) = %v\n", i+1, ex.arr, uniqueOccurrences(ex.arr))
	}
}

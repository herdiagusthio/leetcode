package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	grid := [][]int{{3, 2, 1}, {1, 7, 6}, {2, 7, 7}}
	fmt.Printf("Input: %v\n", grid)
	fmt.Printf("Output: %d\n", equalPairs(grid))
}

func equalPairs(grid [][]int) int {
	var pairs int

	rowsMap := make(map[string]int)

	// convert row to string
	for i := 0; i < len(grid); i++ {
		key := rowToString(grid[i])

		rowsMap[key]++
	}

	// convert column to string
	for col := 0; col < len(grid); col++ {
		var sb strings.Builder
		for row := 0; row < len(grid); row++ {
			sb.WriteString(strconv.Itoa(grid[row][col]))
			if row < len(grid)-1 {
				sb.WriteByte(',')
			}
		}

		key := sb.String()

		if val, exist := rowsMap[key]; exist {
			pairs += val
		}
	}

	return pairs
}

func rowToString(row []int) string {
	var sb strings.Builder

	for i, num := range row {
		sb.WriteString(strconv.Itoa(num))
		if i < len(row)-1 {
			sb.WriteByte(',')
		}
	}

	return sb.String()
}

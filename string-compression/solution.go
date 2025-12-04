package main

import (
	"fmt"
	"strconv"
)

func compress(chars []byte) int {
	result := []byte{chars[0]}
	pointer := chars[0]
	count := 1
	for i := 1; i < len(chars); i++ {
		if chars[i] == pointer {
			count++
		} else {
			if count > 1 {
				countStr := strconv.Itoa(count)
				result = append(result, []byte(countStr)...)
				count = 1
			}
			pointer = chars[i]
			result = append(result, chars[i])
		}
	}

	if count > 1 {
		countStr := strconv.Itoa(count)
		result = append(result, []byte(countStr)...)
	}

	copy(chars, result)

	return len(result)
}

func main() {
	examples := [][]byte{
		{'a', 'a', 'b', 'b', 'c', 'c', 'c'},
		{'a'},
		{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'},
	}

	for _, ex := range examples {
		// Make a copy to avoid modifying the original example for printing (though we print after modification anyway)
		input := make([]byte, len(ex))
		copy(input, ex)
		newLen := compress(input)
		fmt.Printf("input: %s -> output length: %d, content: %s\n", string(ex), newLen, string(input[:newLen]))
	}
}

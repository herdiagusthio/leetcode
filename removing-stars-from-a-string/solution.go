package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "leet**cod*e"
	fmt.Printf("Input: %s\n", input)
	fmt.Printf("Output: %s\n", removeStars(input))
}

// removeStars removes stars from the string using a stack-based approach.
func removeStars(s string) string {
	stack := make([]byte, 0, len(s))

	for i := 0; i < len(s); i++ {
		if s[i] == '*' {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, s[i])
		}
	}

	return string(stack)
}

// removeStarsBackward removes stars by iterating backwards and skipping characters.
func removeStarsBackward(s string) string {
	var counter int
	sliceChar := make([]byte, 0, len(s))

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '*' {
			counter++
		} else {
			if counter == 0 {
				sliceChar = append(sliceChar, s[i])
			} else {
				counter--
			}
		}
	}

	var sb strings.Builder
	for i := len(sliceChar) - 1; i >= 0; i-- {
		sb.WriteByte(sliceChar[i])
	}

	return sb.String()
}

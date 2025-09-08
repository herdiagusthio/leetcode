package main

import (
	"fmt"
)

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	mapCounter := make(map[rune]int)

	for _, char := range s {
		mapCounter[char]++
	}

	for _, char := range t {
		mapCounter[char]--
		if mapCounter[char] < 0 {
			return false
		}
	}

	return true
}

func main() {

	// Otherwise run a few example cases and print readable output.
	examples := []struct{ s, t string }{
		{"anagram", "nagaram"},
		{"rat", "car"},
		{"listen", "silent"},
		{"", ""},
	}

	for _, ex := range examples {
		fmt.Printf("isAnagram(%q, %q) = %v\n", ex.s, ex.t, isAnagram(ex.s, ex.t))
	}
}

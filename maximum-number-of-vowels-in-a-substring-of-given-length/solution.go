package main

import "fmt"

// maxVowels returns the maximum number of vowel letters in any substring of s with length k.
func maxVowels(s string, k int) int {
	var maxNum, counter int
	isVowel := func(c byte) bool {
		switch c {
		case 'a', 'e', 'i', 'o', 'u':
			return true
		default:
			return false
		}
	}

	// Open the first window
	for i := 0; i < k; i++ {
		if isVowel(s[i]) {
			counter++
		}
	}

	maxNum = counter

	// Slide the window
	for i := k; i < len(s); i++ {
		if isVowel(s[i-k]) {
			counter--
		}

		if isVowel(s[i]) {
			counter++
		}

		if counter > maxNum {
			maxNum = counter
		}
	}

	return maxNum
}

func main() {
	examples := []struct {
		s string
		k int
	}{
		{"abciiidef", 3},
		{"aeiou", 2},
		{"leetcode", 3},
		{"rhythms", 4},
		{"tryhard", 4},
	}

	for _, ex := range examples {
		fmt.Printf("maxVowels(%q, %d) = %d\n", ex.s, ex.k, maxVowels(ex.s, ex.k))
	}
}

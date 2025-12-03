package main

import "fmt"

// reverseVowels reverses only the vowels in the input string.
// It works on runes to handle multi-byte characters correctly.
func reverseVowels(s string) string {
	runes := []rune(s)
	i, j := 0, len(runes)-1

	for i < j {
		for i < j && !isVowel(runes[i]) {
			i++
		}

		for i < j && !isVowel(runes[j]) {
			j--
		}

		if i < j {
			runes[i], runes[j] = runes[j], runes[i]
			i++
			j--
		}
	}

	return string(runes)
}

func isVowel(r rune) bool {
	switch r {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	default:
		return false
	}
}

func main() {
	examples := []string{
		"IceCreAm",
		"leetcode",
		"",
		"a",
		"b",
		"Aa",
		"AEIOU",
	}

	for _, ex := range examples {
		fmt.Printf("input: %q -> output: %q\n", ex, reverseVowels(ex))
	}
}

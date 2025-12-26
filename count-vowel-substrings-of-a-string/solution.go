package main

import "fmt"

// countVowelSubstrings returns the number of vowel substrings that contain all five vowels.
func countVowelSubstrings(word string) int {
	var result int
	vowFreq := make(map[byte]int)

	lastConsonant := -1
	left := 0
	for right := 0; right < len(word); right++ {
		if !isVowel(word[right]) {
			lastConsonant = right
			vowFreq = make(map[byte]int)
			left = right + 1
			continue
		}

		vowFreq[word[right]]++

		for len(vowFreq) == 5 && vowFreq[word[left]] > 1 {
			vowFreq[word[left]]--
			left++
		}

		if len(vowFreq) == 5 {
			result += left - lastConsonant
		}
	}

	return result
}

func isVowel(c byte) bool {
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
}

func main() {
	examples := []struct {
		word string
	}{
		{"aeiouu"},
		{"unicornarihan"},
		{"cuaieuouac"},
	}

	for i, ex := range examples {
		fmt.Printf("Example %d: countVowelSubstrings(%q) = %d\n", i+1, ex.word, countVowelSubstrings(ex.word))
	}
}
package main

import (
	"fmt"
	"sort"
)

// closeStrings returns true if word1 and word2 are "close" strings.
// Two strings are close if they can be transformed into each other using:
// 1. Swapping any two characters, or
// 2. Transforming all occurrences of one character into another (and vice versa).
func closeStrings(word1 string, word2 string) bool {
	// Different lengths cannot be close
	if len(word1) != len(word2) {
		return false
	}

	// Count character frequencies
	freqWord1 := make(map[rune]int, len(word1))
	for _, char := range word1 {
		freqWord1[char]++
	}

	freqWord2 := make(map[rune]int, len(word2))
	for _, char := range word2 {
		freqWord2[char]++
	}

	// Check if both strings have the same set of unique characters
	if len(freqWord1) != len(freqWord2) {
		return false
	}

	for char := range freqWord1 {
		if _, exists := freqWord2[char]; !exists {
			return false
		}
	}

	// Extract frequency counts
	countsWord1 := make([]int, 0, len(freqWord1))
	for _, count := range freqWord1 {
		countsWord1 = append(countsWord1, count)
	}

	countsWord2 := make([]int, 0, len(freqWord2))
	for _, count := range freqWord2 {
		countsWord2 = append(countsWord2, count)
	}

	// Sort and compare frequency distributions
	sort.Ints(countsWord1)
	sort.Ints(countsWord2)

	for i := range countsWord1 {
		if countsWord1[i] != countsWord2[i] {
			return false
		}
	}

	return true
}

func main() {
	examples := []struct {
		word1 string
		word2 string
	}{
		{"abc", "bca"},
		{"a", "aa"},
		{"cabbba", "abbccc"},
		{"uau", "ssx"},
		{"aabbcc", "ccbbaa"},
	}

	for _, ex := range examples {
		fmt.Printf("input: word1=%q, word2=%q -> output: %v\n", ex.word1, ex.word2, closeStrings(ex.word1, ex.word2))
	}
}

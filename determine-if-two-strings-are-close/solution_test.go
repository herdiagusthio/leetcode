package main

import (
	"testing"
)

func TestCloseStrings(t *testing.T) {
	tests := []struct {
		name  string
		word1 string
		word2 string
		want  bool
	}{
		// Example test cases from LeetCode
		{"example1", "abc", "bca", true},
		{"example2", "a", "aa", false},
		{"example3", "cabbba", "abbccc", true},

		// Basic cases
		{"identicalStrings", "hello", "hello", true},
		{"singleChar", "a", "a", true},
		{"twoEqualChars", "ab", "ba", true},
		{"emptyStrings", "", "", true},

		// Different lengths
		{"differentLengths1", "abc", "abcd", false},
		{"differentLengths2", "a", "aa", false},
		{"differentLengths3", "hello", "hi", false},

		// Different character sets
		{"differentChars1", "uau", "ssx", false},
		{"differentChars2", "abc", "def", false},
		{"differentChars3", "abbzzca", "babzzcz", false},

		// Same characters, different frequencies
		{"sameCharsDiffFreq1", "aabbcc", "ccbbaa", true},
		{"sameCharsDiffFreq2", "aaabbb", "bbbaaa", true},
		{"sameCharsDiffFreq3", "aaabbbbccddeeeeefffff", "aaaaabbcccdddeeeeffff", false},

		// Edge cases with single occurrences
		{"allUnique", "abcdef", "fedcba", true},
		{"allSame", "aaaa", "aaaa", true},
		{"oneDifferent", "aaaa", "aaab", false},

		// Complex transformations
		{"canTransform1", "aacabb", "bbcbaa", true},
		{"canTransform2", "xxyyzz", "zzyyxx", true},
		{"cannotTransform", "aabbcc", "aabbc", false},

		// Large frequency differences
		{"largeFreq1", "aaaabbbb", "aabbbbaa", true},
		{"largeFreq2", "aaabbb", "aabbbb", false},

		// Mixed cases
		{"mixedCase1", "abcabc", "cbacba", true},
		{"mixedCase2", "aabbcc", "aabbccc", false}, // different lengths
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := closeStrings(tc.word1, tc.word2)
			if got != tc.want {
				t.Fatalf("closeStrings(%q, %q) = %v, want %v", tc.word1, tc.word2, got, tc.want)
			}
		})
	}
}

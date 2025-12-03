package main

import "testing"

func TestReverseVowels(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{"example1", "IceCreAm", "AceCreIm"},
		{"example2", "leetcode", "leotcede"},
		{"empty", "", ""},
		{"singleVowel", "a", "a"},
		{"singleConsonant", "b", "b"},
		{"casePair", "Aa", "aA"},
		{"allVowels", "AEIOU", "UOIEA"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := reverseVowels(tc.in); got != tc.want {
				t.Fatalf("reverseVowels(%q) = %q, want %q", tc.in, got, tc.want)
			}
		})
	}
}

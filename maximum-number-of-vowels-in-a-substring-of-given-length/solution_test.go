package main

import "testing"

func TestMaxVowels(t *testing.T) {
	tests := []struct {
		name string
		s    string
		k    int
		want int
	}{
		{"example1", "abciiidef", 3, 3},
		{"example2", "aeiou", 2, 2},
		{"example3", "leetcode", 3, 2},
		{"no_vowels", "rhythms", 4, 0},
		{"single_vowel", "a", 1, 1},
		{"single_consonant", "b", 1, 0},
		{"all_vowels", "aeiou", 5, 5},
		{"mixed", "tryhard", 4, 1},
		{"window_all_vowels", "weallloveyou", 7, 4},
		{"repeated_vowels", "aaaa", 2, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxVowels(tt.s, tt.k)
			if got != tt.want {
				t.Fatalf("maxVowels(%q, %d) = %d; want %d", tt.s, tt.k, got, tt.want)
			}
		})
	}
}

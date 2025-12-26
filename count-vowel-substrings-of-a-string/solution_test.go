package main

import "testing"

func TestCountVowelSubstrings(t *testing.T) {
	tests := []struct {
		name string
		word string
		want int
	}{
		{"example1", "aeiouu", 2},
		{"example2", "unicornarihan", 0},
		{"example3", "cuaieuouac", 7},
		{"single_vowel", "a", 0},
		{"all_vowels_once", "aeiou", 1},
		{"no_vowels", "bcdfg", 0},
		{"repeated_vowels", "aaaaaa", 0},
		{"vowels_with_consonant", "aaeiouux", 4},
		{"multiple_segments", "aeioubcaeiou", 2},
		{"long_vowel_string", "aeiouaeiou", 21},
		{"minimal_valid", "aeiou", 1},
		{"extended_vowels", "aeioua", 3},
		{"all_five_repeated", "aeiouaeiou", 21},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countVowelSubstrings(tt.word)
			if got != tt.want {
				t.Fatalf("countVowelSubstrings(%q) = %d; want %d", tt.word, got, tt.want)
			}
		})
	}
}

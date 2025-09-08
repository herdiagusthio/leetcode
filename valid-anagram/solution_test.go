package main

import "testing"

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		name string
		s, t string
		want bool
	}{
		{"example", "anagram", "nagaram", true},
		{"not_anagram", "rat", "car", false},
		{"empty", "", "", true},
		{"single_diff", "a", "b", false},
		{"same_letters_diff_order", "listen", "silent", true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := isAnagram(tc.s, tc.t)
			if got != tc.want {
				t.Fatalf("%s: isAnagram(%q,%q) = %v, want %v", tc.name, tc.s, tc.t, got, tc.want)
			}
		})
	}
}

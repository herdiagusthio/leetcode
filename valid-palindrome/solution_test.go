package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{"empty", "", true},
		{"simple palindrome", "A man, a plan, a canal: Panama", true},
		{"not palindrome", "race a car", false},
		{"mixed chars", "0P", false},
		{"punct only", ".,", true},
		{"underscore treated non-alnum", "ab_a", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.s); got != tt.want {
				t.Fatalf("isPalindrome(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}

package main

import "testing"

func TestIsSubsequence(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want bool
	}{
		{"example1", "abc", "ahbgdc", true},
		{"example2", "axc", "ahbgdc", false},
		{"emptyS", "", "", true},
		{"emptySNonEmptyT", "", "abc", true},
		{"sLongerThanT", "abcd", "abc", false},
		{"nonContiguous", "ace", "abcde", true},
		{"orderMismatch", "aec", "abcde", false},
		{"sameStrings", "abc", "abc", true},
		{"singleCharTrue", "a", "a", true},
		{"singleCharFalse", "b", "a", false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := isSubsequence(tc.s, tc.t); got != tc.want {
				t.Fatalf("isSubsequence(%q, %q) = %v, want %v", tc.s, tc.t, got, tc.want)
			}
		})
	}
}

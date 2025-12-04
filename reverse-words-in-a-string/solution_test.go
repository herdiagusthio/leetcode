package main

import "testing"

func TestReverseWords(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{"simple", "the sky is blue", "blue is sky the"},
		{"leadingTrailing", "  hello world!  ", "world! hello"},
		{"multipleSpaces", "a good   example", "example good a"},
		{"empty", "", ""},
		{"spacesOnly", "    ", ""},
		{"single", "one", "one"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := reverseWords(tc.in); got != tc.want {
				t.Fatalf("reverseWords(%q) = %q, want %q", tc.in, got, tc.want)
			}
		})
	}
}

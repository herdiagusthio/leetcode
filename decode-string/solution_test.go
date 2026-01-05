package main

import (
	"testing"
)

func TestDecodeString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Example 1: Simple repetition",
			input:    "3[a]2[bc]",
			expected: "aaabcbc",
		},
		{
			name:     "Example 2: Nested encoding",
			input:    "3[a2[c]]",
			expected: "accaccacc",
		},
		{
			name:     "Example 3: Mixed with regular characters",
			input:    "2[abc]3[cd]ef",
			expected: "abcabccdcdcdef",
		},
		{
			name:     "Single character repetition",
			input:    "3[a]",
			expected: "aaa",
		},
		{
			name:     "No encoding",
			input:    "abcd",
			expected: "abcd",
		},
		{
			name:     "Double digit repetition",
			input:    "10[a]",
			expected: "aaaaaaaaaa",
		},
		{
			name:     "Triple digit repetition",
			input:    "100[a]",
			expected: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
		},
		{
			name:     "Multiple same level encodings",
			input:    "2[a]3[b]4[c]",
			expected: "aabbbcccc",
		},
		{
			name:     "Deep nesting",
			input:    "2[2[2[a]]]",
			expected: "aaaaaaaa",
		},
		{
			name:     "Complex nested with letters",
			input:    "2[abc3[cd]xyz]",
			expected: "abccdcdcdxyzabccdcdcdxyz",
		},
		{
			name:     "Leading characters",
			input:    "abc3[cd]",
			expected: "abccdcdcd",
		},
		{
			name:     "Trailing characters",
			input:    "3[cd]xyz",
			expected: "cdcdcdxyz",
		},
		{
			name:     "Characters between encodings",
			input:    "2[a]b3[c]",
			expected: "aabccc",
		},
		{
			name:     "Single digit single char",
			input:    "1[a]",
			expected: "a",
		},
		{
			name:     "Nested at different levels",
			input:    "3[z2[xy]]",
			expected: "zxyxyzxyxyzxyxy",
		},
		{
			name:     "Multiple nested brackets",
			input:    "2[2[y]pq4[2[jk]e1[f]]]",
			expected: "yypqjkjkefjkjkefjkjkefjkjkefyypqjkjkefjkjkefjkjkefjkjkef",
		},
		{
			name:     "Large number repetition",
			input:    "20[a]",
			expected: "aaaaaaaaaaaaaaaaaaaa",
		},
		{
			name:     "Mixed complexity",
			input:    "ab2[c3[d]e]f",
			expected: "abcdddecdddef",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := decodeString(tt.input)
			if result != tt.expected {
				t.Errorf("decodeString(%q) = %q, expected %q", tt.input, result, tt.expected)
			}
		})
	}
}

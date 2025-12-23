package main

import (
	"testing"
)

func TestRemoveStars(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"leet**cod*e", "lecoe"},
		{"erase*****", ""},
		{"abc", "abc"},
		{"a*b*c*", ""},
	}

	for _, test := range tests {
		t.Run("Stack_"+test.input, func(t *testing.T) {
			result := removeStars(test.input)
			if result != test.expected {
				t.Errorf("expected %s, got %s", test.expected, result)
			}
		})

		t.Run("Backward_"+test.input, func(t *testing.T) {
			result := removeStarsBackward(test.input)
			if result != test.expected {
				t.Errorf("expected %s, got %s", test.expected, result)
			}
		})
	}
}

func BenchmarkRemoveStars(b *testing.B) {
	input := "leet**cod*e"
	for i := 0; i < b.N; i++ {
		removeStars(input)
	}
}

func BenchmarkRemoveStarsBackward(b *testing.B) {
	input := "leet**cod*e"
	for i := 0; i < b.N; i++ {
		removeStarsBackward(input)
	}
}

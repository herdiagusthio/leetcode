package main

import (
	"reflect"
	"testing"
)

func TestCompress(t *testing.T) {
	tests := []struct {
		name     string
		chars    []byte
		expected []byte
		wantLen  int
	}{
		{
			name:     "Example 1",
			chars:    []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'},
			expected: []byte{'a', '2', 'b', '2', 'c', '3'},
			wantLen:  6,
		},
		{
			name:     "Example 2",
			chars:    []byte{'a'},
			expected: []byte{'a'},
			wantLen:  1,
		},
		{
			name:     "Example 3",
			chars:    []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'},
			expected: []byte{'a', 'b', '1', '2'},
			wantLen:  4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Make a copy of chars to avoid modifying the test case input for other runs if we were reusing it (though here we don't)
			input := make([]byte, len(tt.chars))
			copy(input, tt.chars)

			gotLen := compress(input)

			if gotLen != tt.wantLen {
				t.Errorf("compress() = %v, want %v", gotLen, tt.wantLen)
			}

			if !reflect.DeepEqual(input[:gotLen], tt.expected) {
				t.Errorf("compress() buffer = %v, want %v", input[:gotLen], tt.expected)
			}
		})
	}
}

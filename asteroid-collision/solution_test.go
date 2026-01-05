package main

import (
	"reflect"
	"testing"
)

func TestAsteroidCollision(t *testing.T) {
	tests := []struct {
		name      string
		asteroids []int
		expected  []int
	}{
		{
			name:      "Example 1: Right wins collision",
			asteroids: []int{5, 10, -5},
			expected:  []int{5, 10},
		},
		{
			name:      "Example 2: Equal size collision",
			asteroids: []int{8, -8},
			expected:  []int{},
		},
		{
			name:      "Example 3: Multiple collisions",
			asteroids: []int{10, 2, -5},
			expected:  []int{10},
		},
		{
			name:      "Example 4: Complex scenario",
			asteroids: []int{-2, -1, 1, 2},
			expected:  []int{-2, -1, 1, 2},
		},
		{
			name:      "All moving left",
			asteroids: []int{-5, -10, -15},
			expected:  []int{-5, -10, -15},
		},
		{
			name:      "All moving right",
			asteroids: []int{5, 10, 15},
			expected:  []int{5, 10, 15},
		},
		{
			name:      "Left destroys all right",
			asteroids: []int{1, 2, 3, -10},
			expected:  []int{-10},
		},
		{
			name:      "Multiple left destroy multiple right",
			asteroids: []int{5, 10, -5, -10},
			expected:  []int{5},
		},
		{
			name:      "Chain collisions",
			asteroids: []int{10, 5, -5, -10},
			expected:  []int{},
		},
		{
			name:      "Left then right then left",
			asteroids: []int{-5, 10, -15},
			expected:  []int{-5, -15},
		},
		{
			name:      "Single asteroid positive",
			asteroids: []int{10},
			expected:  []int{10},
		},
		{
			name:      "Single asteroid negative",
			asteroids: []int{-10},
			expected:  []int{-10},
		},
		{
			name:      "Two moving same direction right",
			asteroids: []int{5, 10},
			expected:  []int{5, 10},
		},
		{
			name:      "Two moving same direction left",
			asteroids: []int{-10, -5},
			expected:  []int{-10, -5},
		},
		{
			name:      "Large left destroys small right",
			asteroids: []int{1, -2},
			expected:  []int{-2},
		},
		{
			name:      "Small left destroyed by large right",
			asteroids: []int{10, -1},
			expected:  []int{10},
		},
		{
			name:      "Multiple equal collisions",
			asteroids: []int{5, -5, 10, -10},
			expected:  []int{},
		},
		{
			name:      "Complex with survivors on both sides",
			asteroids: []int{-5, -3, 1, 2, -2, 3, -10},
			expected:  []int{-5, -3, -10},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := asteroidCollision(tt.asteroids)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("asteroidCollision(%v) = %v, expected %v", tt.asteroids, result, tt.expected)
			}
		})
	}
}

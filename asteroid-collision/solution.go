package main

import "fmt"

// asteroidCollision simulates asteroid collisions in a row.
// Positive values represent asteroids moving right, negative values moving left.
// When two asteroids collide, the smaller one explodes. If both are equal, both explode.
// Asteroids moving in the same direction never collide.
// Returns the final state after all collisions.
func asteroidCollision(asteroids []int) []int {
    var (
        result []int
    )    

    outer_loop:
    for _, asteroid := range asteroids {
        // No collision: empty stack, top is left-moving, or current is right-moving
        if len(result) == 0 || result[len(result)-1] < 0 || asteroid > 0 {
            result = append(result, asteroid)
            continue
        }
        
        // Handle collisions: right-moving asteroids in stack meet left-moving current
        for len(result) > 0 && result[len(result)-1] > 0 {
            lastAsteroid := result[len(result)-1]
            
            if lastAsteroid == -asteroid {
                // Equal size: both asteroids explode
                result = result[:len(result)-1]
                continue outer_loop
            } else if lastAsteroid > -asteroid {
                // Right-moving asteroid wins: current explodes
                continue outer_loop
            } else {
                // Left-moving asteroid wins: pop top and continue checking
                result = result[:len(result)-1]
            }
        }

        // Current asteroid survived all collisions
        result = append(result, asteroid)

    }

    return result
}

func main() {
	// Example 1: 10 and -5 collide, 10 wins
	fmt.Println(asteroidCollision([]int{5, 10, -5})) // [5 10]

	// Example 2: 8 and -8 collide, both explode
	fmt.Println(asteroidCollision([]int{8, -8})) // []

	// Example 3: Multiple collisions
	fmt.Println(asteroidCollision([]int{10, 2, -5})) // [10]

	// Example 4: Complex scenario with multiple survivors
	fmt.Println(asteroidCollision([]int{-2, -1, 1, 2})) // [-2 -1 1 2]
}
package main

import "fmt"

// largestAltitude returns the highest altitude of a point during a road trip.
func largestAltitude(gain []int) int {
	highestAltitude, altitude := 0, 0
	for _, g := range gain {
		altitude += g

		if altitude > highestAltitude {
			highestAltitude = altitude
		}
	}

	return highestAltitude
}

func main() {
	examples := []struct {
		gain []int
	}{
		{[]int{-5, 1, 5, 0, -7}},
		{[]int{-4, -3, -2, -1, 4, 3, 2}},
	}

	for _, ex := range examples {
		fmt.Printf("largestAltitude(%v) = %d\n", ex.gain, largestAltitude(ex.gain))
	}
}

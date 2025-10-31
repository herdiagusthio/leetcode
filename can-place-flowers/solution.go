package main

import "fmt"

func canPlaceFlowers(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	}

	count := 0
	length := len(flowerbed)

	for i := 0; i < length; i++ {
		// Check if we can place a flower here: current spot empty, and no adjacent flowers
		if flowerbed[i] == 0 &&
			(i == 0 || flowerbed[i-1] == 0) && // No flower to the left (or out of bounds)
			(i == length-1 || flowerbed[i+1] == 0) { // No flower to the right (or out of bounds)
			// Place flower (simulate without mutating the array)
			count++
			if count >= n {
				return true // Early exit if we've placed enough
			}
			i++ // Skip next position to avoid adjacent placement
		}
	}

	return count >= n
}

func main() {
	samples := []struct {
		bed []int
		n   int
	}{
		{[]int{1, 0, 0, 0, 1}, 1},
		{[]int{1, 0, 0, 0, 1}, 2},
		{[]int{0}, 1},
		{[]int{0, 0, 1, 0, 0}, 2},
	}

	for _, s := range samples {
		fmt.Printf("bed=%v n=%d -> %v\n", s.bed, s.n, canPlaceFlowers(append([]int(nil), s.bed...), s.n))
	}
}

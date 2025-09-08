package main

import (
	"fmt"
)

func findClosestNumber(nums []int) int {
    if len(nums) == 0 {return 0}

    closestDistance := absInt(nums[0])
    largestValue := nums[0]

    for _, num := range nums[1:] {
        distance := absInt(num)

        if (distance < closestDistance) || (distance == closestDistance && num > largestValue) {
            closestDistance = distance
            largestValue = num
        }
    }

    return largestValue    
}

func absInt(x int) int {
    if x < 0 {return -x}
    return x
}

func main() {
	nums := []int{-4,-1, 1, 4, 8}
	result := findClosestNumber(nums)
	fmt.Println(result) // Output: 1
}
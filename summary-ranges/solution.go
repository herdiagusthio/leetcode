package main

import (
	"strconv"
)

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return nil
	}

	output := make([]string, 0, len(nums))
	start := nums[0]
	end := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i]-end == 1 {
			end = nums[i]
		} else {
			if start == end {
				output = append(output, strconv.Itoa(start))
			} else {
				output = append(output, strconv.Itoa(start)+"->"+strconv.Itoa(end))
			}
			start = nums[i]
			end = nums[i]
		}
	}

	if start == end {
		output = append(output, strconv.Itoa(start))
	} else {
		output = append(output, strconv.Itoa(start)+"->"+strconv.Itoa(end))
	}

	return output
}

# Max Consecutive Ones III

**LeetCode problem:**  
https://leetcode.com/problems/max-consecutive-ones-iii/

Given a binary array `nums` and an integer `k`, return the maximum number of consecutive 1's in the array if you can flip at most `k` 0's.

**Example 1:**
```
Input: nums = [1,1,1,0,0,0,1,1,1,1,0], k = 2
Output: 6
Explanation: [1,1,1,0,0,1,1,1,1,1,1]
Bolded numbers were flipped from 0 to 1. The longest subarray is underlined.
```

**Example 2:**
```
Input: nums = [0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1], k = 3
Output: 10
Explanation: [0,0,1,1,1,1,1,1,1,1,1,1,0,0,0,1,1,1,1]
Bolded numbers were flipped from 0 to 1. The longest subarray is underlined.
```

**Constraints:**
- `1 <= nums.length <= 10^5`
- `nums[i]` is either `0` or `1`
- `0 <= k <= nums.length`

## Solution

The solution uses a sliding window approach with two pointers to find the longest subarray containing at most `k` zeros. The window expands by moving the right pointer and shrinks from the left when the zero count exceeds `k`. This efficiently tracks the maximum consecutive 1's achievable by flipping at most `k` zeros.

**Time Complexity:** O(n) where n is the length of the nums array  
**Space Complexity:** O(1)

See [EXPLANATION.md](./EXPLANATION.md) for a detailed breakdown of the algorithm.

## How to run

```bash
# Run tests
go test -v

# Run the solution with example inputs
go run .
```

## Files

- `solution.go` - Main solution implementation with sliding window approach
- `solution_test.go` - Table-driven tests covering examples and edge cases
- `EXPLANATION.md` - Detailed explanation of the algorithm and approach
- `README.md` - This file
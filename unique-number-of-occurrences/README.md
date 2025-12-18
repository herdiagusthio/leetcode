# Unique Number of Occurrences

**LeetCode problem:**  
https://leetcode.com/problems/unique-number-of-occurrences/

Given an array of integers `arr`, return true if the number of occurrences of each value in the array is unique or false otherwise.

**Example 1:**
```
Input: arr = [1,2,2,1,1,3]
Output: true
Explanation: The value 1 has 3 occurrences, 2 has 2 and 3 has 1. No two values have the same number of occurrences.
```

**Example 2:**
```
Input: arr = [1,2]
Output: false
```

**Example 3:**
```
Input: arr = [-3,0,1,-3,1,1,1,-3,10,0]
Output: true
```

**Constraints:**
- `1 <= arr.length <= 1000`
- `-1000 <= arr[i] <= 1000`

## Solution

The solution uses a two-pass approach: first counting the occurrences of each value using a hash map, then checking if all occurrence counts are unique using a boolean set. The optimized version uses `map[int]bool` instead of `map[int]int` for better memory efficiency and cleaner semantics.

**Time Complexity:** O(n) where n is the length of the array  
**Space Complexity:** O(n)

See [EXPLANATION.md](./EXPLANATION.md) for a detailed breakdown of the algorithm.

## How to run

```bash
# Run tests
go test -v

# Run the solution with example inputs
go run .
```

## Files

- `solution.go` - Main solution implementation with two-pass hash map approach
- `solution_test.go` - Table-driven tests covering examples and edge cases
- `EXPLANATION.md` - Detailed explanation of the algorithm and approach
- `README.md` - This file
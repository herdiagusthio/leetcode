# Find the Highest Altitude

**LeetCode problem:**  
https://leetcode.com/problems/find-the-highest-altitude/

There is a biker going on a road trip. The road trip consists of n + 1 points at different altitudes. The biker starts his trip on point 0 with altitude equal 0.

You are given an integer array `gain` of length `n` where `gain[i]` is the net gain in altitude between points `i` and `i + 1` for all (0 <= i < n). Return the highest altitude of a point.

**Example 1:**
```
Input: gain = [-5,1,5,0,-7]
Output: 1
Explanation: The altitudes are [0,-5,-4,1,1,-6]. The highest is 1.
```

**Example 2:**
```
Input: gain = [-4,-3,-2,-1,4,3,2]
Output: 0
Explanation: The altitudes are [0,-4,-7,-9,-10,-6,-3,-1]. The highest is 0.
```

**Constraints:**
- `n == gain.length`
- `1 <= n <= 100`
- `-100 <= gain[i] <= 100`

## Solution

The solution uses a simple iterative approach to calculate cumulative altitude at each point while tracking the maximum altitude encountered. Starting from altitude 0, we add each gain value and compare with the highest altitude seen so far.

**Time Complexity:** O(n) where n is the length of the gain array  
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

- `solution.go` - Main solution implementation with iterative approach
- `solution_test.go` - Table-driven tests covering examples and edge cases
- `EXPLANATION.md` - Detailed explanation of the algorithm and approach
- `README.md` - This file
# Longest Subarray of 1's After Deleting One Element

**LeetCode problem:**  
https://leetcode.com/problems/longest-subarray-of-1s-after-deleting-one-element/

Given a binary array `nums`, you should delete one element from it.

Return the size of the longest non-empty subarray containing only 1's in the resulting array. Return 0 if there is no such subarray.

**Example 1:**
```
Input: nums = [1,1,0,1]
Output: 3
Explanation: After deleting the number in position 2, [1,1,1] contains 3 numbers with value of 1's.
```

**Example 2:**
```
Input: nums = [0,1,1,1,0,1,1,0,1]
Output: 5
Explanation: After deleting the number in position 4, [0,1,1,1,1,1,0,1] longest subarray with value of 1's is [1,1,1,1,1].
```

**Example 3:**
```
Input: nums = [1,1,1]
Output: 2
Explanation: You must delete one element.
```

**Constraints:**
- `1 <= nums.length <= 10^5`
- `nums[i]` is either `0` or `1`

## Solution

The solution uses a sliding window approach to find the longest subarray with at most 1 zero. The key insight is that we must delete exactly one element, so finding a window with at most 1 zero allows us to either delete that zero (connecting 1's) or delete a 1 (if all are 1's). The length is calculated as `right - left` to account for the mandatory deletion.

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
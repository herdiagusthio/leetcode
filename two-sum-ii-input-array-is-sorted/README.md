# Two Sum II - Input Array Is Sorted

LeetCode problem: Two Sum II - Input Array Is Sorted

Link: https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/

Given a 1-indexed array of integers `numbers` that is already sorted in non-decreasing order, find two numbers such that they add up to a specific `target` number. Return the indices of the two numbers (1-based) as an integer array `[index1, index2]` of length 2.

## Solution
This package implements the classic two-pointer approach for sorted arrays:
1. Maintain two pointers, `left` at start and `right` at end.
2. Compute sum = numbers[left] + numbers[right].
   - If sum == target: return [left+1, right+1] (1-based indices).
   - If sum < target: advance left++ to increase sum.
   - If sum > target: decrement right-- to decrease sum.

This is O(n) time and O(1) extra space.

## How to run
From this folder:

```bash
go test -v
```

To run the example `main` in this package:

```bash
go run .
```

## Files
- `solution.go` — two-pointer `twoSum` function and small demo `main()`.
- `solution_test.go` — unit tests.
- `EXPLANATION.md` — step-by-step explanation of the approach.

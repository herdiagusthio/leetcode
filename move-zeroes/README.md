# Move Zeroes

LeetCode problem: Move Zeroes

Link: https://leetcode.com/problems/move-zeroes/

Given an integer array `nums`, move all 0's to the end of it while maintaining the relative order of the non-zero elements. You must do this in-place without making a copy of the array.

## Solution
This package implements a two-pointer solution that scans the array once, swapping non-zero elements to the front while maintaining their relative order. The approach is O(n) time and O(1) space.

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
- `solution.go` — implementation and example `main`.
- `solution_test.go` — unit tests.
- `EXPLANATION.md` — step-by-step explanation of the approach.

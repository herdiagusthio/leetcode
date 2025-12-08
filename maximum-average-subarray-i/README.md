# Maximum Average Subarray I

LeetCode problem: Maximum Average Subarray I

Link: https://leetcode.com/problems/maximum-average-subarray-i/

Given an integer array `nums` of `n` elements and an integer `k`, find a contiguous subarray
whose length is equal to `k` that has the maximum average value and return this value.

## Solution
This package implements the optimal sliding window approach. We compute the sum of the first
`k` elements, then slide the window by adding the next element and removing the leftmost element.
We track the maximum sum and return `maxSum / k` as the maximum average. This runs in O(n) time
with O(1) extra space.

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
- `solution_test.go` — table-driven unit tests covering edge cases.
- `EXPLANATION.md` — step-by-step explanation of the approach.
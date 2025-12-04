# Find Closest Number to Zero

LeetCode problem: Find Closest Number to Zero

Link: https://leetcode.com/problems/find-closest-number-to-zero/description/

Given an integer array, return the number that is closest to zero. If two numbers are equally close, return the larger one (for example, between `-3` and `3` return `3`).

## Solution
The implementation scans the slice once and keeps the current best candidate: the element with the smallest absolute value seen so far. When a new element has a strictly smaller absolute value it becomes the candidate. If the absolute values are equal, the larger numerical value is chosen (tie-break rule).

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


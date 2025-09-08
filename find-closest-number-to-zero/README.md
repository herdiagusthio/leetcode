# Find Closest Number to Zero

LeetCode problem: https://leetcode.com/problems/find-closest-number-to-zero/description/

## Problem
Given an integer array, return the number that is closest to zero. If two numbers are equally close, return the larger one (for example, between `-3` and `3` return `3`).

## Solution (short)
The implementation scans the slice once and keeps the current best candidate: the element with the smallest absolute value seen so far. When a new element has a strictly smaller absolute value it becomes the candidate. If the absolute values are equal, the larger numerical value is chosen (tie-break rule).

This is implemented in `solution.go` with an integer `absInt` helper to avoid float conversions.

## Complexity
- Time: O(n) — single pass through the input slice.
- Space: O(1) — constant extra memory.

## Example
Input (in `main` of `solution.go`):

```go
nums := []int{-4, -1, 1, 4, 8}
```

Output:

```
1
```

Explanation: both `-1` and `1` are distance 1 from zero; the tie-breaker picks `1`.

## How to run
- Run the sample program:

```bash
go run solution.go
```

## Run tests

This repository includes unit tests under the package. Choose one of the following depending on how you run Go:

- Module mode (recommended, run tests for the package):

```bash
go test ./find-closest-number-to-zero
```

- Module mode (run all packages under the repo root):

```bash
go test ./...
```

- GOPATH mode (disable modules):

```bash
cd find-closest-number-to-zero
GO111MODULE=off go test
```

## Notes
- Go version: the code targets Go 1.20+ but will work on most recent Go toolchains.
- Contributions and PRs are welcome. If you want the explanation moved to a separate `EXPLANATION.md`, I can add it.


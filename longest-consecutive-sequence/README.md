# Longest Consecutive Sequence

LeetCode problem: Longest Consecutive Sequence

Link: https://leetcode.com/problems/longest-consecutive-sequence

Given an unsorted array of integers, find the length of the longest consecutive elements sequence.

## Solution
This folder contains an O(n) solution using a hash set (`map[int]struct{}`) to detect sequence starts and expand consecutive runs.

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
- `solution.go` — optimized implementation of `longestConsecutive`.
- `solution_test.go` — unit tests and a benchmark.
- `EXPLANATION.md` — step-by-step reasoning for the implementation.
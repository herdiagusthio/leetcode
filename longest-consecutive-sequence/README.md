# Longest Consecutive Sequence

LeetCode problem: Longest Consecutive Sequence

Given an unsorted array of integers, find the length of the longest consecutive elements sequence.

- Link: https://leetcode.com/problems/longest-consecutive/

## Solution
This folder contains an O(n) solution using a hash set (`map[int]struct{}`) to detect sequence starts and expand consecutive runs.

## How to run
From this folder, run:

```bash
go test -v
```

Or from repo root:

```bash
go test ./longest-consecutive-sequence -v
```

Benchmarks:

```bash
go test -bench .
```

## Files
- `solution.go` — optimized implementation of `longestConsecutive`.
- `solution_test.go` — unit tests and a benchmark.
- `EXPLANATION.md` — step-by-step reasoning for the implementation.
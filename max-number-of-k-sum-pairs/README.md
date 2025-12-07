# Max Number of K-Sum Pairs

LeetCode problem: Max Number of K-Sum Pairs

Link: https://leetcode.com/problems/max-number-of-k-sum-pairs/

Given an integer array `nums` and an integer `k`, you may perform an operation that removes
any pair of numbers whose sum equals `k`. Return the maximum number of such operations.

## Solution
This package implements the standard greedy hash-map approach. We iterate once through `nums`,
maintaining a frequency map of values still waiting for a complement. For each value `v`, we
look for `k - v` in the map; if present, we form a pair and decrement the stored count. Otherwise
we record `v` for future matches. This ensures every element is used at most once, producing the
maximum count of disjoint pairs in O(n) time and O(u) additional space (`u` distinct values).

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

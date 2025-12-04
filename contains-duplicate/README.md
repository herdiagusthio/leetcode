# Contains Duplicate

LeetCode problem: Contains Duplicate

Link: https://leetcode.com/problems/contains-duplicate/

Given an integer array `nums`, return `true` if any value appears at least twice in the array, and return `false` if every element is distinct.

## Solution
This package contains a simple O(n) solution using a hash set (`map[int]struct{}`) to track seen values and return early on the first duplicate.

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
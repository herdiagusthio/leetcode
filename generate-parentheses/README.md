# Generate Parentheses

LeetCode problem: Generate Parentheses

Link: https://leetcode.com/problems/generate-parentheses/

Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

## Solution
Use DFS/backtracking while tracking how many open and close parentheses have been placed. Collect the sequence once its length reaches 2*n.

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

## Generate Parentheses

LeetCode problem: https://leetcode.com/problems/generate-parentheses/

## Problem

Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

## Solution (short)

Use DFS/backtracking while tracking how many open and close parentheses have been placed. At each step:

- add '(' if open < n
- add ')' if close < open

Collect the sequence once its length reaches 2*n.

This folder contains the implementation using backtracking and a small demo `main`.

## Complexity

- Time: O(C_n * n) where C_n is the nth Catalan number (number of valid sequences).
- Space: O(C_n * n) for the output, plus O(n) recursion depth.

## Example

Input:

```go
generateParenthesis(3)
```

Output:

```
["((()))","(()())","(())()","()(())","()()()"]
```

## How to run

- Run the sample program (prints example output):

```bash
cd generate-parentheses
go run .
```

- Run tests for this package only:

```bash
go test ./generate-parentheses
```

- Run all tests in the repository:

```bash
go test ./...
```

If you run without Go modules (GOPATH):

```bash
cd generate-parentheses
GO111MODULE=off go test
```

## Files in this folder

- `solution.go` — implementation and small example `main`.
- `solution_test.go` — unit tests.
- `EXPLANATION.md` — detailed explanation and notes.

## Notes

- The backtracking implementation is clear and preserves correctness; use an in-place buffer variant if allocation pressure becomes an issue.

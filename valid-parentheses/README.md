# Valid Parentheses

LeetCode problem: Valid Parentheses

Link: https://leetcode.com/problems/valid-parentheses/

Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:
- Open brackets must be closed by the same type of brackets.
- Open brackets must be closed in the correct order.
- Every close bracket has a corresponding open bracket of the same type.

## Solution
Scan the string once using a stack (LIFO). Push opening brackets. When encountering a closing bracket, check the stack top matches the corresponding opening bracket and pop it. If any mismatch occurs or the stack is non-empty at the end, the string is invalid.

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
- `solution_test_valid.go` — unit tests.
- `EXPLANATION.md` — step-by-step explanation of the approach.

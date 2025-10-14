## Valid Parentheses

LeetCode problem: https://leetcode.com/problems/valid-parentheses/

## Problem

Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

- Open brackets must be closed by the same type of brackets.
- Open brackets must be closed in the correct order.
- Every close bracket has a corresponding open bracket of the same type.

## Solution (short)

Scan the string once using a stack (LIFO). Push opening brackets. When encountering a closing bracket, check the stack top matches the corresponding opening bracket and pop it. If any mismatch occurs or the stack is non-empty at the end, the string is invalid.

## Complexity

- Time: O(n) — single pass over the string.
- Space: O(n) — stack in the worst case.

## Example

Input: s = "()[]{}"

Output: true

## How to run

- Run the sample program in this folder:

```bash
cd valid-parentheses
go run .
```

- Run tests for this package only:

```bash
go test ./valid-parentheses
```

## Files in this folder

- `solution.go` — implementation and small demo `main`.
- `solution_test_valid.go` — unit tests.
- `EXPLANATION.md` — explanation and notes.

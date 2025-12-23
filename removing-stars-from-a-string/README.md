# Removing Stars From a String

LeetCode problem: Removing Stars From a String

Link: https://leetcode.com/problems/removing-stars-from-a-string/description/

You are given a string s, which contains stars *.

In one operation, you can:
- Choose a star in s.
- Remove the closest non-star character to its left, as well as remove the star itself.

Return the string after all stars have been removed.

Note:
- The input will be generated such that the operation is always possible.
- It can be shown that the resulting string will always be unique.

## Solution
We use a stack-based approach (simulated with a byte slice). We iterate through the string character by character. If we encounter a star '*', we pop the last character from our stack. If it's a non-star character, we push it onto the stack.

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
- `solution.go` — implementation.
- `solution_test.go` — unit tests.
- `EXPLANATION.md` — step-by-step explanation of the approach.
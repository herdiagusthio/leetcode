# Reverse Vowels of a String

LeetCode problem: Reverse Vowels of a String

Link: https://leetcode.com/problems/reverse-vowels-of-a-string/

Given a string `s`, reverse only the vowels in the string and return it.

## Solution
This package provides a simple two-pointer solution that converts the input to a `[]rune`, then swaps vowels from both ends until the pointers meet. The approach is correct for multi-byte UTF-8 characters and runs in linear time.

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

# Valid Palindrome

LeetCode problem: Valid Palindrome

Link: https://leetcode.com/problems/valid-palindrome/

Given a string s, return true if it is a palindrome after removing all non-alphanumeric characters and ignoring case. Alphanumeric characters include letters and numbers.

## Solution
This implementation follows three simple steps:
1. Remove non-alphanumeric characters using a precompiled regular expression.
2. Convert the cleaned string to lowercase.
3. Use a two-pointer comparison (front and back) to check for palindrome.

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

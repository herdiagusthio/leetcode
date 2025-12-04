# Reverse Words in a String

LeetCode problem: Reverse Words in a String

Link: https://leetcode.com/problems/reverse-words-in-a-string/

Given an input string `s`, reverse the order of the words. Words are separated by spaces. The returned string should have words in reverse order, separated by a single space, with no leading or trailing spaces.

## Solution
This package provides a straightforward solution using `strings.Fields` to split the input into words (which trims and collapses whitespace), reversing the words in-place, and joining them with a single space using `strings.Join`.

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
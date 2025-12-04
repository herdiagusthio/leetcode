# Valid Anagram

LeetCode problem: Valid Anagram

Link: https://leetcode.com/problems/valid-anagram/description/

Given two strings `s` and `t`, return true if `t` is an anagram of `s`, and false otherwise. Both strings consist of lowercase English letters.

## Solution
The provided `isAnagram` implementation counts characters in `s` into a map (or array), then decrements counts while scanning `t`; if any count goes negative it returns false. This yields O(n) time and O(1) space relative to the alphabet size.

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

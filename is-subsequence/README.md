# Is Subsequence

LeetCode problem: Is Subsequence

Link: https://leetcode.com/problems/is-subsequence/

Given two strings `s` and `t`, determine if `s` is a subsequence of `t`. A subsequence means the characters of `s` appear in `t` in the same order (not necessarily contiguous).

## Solution
This package implements a linear two-pointer scan: iterate `t` and advance a pointer over `s` when characters match. If the `s` pointer reaches the end, `s` is a subsequence of `t`. Time is O(n) where n = len(t), and extra space is O(1).

## How to run
# Is Subsequence

LeetCode problem: Is Subsequence

Link: https://leetcode.com/problems/is-subsequence/

Given two strings `s` and `t`, determine if `s` is a subsequence of `t`. A subsequence means the characters of `s` appear in `t` in the same order (not necessarily contiguous).

## Solution
This package implements a linear two-pointer scan: iterate `t` and advance a pointer over `s` when characters match. If the `s` pointer reaches the end, `s` is a subsequence of `t`. Time is O(n) where n = len(t), and extra space is O(1).

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

## Problem details

Constraints (from LeetCode):

- 0 <= s.length <= 100
- 0 <= t.length <= 10^4
- `s` and `t` consist only of lowercase English letters.

Example:

- `s = "abc"`, `t = "ahbgdc"` → `true`
- `s = "axc"`, `t = "ahbgdc"` → `false`



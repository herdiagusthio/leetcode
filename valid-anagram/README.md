# Valid Anagram

LeetCode problem: https://leetcode.com/problems/valid-anagram/description/

## Problem
Given two strings `s` and `t`, return true if `t` is an anagram of `s`, and false otherwise. Both strings consist of lowercase English letters.

## How the solution works (short)
The provided `isAnagram` implementation counts characters in `s` into a map, then decrements counts while scanning `t`; if any count goes negative it returns false. This yields O(n) time and O(1) space relative to the alphabet size.

## Complexity
- Time: O(n) — single pass to count and single pass to verify.
- Space: O(1) — counts limited to lowercase English letters (effectively constant).

## How to run

```bash
go run solution.go
```

## Run tests

```bash
go test
```

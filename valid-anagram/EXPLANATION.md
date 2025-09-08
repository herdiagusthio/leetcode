# Explanation — Valid Anagram

This document explains the `isAnagram` implementation in `solution.go`.

## Contract
- Input: `s`, `t` strings consisting of lowercase English letters.
- Output: `bool` — true if `t` is an anagram of `s`, false otherwise.

## High-level approach
Count characters of `s` and decrement counts for `t`. If any count becomes negative, `t` contains more of a character than `s` and the function returns false. If all counts match, the strings are anagrams.

## Why this is efficient
- Time: single pass over each string → O(n).
- Space: counts limited to the alphabet size (26) → O(1).

## Edge cases
- Different lengths → immediately false.
- Empty strings of equal length → true.

## Suggestion
Given the constraints (lowercase English letters only), a fixed-size array (`[26]int`) is slightly faster and uses no map allocations. The current map-based implementation is correct and more general (handles Unicode) but uses small extra overhead.

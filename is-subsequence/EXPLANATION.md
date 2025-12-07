# Explanation — Is Subsequence

## Problem
Given strings `s` and `t`, check whether `s` is a subsequence of `t`. A subsequence means that all characters in `s` appear in `t` in the same order (not necessarily contiguous).

Examples:
- `s = "abc"`, `t = "ahbgdc"` -> `true`
- `s = "axc"`, `t = "ahbgdc"` -> `false`

## Idea
Use a two-pointer scan: keep a pointer over `s` and iterate through `t`. When `t[i]` equals the current `s` character, advance the `s` pointer. If the `s` pointer reaches `len(s)`, all characters were matched in order and `s` is a subsequence.

## Algorithm (step-by-step)
1. If `len(s) == 0`, return `true` (empty string is a subsequence of any string).
2. Initialize `iS := 0`.
3. For each index `i` in `0..len(t)-1`:
   - If `t[i] == s[iS]`, increment `iS`.
   - If `iS == len(s)` return `true`.
4. Return `false` if we finish the loop without matching all characters.

## Complexity
- Time: O(len(t)).
- Space: O(1) extra.

## Alternatives
- If many queries are made with the same `t` and different `s` values, precompute next-character jump tables (DP/next pointers) to answer queries faster.

## Implementation notes
- The simple linear scan is optimal for single-query cases and is easy to reason about and test.

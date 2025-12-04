# Explanation — Longest Consecutive Sequence

## Problem
Given an unsorted array of integers, return the length of the longest sequence of consecutive integers.

## Idea
We want an O(n) solution. Sorting would be O(n log n), so instead we use a hash set to get O(1) lookups and detect the starts of sequences.

## Algorithm (step-by-step)
1. Edge case: if the input slice is empty, return 0.
2. Insert every number into a set (`map[int]struct{}`) to allow O(1) membership checks.
3. For each number `n` in the set:
   - If `n-1` exists in the set then `n` is not the start of a sequence, skip it.
   - Otherwise `n` is the start of a sequence. Expand forward by checking `n+1`, `n+2`, ... until the next integer is absent.
   - Track the length of the expanded run and update the global maximum.
4. Return the maximum run length found.

## Complexity
- Time: O(n) average — every number is inserted once and visited at most twice (once as a set element, once during expansion).
- Space: O(n) for the set.

## Implementation notes
- We only expand sequences from numbers that do not have a predecessor. This guarantees each consecutive sequence is expanded exactly once.
- We use `map[int]struct{}` as a set to minimize memory overhead compared to `map[int]bool`.

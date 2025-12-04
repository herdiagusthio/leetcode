# Explanation — Valid Anagram

## Problem
Given two strings `s` and `t`, return true if `t` is an anagram of `s`, and false otherwise. Both strings consist of lowercase English letters.

## Idea
Count characters of `s` and decrement counts for `t`. If any count becomes negative, `t` contains more of a character than `s` and the function returns false. If all counts match, the strings are anagrams.

## Algorithm (step-by-step)
1. Check if lengths of `s` and `t` are different. If so, return `false`.
2. Initialize a frequency map (or array of size 26).
3. Iterate through `s` and increment counts for each character.
4. Iterate through `t` and decrement counts for each character.
   - If a count drops below zero, return `false`.
5. Return `true`.

## Complexity
- Time: O(n) — single pass over each string.
- Space: O(1) — counts limited to the alphabet size (26).

## Alternatives
- **Sorting**: Sort both strings and compare them. Time O(n log n), Space O(1) or O(n) depending on sort implementation.

## Implementation notes
- Different lengths → immediately false.
- Empty strings of equal length → true.
- Given the constraints (lowercase English letters only), a fixed-size array (`[26]int`) is slightly faster and uses no map allocations.

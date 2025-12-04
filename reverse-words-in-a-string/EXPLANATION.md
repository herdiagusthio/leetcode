# Explanation — Reverse Words in a String

## Problem
Given an input string `s`, reverse the order of the words. A word is defined as a sequence of non-space characters. The returned string should have the words in reverse order, separated by a single space, and no leading or trailing spaces.

## Idea
Split the input into words using `strings.Fields`, which trims leading/trailing whitespace and collapses multiple spaces. Then reverse the slice of words and join them with a single space.

This approach is simple and matches the problem requirement of trimming and collapsing whitespace.

## Algorithm (step-by-step)
1. Call `words := strings.Fields(s)` to split `s` into words.
2. Reverse the `words` slice in-place by swapping pairs from the ends toward the center.
3. Return `strings.Join(words, " ")`.

## Complexity
- Time: O(n) — splitting and joining examine each byte/character of the input.
- Space: O(n) — `strings.Fields` and `strings.Join` allocate slices/strings proportional to the input size.

## Alternatives
- Use `strings.Fields` and then construct the result with a `strings.Builder` writing words in reverse order. This is equivalent in complexity and can avoid mutating the `words` slice.
- If preserving exact whitespace is required (not the case for this LeetCode problem), parse manually to keep spaces.

## Implementation notes
- Using `strings.Fields` ensures UTF-8/Unicode whitespace handling via the standard library.
- Reversing in-place and using `strings.Join` is concise and leverages optimized library code for joining strings.

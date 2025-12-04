# Explanation — Reverse Vowels of a String

## Problem
Given a string `s`, reverse only all the vowels in the string and return it. Vowels are `a`, `e`, `i`, `o`, `u` in both lower and upper case. Other characters remain in their original positions.

## Idea
Use a two-pointer approach on a slice of runes:
- Convert the input string to `[]rune` so the algorithm handles multi-byte characters correctly.
- Maintain two indices `i` (left) and `j` (right). Move `i` forward until it points to a vowel, and move `j` backward until it points to a vowel.
- When both `i` and `j` point to vowels and `i < j`, swap the runes at `i` and `j`, then advance `i` and decrement `j`.

This swaps only vowel characters in place (within the rune slice) and preserves order and positions of the non-vowel characters.

## Algorithm (step-by-step)
1. Convert `s` to `runes := []rune(s)`.
2. Set `i := 0`, `j := len(runes)-1`.
3. While `i < j`:
   - Increment `i` until `runes[i]` is a vowel or `i >= j`.
   - Decrement `j` until `runes[j]` is a vowel or `i >= j`.
   - If `i < j`, swap `runes[i]` and `runes[j]`, then `i++`, `j--`.
4. Return `string(runes)`.

## Complexity
- Time: O(n) — each character is examined at most once by the two pointers.
- Space: O(n) additional for the `[]rune` conversion of the input string.

## Alternatives
- Collect all vowels into a separate slice and then write them back in reverse order. This is still O(n) time but uses another O(k) extra memory where `k` is number of vowels.
- If input is ASCII-only and memory matters, operate on `[]byte` and use a small lookup table for vowels.

## Implementation notes
- The implementation in this package uses a `switch`-based `isVowel` function for readability and speed.
- Using `[]rune` ensures correctness for multi-byte characters (UTF-8).

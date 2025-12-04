# Explanation — Valid Palindrome

## Problem
Given a string s, return true if it is a palindrome after removing all non-alphanumeric characters and ignoring case. Alphanumeric characters include letters and numbers.

## Idea
1. Remove characters that are not letters or digits.
2. Convert the resulting string to lowercase so that comparisons are case-insensitive.
3. Use two indices (front and back) and compare characters while moving towards the center. If any pair differs, the string is not a palindrome.

## Algorithm (step-by-step)
1. Clean the string: remove non-alphanumerics using regex or manual check.
2. Lowercase the string.
3. Initialize `left = 0` and `right = len(s) - 1`.
4. While `left < right`:
   - If `s[left] != s[right]`, return `false`.
   - `left++`, `right--`.
5. Return `true`.

## Complexity
- Time: O(n) — single pass (cleaning + comparing).
- Space: O(n) — for storing the cleaned string.

## Alternatives
- **In-place Two Pointers**: Skip non-alphanumeric characters inside the loop without creating a new string. This achieves O(1) space.

## Implementation notes
- Empty string and strings with only non-alphanumeric characters (e.g., "., ") are considered palindromes (returns true).
- This implementation is byte-oriented after cleaning and lowercasing, which is fine for ASCII inputs.

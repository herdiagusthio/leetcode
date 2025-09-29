# Explanation — Valid Palindrome

This document explains the approach used in `solution.go` to determine whether a given string is a palindrome when only alphanumeric characters are considered and case is ignored.

High-level idea
----------------
1. Remove characters that are not letters or digits. We use a precompiled regular expression `[^a-zA-Z0-9]+` to replace non-alphanumeric characters with an empty string.
2. Convert the resulting string to lowercase so that comparisons are case-insensitive.
3. Use two indices (front and back) and compare characters while moving towards the center. If any pair differs, the string is not a palindrome.

Why this works
---------------
- Removing non-alphanumeric characters explicitly matches the problem statement and simplifies later comparisons.
- Lowercasing before comparing avoids having to handle case differences during the two-pointer sweep.
- Two-pointer comparison is efficient (single pass) and obvious to reason about: a palindrome reads the same forwards and backwards, so every corresponding pair must match.

Step-by-step example
--------------------
Input: "A man, a plan, a canal: Panama"

1. Remove non-alphanumerics: "AmanaplanacanalPanama"
2. Lowercase: "amanaplanacanalpanama"
3. Two-pointer checks:
   - Compare index 0 and last -> 'a' == 'a'
   - Compare index 1 and second-last -> 'm' == 'm'
   - ... continue until the center. All comparisons match → palindrome.

Edge cases
----------
- Empty string and strings with only non-alphanumeric characters (e.g., "., ") are considered palindromes (returns true).
- Mixed ASCII and punctuation are handled correctly by the cleaning step.
- This implementation is byte-oriented after cleaning and lowercasing, which is fine for ASCII inputs. For inputs with multi-byte Unicode letters, consider a rune-aware approach.

Complexity recap
-----------------
- Time: O(n)
- Space: O(n)

Follow-ups
----------
- Rune-aware version: if you expect Unicode letters/digits, replace the cleaning and byte comparisons with a rune-based two-pointer scan (or decode runes on the fly using utf8.DecodeRuneInString).
- Normalization: For user-facing text, consider Unicode normalization (NFC/NFD) if combining marks may appear.

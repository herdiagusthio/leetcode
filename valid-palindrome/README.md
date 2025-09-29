# Valid Palindrome

LeetCode: https://leetcode.com/problems/valid-palindrome/

Problem
-------
Given a string s, return true if it is a palindrome after removing all non-alphanumeric characters and ignoring case. Alphanumeric characters include letters and numbers.

Examples
--------
- Input: "A man, a plan, a canal: Panama"
  Output: true

- Input: "race a car"
  Output: false

- Input: " "
  Output: true

Approach (this folder)
----------------------
This implementation follows three simple steps:

1. Remove non-alphanumeric characters using a precompiled regular expression.
2. Convert the cleaned string to lowercase.
3. Use a two-pointer comparison (front and back) to check for palindrome.

The implementation is compact and works efficiently for typical ASCII inputs used in the problem.

Complexity
----------
- Time: O(n) — the input is scanned a small, constant number of times (cleaning, lowercasing, and comparing).
- Space: O(n) — cleaning and lowercasing allocate new strings proportional to the input size.

Notes
-----
- The provided solution is byte-oriented after cleaning and lowercasing and assumes ASCII alphanumerics (matching LeetCode constraints). For full Unicode-aware behavior (non-ASCII letters/digits, combining marks, normalization), use a rune-aware two-pointer approach and optional Unicode normalization.

Constraints
-----------
- 1 <= s.length <= 2 * 10^5
- s consists only of printable ASCII characters (per problem statement).

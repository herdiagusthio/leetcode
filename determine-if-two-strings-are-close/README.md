# Determine if Two Strings Are Close

LeetCode problem: Determine if Two Strings Are Close

Link: https://leetcode.com/problems/determine-if-two-strings-are-close/

Two strings are considered **close** if you can attain one from the other using the following operations:

**Operation 1:** Swap any two existing characters.
- For example, `abcde` → `aecdb`

**Operation 2:** Transform every occurrence of one existing character into another existing character, and do the same with the other character.
- For example, `aacabb` → `bbcbaa` (all `a`'s turn into `b`'s, and all `b`'s turn into `a`'s)

You can use the operations on either string as many times as necessary.

Given two strings, `word1` and `word2`, return `true` if `word1` and `word2` are close, and `false` otherwise.

## Solution
This package implements a solution that checks three essential conditions: (1) both strings have the same length, (2) both strings have the same set of unique characters, and (3) both strings have the same frequency distribution (when sorted). The implementation uses hash maps to count character frequencies and runs in O(n + m log m) time where n is the string length and m is the number of unique characters.

## How to run
From this folder:

```bash
go test -v
```

To run the example `main` in this package:

```bash
go run .
```

## Files
- `solution.go` — implementation and example `main`.
- `solution_test.go` — unit tests.
- `EXPLANATION.md` — step-by-step explanation of the approach.

---

## Examples

**Example 1:**
```
Input: word1 = "abc", word2 = "bca"
Output: true
Explanation: You can attain word2 from word1 in 2 operations.
Apply Operation 1: "abc" → "acb"
Apply Operation 1: "acb" → "bca"
```

**Example 2:**
```
Input: word1 = "a", word2 = "aa"
Output: false
Explanation: It is impossible to attain word2 from word1, or vice versa, in any number of operations.
```

**Example 3:**
```
Input: word1 = "cabbba", word2 = "abbccc"
Output: true
Explanation: You can attain word2 from word1 in 3 operations.
Apply Operation 1: "cabbba" → "caabbb"
Apply Operation 2: "caabbb" → "baaccc"
Apply Operation 2: "baaccc" → "abbccc"
```

## Constraints

- `1 <= word1.length, word2.length <= 10⁵`
- `word1` and `word2` contain only lowercase English letters.

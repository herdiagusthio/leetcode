# Reverse Words in a String

LeetCode problem: Reverse Words in a String

Given an input string `s`, reverse the order of the words. Words are separated by spaces. The returned string should have words in reverse order, separated by a single space, with no leading or trailing spaces.

Link: https://leetcode.com/problems/reverse-words-in-a-string/

## Solution
This package provides a straightforward solution using `strings.Fields` to split the input into words (which trims and collapses whitespace), reversing the words in-place, and joining them with a single space using `strings.Join` (or alternatively constructing with `strings.Builder`).

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

Leetcode Problem: Reverse Words in a String

Given an input string s, reverse the order of the words.

A word is defined as a sequence of non-space characters. The words in s will be separated by at least one space.

Return a string of the words in reverse order concatenated by a single space.

Note that s may contain leading or trailing spaces or multiple spaces between two words. The returned string should only have a single space separating the words. Do not include any extra spaces.

 

Example 1:

Input: s = "the sky is blue"
Output: "blue is sky the"
Example 2:

Input: s = "  hello world  "
Output: "world hello"
Explanation: Your reversed string should not contain leading or trailing spaces.
Example 3:

Input: s = "a good   example"
Output: "example good a"
Explanation: You need to reduce multiple spaces between two words to a single space in the reversed string.
 

Constraints:

1 <= s.length <= 104
s contains English letters (upper-case and lower-case), digits, and spaces ' '.
There is at least one word in s.
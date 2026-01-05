# Decode String

**LeetCode problem:**  
https://leetcode.com/problems/decode-string/

Given an encoded string, return its decoded string.

The encoding rule is: `k[encoded_string]`, where the `encoded_string` inside the square brackets is being repeated exactly `k` times. Note that `k` is guaranteed to be a positive integer.

You may assume that the input string is always valid; there are no extra white spaces, square brackets are well-formed, etc. Furthermore, you may assume that the original data does not contain any digits and that digits are only for those repeat numbers, `k`. For example, there will not be input like `3a` or `2[4]`.

The test cases are generated so that the length of the output will never exceed 10^5.

**Example 1:**
```
Input: s = "3[a]2[bc]"
Output: "aaabcbc"
```

**Example 2:**
```
Input: s = "3[a2[c]]"
Output: "accaccacc"
```

**Example 3:**
```
Input: s = "2[abc]3[cd]ef"
Output: "abcabccdcdcdef"
```

**Constraints:**
- `1 <= s.length <= 30`
- `s` consists of lowercase English letters, digits, and square brackets `[]`.
- `s` is guaranteed to be a valid input.
- All the integers in `s` are in the range `[1, 300]`.

## Solution

The solution uses a stack-based state machine to handle nested encodings. As we traverse the string, we track the current string being built and the current repeat count. When encountering `[`, we save the current state (previous string and repeat count) onto a stack and reset for the new nested level. When encountering `]`, we pop the saved state, repeat the current string the specified number of times, and concatenate it with the previous string. This approach elegantly handles arbitrary nesting depths.

**Time Complexity:** O(maxK * n) where maxK is the maximum value of k and n is the length of the input string  
**Space Complexity:** O(m) where m is the maximum depth of nested encodings

See [EXPLANATION.md](./EXPLANATION.md) for a detailed breakdown of the algorithm.

## How to run

```bash
# Run tests
go test -v

# Run the solution with example inputs
go run .
```

## Files

- `solution.go` - Main solution implementation with stack-based state machine
- `solution_test.go` - Table-driven tests covering examples and edge cases
- `EXPLANATION.md` - Detailed explanation of the algorithm and approach
- `README.md` - This file

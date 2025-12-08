# Maximum Number of Vowels in a Substring of Given Length

**LeetCode problem:**  
https://leetcode.com/problems/maximum-number-of-vowels-in-a-substring-of-given-length/

Given a string `s` and an integer `k`, return the maximum number of vowel letters in any substring of `s` with length `k`.

Vowel letters in English are 'a', 'e', 'i', 'o', and 'u'.

**Example 1:**
```
Input: s = "abciiidef", k = 3
Output: 3
Explanation: The substring "iii" contains 3 vowel letters.
```

**Example 2:**
```
Input: s = "aeiou", k = 2
Output: 2
Explanation: Any substring of length 2 contains 2 vowels.
```

**Example 3:**
```
Input: s = "leetcode", k = 3
Output: 2
Explanation: "lee", "eet" and "ode" contain 2 vowels.
```

**Constraints:**
- `1 <= s.length <= 10^5`
- `s` consists of lowercase English letters
- `1 <= k <= s.length`

## Solution

The solution uses a sliding window approach to efficiently count vowels in each k-length substring. It maintains a running count of vowels as the window moves through the string, updating the count when characters enter or leave the window. This avoids recounting vowels for overlapping substrings.

**Time Complexity:** O(n) where n is the length of string s  
**Space Complexity:** O(1)

See [EXPLANATION.md](./EXPLANATION.md) for a detailed breakdown of the algorithm.

## How to run

```bash
# Run tests
go test -v

# Run the solution with example inputs
go run .
```

## Files

- `solution.go` - Main solution implementation with sliding window approach
- `solution_test.go` - Table-driven tests covering examples and edge cases
- `EXPLANATION.md` - Detailed explanation of the algorithm and approach
- `README.md` - This file
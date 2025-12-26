# Count Vowel Substrings of a String

**LeetCode problem:**  
https://leetcode.com/problems/count-vowel-substrings-of-a-string/

A substring is a contiguous (non-empty) sequence of characters within a string.

A vowel substring is a substring that only consists of vowels ('a', 'e', 'i', 'o', and 'u') and has all five vowels present in it.

Given a string `word`, return the number of vowel substrings in `word`.

**Example 1:**
- Input: word = "aeiouu"
- Output: 2
- Explanation:
    The vowel substrings are "aeiou" and "aeiouu".

**Example 2:**
- Input: word = "unicornarihan"
- Output: 0
- Explanation:
    Not all 5 vowels are present, so there are no vowel substrings.

**Example 3:**
- Input: word = "cuaieuouac"
- Output: 7
- Explanation:
    The vowel substrings are (shown with underlining in the original word):
    1. "c<u>uaieuo</u>uac" → "uaieuo"
    2. "c<u>uaieuou</u>ac" → "uaieuou"
    3. "c<u>uaieuoua</u>c" → "uaieuoua"
    4. "cu<u>aieuo</u>uac" → "aieuo"
    5. "cu<u>aieuou</u>ac" → "aieuou"
    6. "cu<u>aieuoua</u>c" → "aieuoua"
    7. "cua<u>ieuoua</u>c" → "ieuoua"
    
    Each of these 7 substrings contains all 5 vowels (a, e, i, o, u).

**Constraints:**
- `1 <= word.length <= 100`
- `word` consists of lowercase English letters only

## Solution

The solution uses a sliding window approach with position tracking. It maintains the last consonant position and a frequency map of vowels in the current window. When all 5 vowels are present, it shrinks the window from the left while maintaining all vowels, then counts valid substrings by calculating the number of possible starting positions (`left - lastConsonant`).

**Time Complexity:** O(n) where n is the length of the word  
**Space Complexity:** O(1) - map stores at most 5 vowels

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

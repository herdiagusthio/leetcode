# Maximum Number of Vowels in a Substring of Given Length - Explanation

## Problem Summary
Find the maximum number of vowel letters in any substring of string `s` with length `k`.

## Approach
This solution uses a **sliding window** technique to efficiently count vowels in each k-length substring:

1. **Initial Window**: Count the vowels in the first k characters
2. **Slide the Window**: Move the window one position to the right
   - Remove the leftmost character (if it was a vowel, decrement count)
   - Add the new rightmost character (if it's a vowel, increment count)
3. **Track Maximum**: Keep track of the highest vowel count seen

## Why This Works
The sliding window approach works because:
- We only need to consider substrings of exactly length k
- When moving from one k-length window to the next, only two characters change:
  - One character leaves the window (left side)
  - One character enters the window (right side)
- Instead of recounting all k characters each time, we update the count based on these two characters

## Time Complexity
**O(n)** where n is the length of string s
- We iterate through the string once
- Each character is examined at most twice (once when entering, once when leaving the window)
- The vowel check is O(1) using a simple switch statement

## Space Complexity
**O(1)** - We only use a constant amount of extra space:
- A few integer variables for counting
- No data structures that grow with input size

## Edge Cases
1. **No vowels**: Returns 0 (e.g., "rhythms")
2. **All vowels**: Returns k (e.g., "aeiou" with k=5)
3. **k equals string length**: Checks only one window (the entire string)
4. **k equals 1**: Finds any single vowel in the string

## Worked Example
For `s = "abciiidef"`, `k = 3`:

```
Initial window: "abc"
Vowels: a (count = 1)
Max = 1

Window: "bci"
Remove 'a' (vowel): count = 0
Add 'i' (vowel): count = 1
Max = 1

Window: "cii"
Remove 'b': count = 1
Add 'i' (vowel): count = 2
Max = 2

Window: "iii"
Remove 'c': count = 2
Add 'i' (vowel): count = 3
Max = 3

Window: "iid"
Remove 'i' (vowel): count = 2
Add 'd': count = 2
Max = 3

Window: "ide"
Remove 'i' (vowel): count = 1
Add 'e' (vowel): count = 2
Max = 3

Window: "def"
Remove 'd': count = 2
Add 'f': count = 2
Max = 3
```

Result: **3**

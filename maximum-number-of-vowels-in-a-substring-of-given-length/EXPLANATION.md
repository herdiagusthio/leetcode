# Maximum Number of Vowels in a Substring of Given Length - Explanation

## Problem Summary
Given a string `s` and an integer `k`, find the maximum number of vowel letters in any substring of `s` with length `k`. Vowels are: a, e, i, o, u (lowercase only).

## Approach
This solution uses a **sliding window technique**:

1. **Count Initial Window**: Count vowels in the first `k` characters
2. **Set Initial Maximum**: Record this count as the current maximum
3. **Slide the Window**: For each position from `k` to `n-1`:
   - Remove the character leaving the window (left side)
   - Add the character entering the window (right side)
   - Update vowel count accordingly
4. **Track Maximum**: Update maximum if current count is higher
5. **Return Result**: The maximum vowel count found

## Why This Works
The key insight is incremental counting with a fixed-size window:
- We need exactly k-length substrings, not variable length
- When sliding the window one position right, only 2 characters change
- Instead of recounting all k characters (O(k) per window)
- We update based on the one character leaving and one entering (O(1) per window)

**Why this finds the maximum:**
- We examine every possible k-length substring exactly once
- We track the maximum count across all windows
- No valid substring is skipped

## Time Complexity
**O(n)** where n is the length of string s
- Count vowels in first k characters: O(k)
- Slide window (n-k) times with O(1) operations each: O(n-k)
- Vowel check uses simple comparison: O(1)
- Total: O(k) + O(n-k) = O(n)

## Space Complexity
**O(1)** - constant space
- Only a few integer variables for counting and tracking
- No data structures that grow with input size
- Vowel checking done with direct character comparison

## Edge Cases
1. **No vowels**: Returns 0 (e.g., "bcdfg" with any k)
2. **All vowels**: Returns k (e.g., "aeiou" with k=3 returns 3)
3. **k = 1**: Returns 1 if any vowel exists, 0 otherwise
4. **k = n** (full string): Returns total vowel count in string
5. **Single character string**: Returns 1 if vowel, 0 otherwise (when k=1)

## Worked Example
For `s = "abciiidef"`, `k = 3`:

```
Step 1: Count first window "abc"
  Vowels: 'a' → count = 1
  maxVowels = 1

Step 2: Slide to "bci"
  Remove 'a' (vowel): count = 0
  Add 'i' (vowel): count = 1
  maxVowels = 1

Step 3: Slide to "cii"
  Remove 'b' (not vowel): count = 1
  Add 'i' (vowel): count = 2
  maxVowels = 2

Step 4: Slide to "iii"
  Remove 'c' (not vowel): count = 2
  Add 'i' (vowel): count = 3
  maxVowels = 3 ✓ (new maximum!)

Step 5: Slide to "iid"
  Remove 'i' (vowel): count = 2
  Add 'd' (not vowel): count = 2
  maxVowels = 3

Step 6: Slide to "ide"
  Remove 'i' (vowel): count = 1
  Add 'e' (vowel): count = 2
  maxVowels = 3

Step 7: Slide to "def"
  Remove 'd' (not vowel): count = 2
  Add 'f' (not vowel): count = 2
  maxVowels = 3

Final result: 3 ✓
Best substring: "iii" (all vowels)
```

## Visual Representation

For `s = "abciiidef"`, `k = 3`:

```
Index:  0 1 2 3 4 5 6 7 8
Char:   a b c i i i d e f
        v   v v v     v       (v = vowel)

Window 1: [a b c]             count = 1
          └─────┘

Window 2:   [b c i]           count = 1
            └─────┘

Window 3:     [c i i]         count = 2
              └─────┘

Window 4:       [i i i]       count = 3 ← maximum!
                └─────┘

Window 5:         [i i d]     count = 2
                  └─────┘

Window 6:           [i d e]   count = 2
                    └─────┘

Window 7:             [d e f] count = 1
                      └─────┘

Best window: "iii" with 3 vowels
```

## Sliding Mechanism

```
s = "abciiidef", k = 3

Step 1: abc → count 'a' = 1 vowel
        ╚═╝

Step 2: slide right (remove 'a', add 'i')
         bci
         ╚═╝
        remove: a (vowel) → count = 0
        add:    i (vowel) → count = 1

Step 3: slide right (remove 'b', add 'i')
          cii
          ╚═╝
        remove: b (not vowel) → count = 1
        add:    i (vowel) → count = 2

Step 4: slide right (remove 'c', add 'i')
           iii
           ╚═╝
        remove: c (not vowel) → count = 2
        add:    i (vowel) → count = 3 ✓ MAX!

Continue sliding... max remains 3
```

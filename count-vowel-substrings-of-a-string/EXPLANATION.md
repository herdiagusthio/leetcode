# Count Vowel Substrings of a String - Explanation

## Problem Summary
Given a string `word`, return the number of vowel substrings. A vowel substring is a contiguous sequence of characters that only consists of vowels ('a', 'e', 'i', 'o', 'u') and contains all five vowels at least once.

## Approach
This solution uses a **sliding window technique** with position tracking:

1. **Track Last Consonant**: Keep track of the last consonant position to know where valid vowel sequences can start
2. **Expand Window**: Move right pointer, adding vowels to frequency map
3. **Shrink Window**: When all 5 vowels are present, shrink from left while maintaining all 5
4. **Count Substrings**: For each valid window, count all possible starting positions from `lastConsonant + 1` to `left`
5. **Reset on Consonant**: When encountering a consonant, reset the window

## Why This Works
The key insight is understanding how many substrings end at the current position:
- When we have a window `[left, right]` with all 5 vowels
- Any substring starting from `lastConsonant + 1` to `left` and ending at `right` is valid
- This gives us `left - lastConsonant` valid substrings ending at position `right`
- We shrink `left` as much as possible while maintaining all 5 vowels to maximize count

**Why shrink the window?**
- A smaller valid window means more possible starting positions
- Example: If "aeiou" works, we can start from positions that include just "aeiou", or "xaeiou", etc.
- By minimizing the window, we count all valid longer substrings

## Time Complexity
**O(n)** where n is the length of the word
- Right pointer traverses the string once: O(n)
- Left pointer only moves forward, never backwards: O(n) total movement
- Each position visited at most twice (once by right, once by left)
- Map operations are O(1) since map size is bounded by 5 vowels

## Space Complexity
**O(1)** - constant space
- Vowel frequency map stores at most 5 entries (for 5 vowels)
- A few integer variables for pointers
- No data structures that grow with input size

## Edge Cases
1. **Single vowel**: Returns 0 (need all 5 vowels)
2. **No vowels**: Returns 0
3. **All 5 vowels once**: Returns 1
4. **Repeated same vowel**: Returns 0 (missing other vowels)
5. **Consonants in middle**: Splits into separate vowel segments
6. **Short string (< 5)**: Returns 0 (impossible to have all 5 vowels)

## Worked Example
For `word = "aeiouu"`:

```
Initial state: lastConsonant = -1, left = 0, result = 0

Process each character:
- Indices 0-3 ('a','e','i','o'): Add to vowFreq, but only 4 unique vowels
  
- Index 4 ('u'): 
  vowFreq = {a:1, e:1, i:1, o:1, u:1} → All 5 vowels present! ✓
  Window: "aeiou" (left=0, right=4)
  Cannot shrink (all frequencies = 1)
  result += 0 - (-1) = 1
  Valid substrings ending at index 4: ["aeiou"]

- Index 5 ('u'):
  vowFreq = {a:1, e:1, i:1, o:1, u:2} → All 5 vowels present! ✓
  Window: "aeiouu" (left=0, right=5)
  Cannot shrink (removing 'a' would eliminate it)
  result += 0 - (-1) = 1
  Total result = 2
  Valid substrings ending at index 5: ["aeiouu"]
  
Final result: 2 ✓
```

## Complex Example
For `word = "cuaieuouac"`:

```
Index:  0 1 2 3 4 5 6 7 8 9
Char:   c u a i e u o u a c
        ^                 ^ consonants (reset positions)

Process:
- Index 0 ('c'): Consonant, lastConsonant = 0, reset window

- Index 6 ('o'):
  vowFreq = {u:2, a:1, i:1, e:1, o:1} → All 5 vowels! ✓
  Window: "uaieuo" (left=1, right=6)
  Shrink: vowFreq['u']=2, can remove one 'u'
    left → 2, vowFreq = {u:1, a:1, i:1, e:1, o:1}
    vowFreq['a']=1, cannot shrink further
  result += 2 - 0 = 2
  Valid substrings: ["uaieuo", "aieuo"]

- Index 7 ('u'):
  vowFreq = {u:2, a:1, i:1, e:1, o:1} → All 5 vowels! ✓
  Shrink to left=2
  result += 2 - 0 = 2
  Valid substrings: ["uaieuou", "aieuou"]
  Total: 4

- Index 8 ('a'):
  vowFreq = {u:2, a:2, i:1, e:1, o:1} → All 5 vowels! ✓
  Shrink to left=2
  result += 2 - 0 = 2
  Valid substrings: ["uaieuoua", "aieuoua"]
  Total: 6

- Index 9 ('c'): Consonant, continue but check if window still valid
  Before consonant, we had one more valid position at left=4
  result += 1
  Total: 7

Final result: 7 ✓
```

## Visual Representation

For `word = "aeiouu"`:
```
Index:  0 1 2 3 4 5
Char:   a e i o u u

Substring 1: "aeiou" (indices 0-4)
              └───┘

Substring 2: "aeiouu" (indices 0-5)
              └────┘
```

For `word = "cuaieuouac"`:
```
Index:  0 1 2 3 4 5 6 7 8 9
Char:   c u a i e u o u a c
        ^                 ^ consonants split segments

Valid vowel segment: "uaieuoua" (indices 1-8)
                      └──────┘

The 7 valid substrings (all containing all 5 vowels):
1. c[uaieuo]uac   → "uaieuo"   (indices 1-6)
2. cu[aieuo]uac   → "aieuo"    (indices 2-6)
3. c[uaieuou]ac   → "uaieuou"  (indices 1-7)
4. cu[aieuou]ac   → "aieuou"   (indices 2-7)
5. c[uaieuoua]c   → "uaieuoua" (indices 1-8)
6. cu[aieuoua]c   → "aieuoua"  (indices 2-8)
7. cua[ieuoua]c   → "ieuoua"   (indices 4-8)
```

# Valid Palindrome - Explanation

## Problem Summary
Given a string `s`, return true if it is a palindrome after converting all uppercase letters to lowercase and removing all non-alphanumeric characters. A palindrome reads the same forward and backward.

## Approach
This solution uses a **two-pointer technique with in-place character validation**:

1. **Initialize Pointers**: Start with `left = 0` and `right = len(s) - 1`
2. **Skip Invalid Characters**: Move pointers to skip non-alphanumeric characters
   - Increment `left` while pointing to non-alphanumeric
   - Decrement `right` while pointing to non-alphanumeric
3. **Compare Characters**: Once both pointers are on valid characters:
   - Compare characters (case-insensitive)
   - If they don't match, return false
   - Move both pointers inward
4. **Continue**: Repeat until pointers meet or cross
5. **Return True**: If all comparisons matched

## Why This Works
The key insight is using two pointers to avoid creating a new cleaned string:
- We skip invalid characters on-the-fly instead of preprocessing
- Case-insensitive comparison handles uppercase/lowercase
- Two pointers converge from both ends, comparing mirror positions
- If any mismatch occurs, it's not a palindrome
- O(1) space by not allocating a new string

**Why in-place is better:**
- Avoids O(n) space for cleaned string
- Single pass through the string
- Early exit on first mismatch

## Time Complexity
**O(n)** where n is the length of string s
- Each character examined at most once
- Left pointer moves right at most n times
- Right pointer moves left at most n times
- Character comparison is O(1)

## Space Complexity
**O(1)** - constant space
- Only two integer pointers
- No additional data structures
- Case conversion done on-the-fly

## Edge Cases
1. **Empty string**: Returns true (empty is considered palindrome)
2. **Single character**: Returns true (always palindromic)
3. **Only non-alphanumeric**: Returns true (e.g., ".,!")
4. **All lowercase**: Direct comparison
5. **Mixed case**: Case-insensitive comparison handles it
6. **With spaces**: Spaces are skipped as non-alphanumeric

## Worked Example
For `s = "A man, a plan, a canal: Panama"`:

```
Clean view: "amanaplanacanalpanama"

Initial: left=0, right=30

Step 1: Skip to first alphanumeric
  left=0: 'A' (alphanumeric) ✓
  right=30: 'a' (alphanumeric) ✓
  Compare: 'A' (lowercase 'a') == 'a' ✓
  left++, right--

Step 2: left=1, right=29
  left=1: ' ' (skip)
  left=2: 'm' ✓
  right=29: 'm' ✓
  Compare: 'm' == 'm' ✓
  left++, right--

...continue comparing...

Final: All characters match
Result: true ✓
```

## Visual Representation

For `s = "A man, a plan, a canal: Panama"`:

```
Original:  A   m a n ,   a   p l a n ,   a   c a n a l :   P a n a m a
           ↑                                                           ↑
         left                                                        right

After cleaning (conceptual):
           a m a n a p l a n a c a n a l p a n a m a
           ↑                                       ↑
         left                                   right

Comparisons:
a == a ✓
 m == m ✓
  a == a ✓
   n == n ✓
    a == a ✓
     p == p ✓
      l == l ✓
       a == a ✓
        n == n ✓
         a == a ✓
          c (middle)
          
All match → palindrome!
```

## Example: Not a Palindrome
For `s = "race a car"`:

```
Clean view: "raceacar"

left=0 ('r'), right=7 ('r'): r == r ✓
left=1 ('a'), right=6 ('a'): a == a ✓  
left=2 ('c'), right=5 ('c'): c == c ✓
left=3 ('e'), right=4 ('a'): e != a ✗

Result: false
```

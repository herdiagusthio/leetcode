# Is Subsequence - Explanation

## Problem Summary
Given two strings `s` and `t`, return true if `s` is a subsequence of `t`. A subsequence is formed by deleting some (or no) characters from `t` without changing the order of remaining characters.

## Approach
This solution uses a **two-pointer greedy matching** approach:

1. **Initialize Pointer**: Start with pointer at index 0 of `s`
2. **Iterate Through `t`**: For each character in `t`:
   - If it matches current character in `s`, advance `s` pointer
   - If `s` pointer reaches end of `s`, all characters matched
3. **Check Completion**: Return true if all characters of `s` were matched

## Why This Works
The key insight is that greedy matching is optimal:
- **Order must be preserved**: We can only match characters in the order they appear
- **Greedy is optimal**: If we can match a character now, we should
  - Skipping a match never helps - it only reduces future opportunities
  - Earlier matches give us more chances to match remaining characters
- **Single pass suffices**: We don't need to backtrack
- **No need to remember positions**: Just track progress in `s`

**Why we don't need to try different matchings**:
- If `s[i]` matches `t[j]`, using this match is always optimal or equal to any alternative
- Delaying the match (using a later occurrence) never provides an advantage

## Time Complexity
**O(n)** where n is the length of string `t`
- Single pass through `t`
- Each character in `t` is examined exactly once
- Pointer operations are O(1)
- String indexing is O(1)

Note: Independent of `s` length since we only advance through `s` when matching

## Space Complexity
**O(1)** - constant space
- Only one integer pointer variable
- No data structures that grow with input
- Input strings are not copied or modified

## Edge Cases
1. **Empty `s`**: Empty string is subsequence of any string, return true
2. **Empty `t`**: Only true if `s` is also empty
3. **`s` longer than `t`**: Cannot be subsequence if `s` has more characters
4. **Identical strings**: `s` is subsequence of itself
5. **All characters match**: `s = "abc"`, `t = "aabbcc"` → true
6. **No matches**: `s = "abc"`, `t = "def"` → false
7. **Single character**: Works correctly

## Worked Example
For `s = "abc"`, `t = "ahbgdc"`:

```
s = "abc"
t = "ahbgdc"
     ^

Initial: sPointer = 0 (looking for 'a')

Index 0: t[0] = 'a'
'a' == s[0] ('a') ✓
sPointer = 1 (now looking for 'b')

Index 1: t[1] = 'h'
'h' != s[1] ('b')
sPointer = 1 (still looking for 'b')

Index 2: t[2] = 'b'
'b' == s[1] ('b') ✓
sPointer = 2 (now looking for 'c')

Index 3: t[3] = 'g'
'g' != s[2] ('c')
sPointer = 2 (still looking for 'c')

Index 4: t[4] = 'd'
'd' != s[2] ('c')
sPointer = 2 (still looking for 'c')

Index 5: t[5] = 'c'
'c' == s[2] ('c') ✓
sPointer = 3

sPointer (3) == len(s) (3) ✓
All characters matched!

Result: true
```

## Visual Representation
```
s = "abc"
     012

t = "ahbgdc"
     012345

Matching process:
t[0]='a' → matches s[0]='a' ✓
t[1]='h' → skip
t[2]='b' → matches s[1]='b' ✓
t[3]='g' → skip
t[4]='d' → skip
t[5]='c' → matches s[2]='c' ✓

Visual matching:
a h b g d c
↓   ↓     ↓
a   b     c  ← All characters of s found in order
```

## Example: Not a Subsequence
For `s = "axc"`, `t = "ahbgdc"`:

```
s = "axc"
t = "ahbgdc"

Index 0: t[0] = 'a' → matches s[0] = 'a' ✓, sPointer = 1
Index 1: t[1] = 'h' → doesn't match s[1] = 'x'
Index 2: t[2] = 'b' → doesn't match s[1] = 'x'
Index 3: t[3] = 'g' → doesn't match s[1] = 'x'
Index 4: t[4] = 'd' → doesn't match s[1] = 'x'
Index 5: t[5] = 'c' → doesn't match s[1] = 'x'

End of t reached, sPointer = 1 (not 3)
Result: false (couldn't find 'x')
```

## Example: Empty String
For `s = ""`, `t = "abc"`:

```
len(s) = 0
Return true immediately
(Empty string is subsequence of any string)
```

## Pattern Visualization
```
Valid subsequence (greedy matching):
t:  a → h → b → g → d → c
s:  a ─────→ b ─────────→ c
    └─ match ─┘           └─ match

Invalid (missing character):
t:  a → h → b → g → d → c
s:  a ─────→ x ← not found!
    └─ match
```

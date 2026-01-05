# Decode String - Explanation

## Problem Summary
Given an encoded string with the pattern `k[encoded_string]`, where `k` is a positive integer and `encoded_string` is repeated exactly `k` times, return the decoded string. The encoding can be nested, meaning `encoded_string` can itself contain encoded patterns. The input is always valid with well-formed brackets, no extra spaces, and digits only appear as repeat counts.

## Approach
This solution uses a **stack-based state machine** to handle nested encodings:

1. **Initialize State**: Track current string being built and current repeat count
2. **Process Each Character**: Iterate through the input string character by character
3. **State Machine Logic**:
   - **Digit**: Build the repeat count (handle multi-digit numbers)
   - **Letter**: Append to current string using `strings.Builder`
   - **Opening bracket `[`**: Save current state to stack, reset for new nested level
   - **Closing bracket `]`**: Pop saved state, repeat current string, concatenate with previous
4. **Return Result**: The final current string contains the fully decoded result

## Why This Works
The key insight is that encoded strings can be nested, requiring us to track multiple levels of context:
- When we encounter `[`, we're entering a new nested level
- We need to remember what came before this nested level
- When we encounter `]`, we complete the current level and merge back with the previous level

**Why stack is the right data structure:**
- Stack naturally handles nested structures (LIFO - Last In, First Out)
- Each `[` pushes a new context onto the stack
- Each `]` pops the most recent context
- The stack depth equals the nesting depth at any point
- This matches the structure of nested encodings perfectly

**Why strings.Builder is important:**
- String concatenation in Go creates new strings (O(n) per operation)
- For k repetitions of a string of length n, naive concatenation is O(k * n²)
- `strings.Builder` uses an internal byte buffer that grows as needed
- Appending to `strings.Builder` is amortized O(1)
- This reduces the repetition operation from O(k * n²) to O(k * n)

**Why the algorithm is correct:**
- Each character is processed exactly once in left-to-right order
- State is saved before entering nested levels and restored when exiting
- The repeat count is correctly parsed for multi-digit numbers (e.g., "100[a]")
- String repetition happens at the right time (when closing bracket is encountered)
- The final result accumulates all decoded strings in the correct order

## Time Complexity
**O(maxK * n)** where maxK is the maximum value of k and n is the length of the input string
- We traverse the input string once: O(n)
- For each encoded segment, we may repeat the string up to k times
- In worst case with deep nesting like `2[2[2[a]]]`, we get exponential expansion
- More precisely: O(sum of all characters in the decoded output)
- For input "3[a]2[bc]", we process 7 input chars and generate 6 output chars

## Space Complexity
**O(m + k * n)** where m is the maximum depth of nested encodings
- Stack stores at most m states (one per nesting level): O(m)
- Each state contains a string and a count
- `strings.Builder` holds the current string being built: O(k * n) in worst case
- Total space depends on both nesting depth and string lengths at each level

## Edge Cases
1. **No encoding**: Plain string (e.g., `"abc"` → `"abc"`)
2. **Single level encoding**: No nesting (e.g., `"3[a]"` → `"aaa"`)
3. **Multiple separate encodings**: No nesting (e.g., `"3[a]2[bc]"` → `"aaabcbc"`)
4. **Nested encoding**: Multiple levels (e.g., `"3[a2[c]]"` → `"accaccacc"`)
5. **Deep nesting**: Many levels (e.g., `"2[2[2[a]]]"` → `"aaaaaaaa"`)
6. **Multi-digit repeat count**: Large numbers (e.g., `"100[a]"` → "a" repeated 100 times)
7. **Mixed content**: Encoding with regular characters (e.g., `"ab2[c]d"` → `"abccd"`)
8. **Leading/trailing characters**: Non-encoded text around encodings
9. **Complex patterns**: Multiple nested levels with regular characters

## Worked Example
For `s = "3[a2[c]]"`:

```
Initial: currentString = "", currentRepeat = 0, stack = []

Step 1: c = '3' (digit)
  currentRepeat = 0 * 10 + 3 = 3

Step 2: c = '[' (opening bracket)
  Save state: push {prevString: "", repeat: 3} to stack
  Reset: currentString = "", currentRepeat = 0
  stack = [{prevString:"", repeat:3}]

Step 3: c = 'a' (letter)
  currentString.WriteString("a")
  currentString = "a"

Step 4: c = '2' (digit)
  currentRepeat = 0 * 10 + 2 = 2

Step 5: c = '[' (opening bracket)
  Save state: push {prevString: "a", repeat: 2} to stack
  Reset: currentString = "", currentRepeat = 0
  stack = [{prevString:"", repeat:3}, {prevString:"a", repeat:2}]

Step 6: c = 'c' (letter)
  currentString.WriteString("c")
  currentString = "c"

Step 7: c = ']' (closing bracket)
  Pop state: {prevString: "a", repeat: 2}
  Repeat: "c" * 2 = "cc"
  Concatenate: "a" + "cc" = "acc"
  currentString = "acc"
  stack = [{prevString:"", repeat:3}]

Step 8: c = ']' (closing bracket)
  Pop state: {prevString: "", repeat: 3}
  Repeat: "acc" * 3 = "accaccacc"
  Concatenate: "" + "accaccacc" = "accaccacc"
  currentString = "accaccacc"
  stack = []

Final result: "accaccacc" ✓
```

## Complex Example
For `s = "2[abc]3[cd]ef"`:

```
Initial: currentString = "", currentRepeat = 0, stack = []

Step 1: c = '2'
  currentRepeat = 2

Step 2: c = '['
  Push: {prevString:"", repeat:2}
  Reset: currentString = "", currentRepeat = 0

Step 3-5: c = 'a', 'b', 'c'
  currentString = "abc"

Step 6: c = ']'
  Pop: {prevString:"", repeat:2}
  Repeat: "abc" * 2 = "abcabc"
  currentString = "abcabc"

Step 7: c = '3'
  currentRepeat = 3

Step 8: c = '['
  Push: {prevString:"abcabc", repeat:3}
  Reset: currentString = "", currentRepeat = 0

Step 9-10: c = 'c', 'd'
  currentString = "cd"

Step 11: c = ']'
  Pop: {prevString:"abcabc", repeat:3}
  Repeat: "cd" * 3 = "cdcdcd"
  Concatenate: "abcabc" + "cdcdcd" = "abcabccdcdcd"
  currentString = "abcabccdcdcd"

Step 12-13: c = 'e', 'f'
  currentString = "abcabccdcdcdef"

Final result: "abcabccdcdcdef" ✓
```

## Visual Representation

For `s = "3[a2[c]]"`:

```
Parse tree structure:
    3[...]
      ├── a
      └── 2[...]
            └── c

Execution trace:
Level 0: ""
    ↓ see '3['
Level 1: ""
    ↓ see 'a'
Level 1: "a"
    ↓ see '2['
Level 2: ""
    ↓ see 'c'
Level 2: "c"
    ↓ see ']' → repeat "c" × 2
Level 1: "a" + "cc" = "acc"
    ↓ see ']' → repeat "acc" × 3
Level 0: "" + "accaccacc" = "accaccacc"

Result: "accaccacc"
```

For `s = "2[abc]3[cd]ef"`:

```
Timeline:
Input:  2 [ a b c ] 3 [ c d ] e f
        ↓
State:  ""
        ↓ 2[
        "" (saved: "", repeat=2)
        ↓ abc
        "abc"
        ↓ ]
        "abcabc" (from "" + "abc"*2)
        ↓ 3[
        "" (saved: "abcabc", repeat=3)
        ↓ cd
        "cd"
        ↓ ]
        "abcabccdcdcd" (from "abcabc" + "cd"*3)
        ↓ ef
        "abcabccdcdcdef"

Result: "abcabccdcdcdef"
```

## Algorithm Flow Diagram

```
For each character c:
    ┌──────────────────────────────┐
    │ What type is c?              │
    └───┬──────┬──────┬────────┬───┘
        │      │      │        │
     Digit  Letter  '['      ']'
        │      │      │        │
        v      v      v        v
    ┌────┐ ┌─────┐ ┌────┐  ┌─────┐
    │Update│Append│Save │  │Pop  │
    │repeat│ to   │state│  │state│
    │count │string│     │  │     │
    └────┘ └─────┘ └────┘  └─────┘
                     │        │
                     v        v
                  ┌──────┐ ┌──────┐
                  │Reset │ │Repeat│
                  │for   │ │string│
                  │nested│ │     │
                  └──────┘ └──────┘
                              │
                              v
                          ┌──────┐
                          │Concat│
                          │with  │
                          │prev  │
                          └──────┘

Stack behavior:
  '[' → PUSH {prevString, repeat}
  ']' → POP → repeat → concatenate
```

## Key Implementation Details

**Multi-digit number parsing:**
```go
currentRepeat = currentRepeat*10 + int(c-'0')
```
- Converts character '0'-'9' to integer 0-9
- Builds number left-to-right (e.g., '1', '2', '3' → 123)

**State structure:**
```go
type state struct {
    prevString string  // String built before this bracket
    repeat     int     // How many times to repeat
}
```

**String repetition with strings.Builder:**
```go
var temp strings.Builder
for i := 0; i < repeat; i++ {
    temp.WriteString(currentString.String())
}
```
- Efficient O(k*n) instead of O(k²*n) with string concatenation

**State restoration:**
```go
currentString.Reset()
currentString.WriteString(s.prevString)
currentString.WriteString(temp.String())
```
- Reset builder to clear previous content
- Write previous string first (maintains order)
- Write repeated string second

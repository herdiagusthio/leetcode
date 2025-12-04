# Explanation — Valid Parentheses

## Problem
Given a string consisting of parentheses characters '()[]{}', determine whether it is valid. A valid string has matching bracket pairs in the correct order.

## Idea
Use a stack to track open brackets. When a closing bracket appears, ensure the top of the stack is the matching opening bracket and pop it. If it doesn't match or the stack is empty, the string is invalid. After processing all characters, the string is valid only if the stack is empty.

## Algorithm (step-by-step)
1. Initialize an empty stack.
2. Iterate through the input string:
   - If character is `(`, `[`, or `{`, push onto stack.
   - If character is `)`, `]`, or `}`, check stack:
     - If stack is empty, return `false`.
     - Pop top element. If it doesn't match the current closing bracket, return `false`.
3. After loop, return `true` if stack is empty, else `false`.

## Complexity
- Time: O(n) — single pass.
- Space: O(n) — in the worst case when all characters are opens.

## Implementation notes
- Odd-length strings are invalid — we early-return for those.
- Empty string is valid.
- The implementation uses a byte stack with a closing->opening map for efficient checks.

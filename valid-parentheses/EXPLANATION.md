Valid Parentheses — Explanation
================================

Problem
-------

Given a string consisting of parentheses characters '()[]{}', determine whether it is valid. A valid string has matching bracket pairs in the correct order.

High-level idea
---------------

Use a stack to track open brackets. When a closing bracket appears, ensure the top of the stack is the matching opening bracket and pop it. If it doesn't match or the stack is empty, the string is invalid. After processing all characters, the string is valid only if the stack is empty.

Approach
--------

1) Stack (recommended)

- Iterate through the input string.
- Push opening brackets onto the stack.
- For closing brackets, check the stack top and pop. If the stack is empty or the types mismatch, return false.

Complexity
----------

- Time: O(n) — single pass.
- Space: O(n) — in the worst case when all characters are opens.

Edge cases
----------

- Odd-length strings are invalid — we early-return for those.
- Empty string is valid.
- Strings with unexpected characters can be treated as invalid (problem constraints don't require this).

Notes
-----

- The implementation uses a byte stack with a closing->opening map for efficient checks.
- This approach is standard and performs well for the problem constraints.

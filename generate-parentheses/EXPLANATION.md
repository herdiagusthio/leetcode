# Explanation — Generate Parentheses

## Problem
Given n pairs of parentheses, generate all combinations of well-formed parentheses.

## Idea
Use DFS/backtracking while tracking two counters: number of open parentheses used and number of close parentheses used. At each step:
- add '(' when open < n
- add ')' when close < open

When the current buffer reaches length 2*n, append it to the results.

## Algorithm (step-by-step)
1. Initialize a results list.
2. Call the recursive backtracking function with empty string (or buffer), `open = 0`, `close = 0`.
3. In the recursive function:
   - Base case: If `len(current) == 2 * n`, add `current` to results and return.
   - If `open < n`:
     - Append `(` to `current`.
     - Recurse with `open + 1`.
     - Backtrack (remove `(`).
   - If `close < open`:
     - Append `)` to `current`.
     - Recurse with `close + 1`.
     - Backtrack (remove `)`).
4. Return results.

## Complexity
- Time: O(C_n * n), where C_n is the nth Catalan number.
- Space: O(C_n * n) for output + O(n) recursion depth.

## Alternatives
- In-place buffer: Preallocate a byte slice of size `2*n` to reduce allocations during recursion.
- Dynamic Programming: Constructive approach using Catalan decomposition.

## Implementation notes
- For clarity use the recursive append-based approach.
- Remember converting a `[]byte` to `string` allocates a new string.

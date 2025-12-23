# Explanation — Removing Stars From a String

## Problem
Given a string `s` containing stars `*`, we need to remove stars. For every star, we remove the closest non-star character to its left and the star itself. We verify that the operation is always possible.

## Approach 1: Stack
This problem naturally fits a **Stack** data structure. As we process characters from left to right:
- A non-star character is "pushed" onto the stack.
- A star `*` means we "pop" the most recently added character (the closest one to the left).

Since we are processing a string and Go strings are immutable, we can use a `[]byte` slice as our stack which is efficient and mutable.

### Algorithm (step-by-step)
1. Initialize an empty byte slice `stack` with capacity equal to the input string length (optimization to avoid resizing).
2. Iterate through each character of the input string `s`.
3. If the current character is `*`:
   - Remove the last element from `stack` (pop).
4. If the current character is NOT `*`:
   - Append it to `stack` (push).
5. Convert `stack` to a string and return it.

## Approach 2: Backwards Iteration
Alternatively, we can iterate through the string **backwards**. This allows us to know immediately if a character should be skipped (removed) because we would have already encountered the `*` that removes it.

### Algorithm (step-by-step)
1. Initialize a `starCounter` to 0 and a byte slice `result` to store kept characters.
2. Iterate through the string `s` from the last index down to 0:
   - If the character is `*`: Increment `starCounter`.
   - If the character is NOT `*`:
     - If `starCounter` > 0: Skip this character (it is removed) and decrement `starCounter`.
     - If `starCounter` == 0: Keep this character (append to `result`).
3. The `result` slice now contains the answer in **reverse order**.
4. Reverse `result` (or build the string in reverse) to get the final output.

## Complexity
Both approaches have the same theoretical complexity:
- **Time**: `O(n)` — We iterate through the string exactly once (and potentially another pass to reverse in Approach 2).
- **Space**: `O(n)` — To store the result.

## Implementation Details
- We use `[]byte` instead of `[]rune` for better performance since the logic doesn't require full UTF-8 decoding for this specific transformation.
- Pre-allocating the slice (`make([]byte, 0, len(s))`) avoids memory re-allocations during append operations.


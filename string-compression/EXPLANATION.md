# Explanation — String Compression

## Problem
Given an array of characters `chars`, compress it in-place by replacing consecutive repeating characters with the character followed by the count (if count > 1). Return the new length of the array.

Examples:
- `chars = ["a","a","b","b","c","c","c"]` -> `["a","2","b","2","c","3"]`, return 6
- `chars = ["a"]` -> `["a"]`, return 1
- `chars = ["a","b","b","b","b","b","b","b","b","b","b","b","b"]` -> `["a","b","1","2"]`, return 4

## Idea
Iterate through the input array `chars` and count consecutive repeating characters. Build the compressed result in a separate buffer (or in-place if optimized) and then update the original array.

## Algorithm (step-by-step)
1. Initialize an empty `result` slice with the first character of `chars`.
2. Initialize `pointer` to the first character and `count` to 1.
3. Iterate through `chars` starting from the second character:
   - If the current character matches `pointer`, increment `count`.
   - If it doesn't match:
     - If `count > 1`, append the string representation of `count` to `result`.
     - Reset `count` to 1.
     - Update `pointer` to the current character.
     - Append the current character to `result`.
4. After the loop, check if the last group had `count > 1` and append the count to `result` if so.
5. Copy the `result` slice back into `chars`.
6. Return the length of `result`.

## Complexity
- Time: O(n) — one pass to count and build the result.
- Space: O(n) — uses a separate `result` slice to build the compressed string. (Note: The problem asks for O(1) space, but this implementation uses O(n)).

## Implementation notes
- The current implementation uses a separate slice `result` for simplicity.
- To achieve O(1) space, one would need to write directly into `chars` using two pointers (read and write).

# Explanation — Equal Row and Column Pairs

## Problem
Given an `n x n` integer matrix `grid`, count the number of pairs `(r_i, c_j)` such that row `r_i` and column `c_j` are identical.

## Idea
Instead of comparing every row with every column which would be `O(n^3)`, we can use a hash map to store the rows. We convert each row into a unique string representation (or use an array/vector as key if supported) and count their occurrences. Then, we iterate through the columns, generate the same string representation, and check how many times that row appeared.

## Algorithm (step-by-step)
1. Initialize a hash map `rowsMap` to store the frequency of each row pattern.
2. Iterate through each row of the `grid`:
   - Convert the row elements into a unique string key (e.g., "1,2,3").
   - Increment the count for this key in `rowsMap`.
3. Initialize `pairs` counter to 0.
4. Iterate through each column of the `grid` (from index 0 to `n-1`):
   - Construct the string key for the current column by collecting elements `grid[row][col]`.
   - Check if this key exists in `rowsMap`.
   - If it exists, add the frequency of that key to `pairs`.
5. Return `pairs`.

## Complexity
- **Time**: `O(n^2)` — We iterate through the grid twice (once for rows, once for columns). String construction for each row/col takes `O(n)`, done `n` times, so total `O(n^2)`.
- **Space**: `O(n^2)` — In the worst case, we store `n` distinct rows, each of length `n` (stringified), in the map.

## Implementation notes
- We use a `strings.Builder` for efficient string concatenation when building keys.
- The key format uses commas to separate numbers to distinctively identify rows (e.g., `[1, 11]` vs `[11, 1]`).

# Explanation — Can Place Flowers

## Problem
Given a flowerbed (array of 0s and 1s), determine if you can plant `n` new flowers without violating the rule that no two flowers can be adjacent.

## Idea
Iterate through the flowerbed and greedily plant a flower at each position if possible (i.e., the current spot and its neighbors are empty). Skip the next spot after planting to avoid adjacency.

## Algorithm (step-by-step)
1. Iterate through the `flowerbed` array.
2. For each position `i`, check if a flower can be planted:
   - `flowerbed[i]` must be 0 (empty).
   - Left neighbor `i-1` must be 0 or out of bounds.
   - Right neighbor `i+1` must be 0 or out of bounds.
3. If all conditions are met:
   - Plant the flower: `flowerbed[i] = 1` (or just increment count).
   - Decrement `n`.
   - Skip the next position (since we can't plant adjacent).
4. If `n <= 0`, return `true`.
5. If the loop finishes and `n > 0`, return `false`.

## Complexity
- Time: O(n) — single pass through the array.
- Space: O(1) — only a counter is needed.

## Implementation notes
- The solution can be written to avoid mutating the input array for purity, but LeetCode allows mutation.
- The greedy approach is optimal: if you can plant a flower at a spot, you should do so immediately.
- Early exit is used for efficiency if n is reached before the end.

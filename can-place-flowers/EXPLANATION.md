# Can Place Flowers — Explanation

## Problem
Given a flowerbed (array of 0s and 1s), determine if you can plant `n` new flowers without violating the rule that no two flowers can be adjacent.

## High-level idea
Iterate through the flowerbed and greedily plant a flower at each position if possible (i.e., the current spot and its neighbors are empty). Skip the next spot after planting to avoid adjacency.

## Approach
- For each position, check:
  - Is the current spot empty?
  - Are both neighbors (or out-of-bounds) empty?
- If so, "plant" a flower (increment a counter or mark the spot) and skip the next position.
- Stop early if enough flowers are planted.

## Complexity
- Time: O(n) — single pass through the array.
- Space: O(1) — only a counter is needed (in-place mutation is allowed by LeetCode, but not required).

## Edge cases
- Single-element flowerbed.
- n = 0 (always true).
- All spots filled or all empty.
- Planting at the start or end of the array.

## Notes
- The solution can be written to avoid mutating the input array for purity, but LeetCode allows mutation.
- The greedy approach is optimal: if you can plant a flower at a spot, you should do so immediately.
- Early exit is used for efficiency if n is reached before the end.

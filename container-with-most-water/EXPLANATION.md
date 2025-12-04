# Explanation — Container With Most Water

## Problem
You are given an integer array `height` of length `n`. There are `n` vertical lines drawn such that the two endpoints of the `i`th line are `(i, 0)` and `(i, height[i])`. Find two lines that together with the x-axis form a container, such that the container contains the most water. Return the maximum amount of water a container can store.

## Idea
We want the maximum value of `(j - i) * min(height[i], height[j])` over all pairs `i < j`.

The implementation uses the standard two-pointer approach: place one pointer at the left end (`i`) and one at the right end (`j`) and move them toward each other, always moving the pointer at the smaller height. This finds the optimal pair in O(n) time with O(1) extra space.

## Algorithm (step-by-step)
1. Initialize `result = 0`.
2. Initialize two pointers: `i = 0` (left) and `j = len(height) - 1` (right).
3. Loop while `i < j`:
   - Calculate current area: `area = (j - i) * min(height[i], height[j])`.
   - Update `result = max(result, area)`.
   - Move the pointer pointing to the shorter line:
     - If `height[i] <= height[j]`, increment `i`.
     - Else, decrement `j`.
4. Return `result`.

## Complexity
- Time: O(n) — each pointer moves at most `n` steps combined.
- Space: O(1) — only a few integer variables are used.

## Implementation notes
- The code uses `height[i] <= height[j]` (less-than-or-equal). When the two heights are equal the left pointer `i` is moved inward. This tie-breaking choice is correct — either pointer could be moved when heights are equal.
- If `len(height) < 2`, the loop condition `i < j` is false initially, so it correctly returns 0.

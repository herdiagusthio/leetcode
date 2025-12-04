# Explanation — Find Closest Number to Zero

## Problem
Given an integer array `nums`, return the number that is closest to zero. If two numbers are equally close, return the larger one (for example, between `-3` and `3` return `3`).

## Idea
Scan the array once, tracking the current best candidate: the element with the smallest absolute value observed so far. When a new element has a strictly smaller absolute value it becomes the candidate. If absolute values tie, choose the numerically larger element (this enforces the tie-break rule: prefer the positive one when distances are equal).

## Algorithm (step-by-step)
1. Handle empty input: if `len(nums) == 0` return `0`.
2. Initialize trackers with the first element:
   - `closestDistance := absInt(nums[0])`
   - `largestValue := nums[0]`
3. Iterate over remaining elements:
   - For each `num`, compute `distance := absInt(num)`.
   - If `distance < closestDistance`:
     - Update `closestDistance = distance` and `largestValue = num`.
   - Else if `distance == closestDistance`:
     - Update `largestValue = max(largestValue, num)` (tie-breaker).
4. Return `largestValue`.

## Complexity
- Time: O(n) — single pass through the input slice.
- Space: O(1) — only a few scalars used for tracking.

## Implementation notes
- Using integer absolute values avoids float conversions and potential precision issues.
- Empty slice returns `0`.
- Zero present: zero has absolute distance 0 and will be returned.

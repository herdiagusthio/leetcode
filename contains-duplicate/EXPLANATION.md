# Explanation — Contains Duplicate

## Problem
Given an integer array `nums`, determine whether any value appears at least twice. Return `true` if a duplicate exists, otherwise return `false`.

Examples:
- `[1,2,3,1]` -> `true`
- `[1,2,3,4]` -> `false`

## Idea
Use a hash set to record numbers we have seen while iterating the array. For each number:
- If it's already in the set, we found a duplicate — return `true` immediately.
- Otherwise insert it into the set and continue.

This approach finds duplicates quickly (early return) and is O(n) time on average and O(n) extra space for the set.

## Algorithm (step-by-step)
1. If `len(nums) < 2`, return `false` (no duplicates possible).
2. Create an empty set: `seen := make(map[int]struct{}, len(nums))`.
3. Iterate over `nums`:
   - If `num` is in `seen`, return `true`.
   - Otherwise set `seen[num] = struct{}{}`.
4. If loop completes, return `false`.

## Complexity
- Time: O(n) average — each element is hashed once.
- Space: O(n) for the set.

## Alternatives
- Sort the array and check adjacent elements; O(n log n) time and O(1) extra space if you can sort in place.
- If memory is constrained and input range is small, a bitmap/bitset may be used.

## Implementation notes
- The solution uses `map[int]struct{}` as a set to avoid counting and minimize memory per entry.
- The function returns as soon as a duplicate is found to minimize work.

*** End Patch
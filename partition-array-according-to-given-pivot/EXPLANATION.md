# Explanation — Partition Array According to Given Pivot

## Problem
Given a 0-indexed integer array `nums` and an integer `pivot`, rearrange `nums` so that:
- Every element less than `pivot` appears before every element greater than `pivot`.
- Every element equal to `pivot` appears in between the elements less than and greater than `pivot`.
- The relative order of the elements less than `pivot` and the elements greater than `pivot` is maintained.

## Idea
A simple, stable approach is to collect elements into three buckets while scanning the array:
- `left` for elements < pivot
- `mid` for elements == pivot
- `right` for elements > pivot

Then concatenate `left + mid + right`. This preserves the relative order inside each partition.

## Algorithm (step-by-step)
1. Initialize three empty slices: `left`, `mid`, `right`.
2. Iterate through `nums`:
   - If `num < pivot`, append to `left`.
   - If `num == pivot`, append to `mid`.
   - If `num > pivot`, append to `right`.
3. Concatenate `left`, `mid`, and `right` to form the result.
4. Return the result.

## Complexity
- Time: O(n) — single pass (or two passes for optimized variants).
- Space: O(n) extra — buckets or one preallocated result slice.

## Alternatives
- **Preallocated buckets**: Same logic but preallocates capacities to reduce reallocations.
- **Two-pass preallocate-and-fill**: First pass counts numbers to compute exact size, second pass fills the result. Best for minimizing allocations.
- **In-place DNF**: The Dutch National Flag algorithm is O(1) space but unstable (does not preserve relative order), so it is not suitable here.

## Implementation notes
- For production code where minimizing allocations matters, prefer the two-pass approach.
- The `pivot` is guaranteed to exist in `nums`.

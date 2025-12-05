# Explanation — Move Zeroes

## Problem
Given an integer array `nums`, move all 0's to the end of it while maintaining the relative order of the non-zero elements. You must do this in-place without making a copy of the array.

Examples:
- `nums = [0,1,0,3,12]` -> `[1,3,12,0,0]`
- `nums = [0]` -> `[0]`

## Idea
Use a two-pointer approach: maintain a pointer that tracks the position where the next non-zero element should be placed. Scan through the array and swap non-zero elements with the element at the pointer position, then advance the pointer.

This ensures non-zero elements maintain their relative order and all zeros end up at the end.

## Algorithm (step-by-step)
1. Initialize `pointer := 0`.
2. For each index `i` from `0` to `len(nums)-1`:
   - If `nums[i] != 0`, swap `nums[pointer]` with `nums[i]` and increment `pointer`.
3. After the loop completes, all non-zero elements are at the front in their original order, and zeros are at the end.

## Complexity
- Time: O(n) — single pass through the array.
- Space: O(1) — in-place modification with only a pointer variable.

## Alternatives
- Count non-zeros, write them to the front, then fill remaining positions with zeros (two passes, similar complexity).
- Use a separate array to collect non-zeros then append zeros (not in-place, uses O(n) extra space).

## Implementation notes
- The swap operation works even when `pointer == i` (swapping with itself), so no special case is needed.
- This approach is optimal for the in-place constraint and maintains stability of non-zero elements.

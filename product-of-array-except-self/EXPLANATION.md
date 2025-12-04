# Explanation — Product of Array Except Self

## Problem
Given an integer array `nums`, return an array `answer` such that `answer[i]` is equal to the product of all the elements of `nums` except `nums[i]`. You must run in O(n) time and not use division.

Examples:
- `nums = [1,2,3,4]` -> `answer = [24,12,8,6]`
- `nums = [-1,1,0,-3,3]` -> `answer = [0,0,9,0,0]`

## Idea
Compute the product of elements to the left (prefix) and to the right (suffix) of each index and multiply them. Do this in two passes using only a constant amount of extra working space (besides the output array):

- First pass (left to right): fill `answer[i]` with the product of all elements left of `i`.
- Second pass (right to left): multiply `answer[i]` by the product of all elements right of `i`.

This avoids division and handles zeros naturally.

## Algorithm (step-by-step)
1. Let `n := len(nums)` and allocate `answer := make([]int, n)`.
2. Set `prefix := 1`. For `i` from `0` to `n-1`:
   - Set `answer[i] = prefix`.
   - Update `prefix *= nums[i]`.
3. Set `suffix := 1`. For `i` from `n-1` down to `0`:
   - Multiply `answer[i] *= suffix`.
   - Update `suffix *= nums[i]`.
4. Return `answer`.

## Complexity
- Time: O(n) — two linear passes over `nums`.
- Space: O(n) for the output array `answer`. Extra working space is O(1) (the `prefix` and `suffix` variables).

## Alternatives
- Using division: compute the product of all elements and divide by `nums[i]` to get `answer[i]`, but division is disallowed and also fails when zeros are present.
- Use additional arrays to store prefix and suffix products separately; this is less space-efficient (still O(n)).

## Implementation notes
- The two-pass approach handles zeros correctly: if more than one zero appears, most answers will be zero; if exactly one zero appears, only the index of the zero will have non-zero product (the product of other elements).
- The solution uses `int` for simplicity; the problem statement guarantees the products fit in 32-bit integers, but in Go `int` is platform-dependent (typically 64-bit on modern machines). If you need to strictly match 32-bit behavior, consider using `int32` where appropriate.

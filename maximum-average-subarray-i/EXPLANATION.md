# Explanation — Maximum Average Subarray I

Problem summary
- Given an integer array `nums` of length `n` and an integer `k`, find a contiguous subarray
  of length `k` that has the maximum average value and return this value.

Approach
- Use the sliding window technique to efficiently compute the maximum sum of any subarray
  of length `k`.
- First, compute the sum of the initial window (first `k` elements).
- Then slide the window one element at a time: add the next element and subtract the element
  that's leaving the window.
- Track the maximum sum encountered and return `maxSum / k` as the maximum average.

Why this is correct
- The maximum average subarray of length `k` is equivalent to finding the maximum sum
  subarray of length `k`, since dividing by `k` (a constant) preserves the ordering.
- The sliding window ensures we examine every possible subarray of length `k` in a single pass,
  updating the sum in O(1) per step.

Complexity
- Time: O(n) — we compute the initial sum in O(k), then slide the window (n - k) times with
  O(1) operations per step.
- Space: O(1) — only a few integer/float variables are used.

Notes and edge cases
- Works correctly when `k = 1` (return the maximum element).
- Works correctly when `k = n` (return the average of the entire array).
- Handles negative numbers correctly since we're tracking the maximum sum, not just positive values.

Small example
- Input: `nums = [1, 12, -5, -6, 50, 3]`, `k = 4`
  - Initial window sum: 1 + 12 + (-5) + (-6) = 2
  - Slide to [12, -5, -6, 50]: sum = 2 - 1 + 50 = 51 (new max)
  - Slide to [-5, -6, 50, 3]: sum = 51 - 12 + 3 = 42
  - Maximum sum = 51, maximum average = 51 / 4 = 12.75

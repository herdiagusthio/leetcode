# Explanation — Container With Most Water

This document explains the `maxArea` implementation in `solution.go` and ensures the explanation matches the actual code.

## Contract (inputs / outputs)
- Input: `height []int` — slice of non-negative integers representing vertical line heights at consecutive x positions.
- Output: `int` — the maximum area (width * limiting height) obtainable by choosing two indices `i < j`. Returns `0` when `len(height) < 2`.

## High-level idea
We want the maximum value of `(j - i) * min(height[i], height[j])` over all pairs `i < j`.

The implementation uses the standard two-pointer approach: place one pointer at the left end (`i`) and one at the right end (`j`) and move them toward each other, always moving the pointer at the smaller height. This finds the optimal pair in O(n) time with O(1) extra space.

## Implementation details (matching the code)
The code initializes these variables exactly as in `solution.go`:

- `result, area, i, j := 0, 0, 0, len(height)-1`

Loop behavior (exactly as implemented):

- `for i < j { ... }`
  - If `height[i] <= height[j]` then the code sets `area = height[i] * (j - i)` and increments `i`.
  - Otherwise the code sets `area = height[j] * (j - i)` and decrements `j`.
  - After computing `area`, the code updates `result` when `area > result`.

Notes about the comparison and tie-breaker:

- The code uses `height[i] <= height[j]` (less-than-or-equal). When the two heights are equal the left pointer `i` is moved inward. This tie-breaking choice is correct — either pointer could be moved when heights are equal; moving the left pointer is a valid deterministic choice and does not affect correctness.

Main function:

- The repository's `solution.go` is in `package main` and includes a `main()` function that demonstrates `maxArea` on several example inputs and prints the results. The `maxArea` function itself is usable independently and is covered by unit tests in `solution_test.go`.

## Example walkthrough (height = [1,8,6,2,5,4,8,3,7])
- Start: `i=0 (h=1)`, `j=8 (h=7)` -> `area = 1 * 8 = 8` -> `result = 8` -> `height[i] <= height[j]` so `i++`
- Next: `i=1 (h=8)`, `j=8 (h=7)` -> `area = 7 * 7 = 49` -> `result = 49` -> `height[i] > height[j]` so `j--`
- The loop continues until `i >= j`; final `result = 49`.

## Complexity
- Time complexity: O(n) — each pointer moves at most `n` steps combined.
- Space complexity: O(1) — only a few integer variables are used.

## Edge cases and notes
- If `len(height) < 2`, `j := len(height)-1` will make the loop condition `i < j` false, so the function returns `0` (the code handles this implicitly).
- All-zero heights -> `0`.
- Monotonic sequences (increasing or decreasing) are handled correctly.
- Given constraints (`height[i] <= 10^4` and `n <= 10^5`), maximum area ≤ 10^9 which fits in a 32-bit signed `int`. If your input domain grows, consider using `int64` for safety.

## Tests
The `solution_test.go` in this folder includes unit tests covering the LeetCode examples and additional edge cases.

## Final note
The explanation now precisely matches the actual code: variable initialization, the `<=` tie-breaker, the `area` update, and the included `main` demonstration. The approach is the standard optimal solution: simple, fast, and memory-efficient.

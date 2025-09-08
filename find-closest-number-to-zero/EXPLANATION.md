# Explanation — Find Closest Number to Zero

This document contains a step-by-step explanation of the solution implemented in `solution.go`.

## Contract
- Input: `nums []int` — a slice of integers (may contain negative, zero, positive values).
- Output: `int` — an element from `nums` that is closest to zero. If two elements have the same distance to zero, the larger numerical value is returned. If `nums` is empty the implementation returns `0`.

## High-level idea
Scan the array once, tracking the current best candidate: the element with the smallest absolute value observed so far. When a new element has a strictly smaller absolute value it becomes the candidate. If absolute values tie, choose the numerically larger element (this enforces the tie-break rule: prefer the positive one when distances are equal).

## Why this is easy and robust
- Single pass (O(n)) and constant extra memory (O(1)).
- Using integer absolute values avoids float conversions and potential precision issues.

## Step-by-step (matching `solution.go`)
1. Handle empty input:
   - The code has `if len(nums) == 0 { return 0 }` to return quickly for an empty slice.

2. Initialize trackers with the first element:
   - `closestDistance := absInt(nums[0])` — integer absolute distance of the first element.
   - `largestValue := nums[0]` — the current chosen value.
   Initializing from the first element avoids special cases inside the loop.

3. Iterate over remaining elements:
   - For each `num` compute `distance := absInt(num)`.

4. Update rule (the key line):
   - `if distance < closestDistance || (distance == closestDistance && num > largestValue) { ... }`
   - Two situations cause an update:
     - `distance < closestDistance`: strictly closer to zero → update.
     - `distance == closestDistance && num > largestValue`: equal distance but numerically larger → update (tie-breaker).

5. Return the chosen value:
   - After the loop the function returns `largestValue` which is the element from the original slice that satisfies the problem constraint.

## Example walkthrough
Input: `nums = [-4, -1, 1, 4, 8]`
- Start with first element `-4` → closestDistance = 4, largestValue = -4
- Next `-1` → distance = 1 < 4 → update: closestDistance = 1, largestValue = -1
- Next `1` → distance = 1 == 1 and `1 > -1` → update: largestValue = 1
- Next `4` → distance = 4 > 1 → ignore
- Next `8` → distance = 8 > 1 → ignore
Return `1`.

## Complexity
- Time: O(n) — single pass through the input slice.
- Space: O(1) — only a few scalars used for tracking.

## Edge cases and behavior
- Empty slice: returns `0` (explicit in implementation).
- Zero present: zero has absolute distance 0 and will be returned.
- Equal absolute values (e.g., `-x` and `x`): the function returns the larger one (e.g., `x`).

## Tests and verification
- There is a small table-driven test in `solution_test.go` that verifies the example, empty input, zero presence, and tie cases. Run `go test` as described in the `README.md` to execute them.

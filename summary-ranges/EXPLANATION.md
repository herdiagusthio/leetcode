# Explanation — Summary Ranges

## Problem
Given a sorted, unique integer array, summarize consecutive number ranges.
Examples:
- `[0,1,2,4,5,7]` -> `["0->2","4->5","7"]`
- `[]` -> `[]`

## Idea
Scan the array once and maintain the current range as `[start, end]`:
- If the next number equals `end + 1`, extend the range (`end = next`).
- Otherwise the current range ends: append it to output (format as `"start"` or `"start->end"`), then start a new range.

At the end, append the final range.

## Algorithm (step-by-step)
1. If `nums` is empty, return empty list.
2. Initialize `start` and `end` with `nums[0]`.
3. Iterate from `i = 1` to `len(nums)-1`:
   - If `nums[i] == end + 1`, update `end = nums[i]`.
   - Else, format range `start` (and `end` if different) and append to result. Reset `start` and `end` to `nums[i]`.
4. Append the final range.
5. Return result.

## Complexity
- Time: O(n) — single pass over the input.
- Space: O(n) — output.

## Implementation notes
- We preallocate the output slice with capacity `len(nums)` because in the worst case each number becomes its own range.
- We use `strconv.Itoa` to convert integers to strings rather than `fmt.Sprintf` to reduce allocations in hot paths.
- Since the input is guaranteed unique and sorted, no special duplicate handling is necessary.

# Explanation — Summary Ranges


## Problem
Given a sorted, unique integer array, summarize consecutive number ranges.

Examples:
- `[0,1,2,4,5,7]` -> `["0->2","4->5","7"]`
- `[]` -> `[]`

## Key idea

Scan the array once and maintain the current range as `[start, end]`:
- If the next number equals `end + 1`, extend the range (`end = next`).
- Otherwise the current range ends: append it to output (format as `"start"` or `"start->end"`), then start a new range.

At the end, append the final range.

## Implementation details
- We preallocate the output slice with capacity `len(nums)` because in the worst case each number becomes its own range.
- We use `strconv.Itoa` to convert integers to strings rather than `fmt.Sprintf` to reduce allocations in hot paths.
 - The implementation assumes the input is sorted and contains unique integers (per the problem statement).

## Complexity
- Time: O(n) — single pass over the input.
- Space: O(n) output.

## Notes
 - Since the input is guaranteed unique and sorted, no special duplicate handling is necessary.
- For heavy workloads, benchmarks (included) compare `strconv` vs `fmt.Sprintf` and show differences in allocations.

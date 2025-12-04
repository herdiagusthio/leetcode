# Explanation — Sort Colors

## Problem
Given an array containing only 0, 1, and 2, reorder it in-place so that all 0s come first, then 1s, then 2s.

## Idea
The problem asks for a single-pass, in-place sort. The canonical solution is the **Dutch National Flag (DNF)** algorithm.

Maintain three pointers:
- `low`: boundary between processed 0s and others (next position to place 0)
- `mid`: current index being inspected
- `high`: boundary between processed 2s and the remainder (next position to place 2)

## Algorithm (step-by-step)
Start with `low = 0`, `mid = 0`, `high = n-1`. While `mid <= high`:
1. If `nums[mid] == 0`: swap `nums[low]` and `nums[mid]`; `low++`; `mid++`.
2. If `nums[mid] == 1`: `mid++`.
3. If `nums[mid] == 2`: swap `nums[mid]` and `nums[high]`; `high--`.

Each swap places at least one element in its final region; `mid` moves forward when we've handled a position, and `high` shrinks when we've moved a 2 out.

## Complexity
- Time: O(n) — single pass.
- Space: O(1) — constant pointers only.

## Alternatives
- **Count-array overwrite**: Count how many 0s, 1s, and 2s appear, then overwrite the input array with that many 0s, followed by 1s, then 2s. This is two passes but still O(n) time and O(1) space.

## Implementation notes
- Empty slice: no-op for both approaches.
- Single color: handled trivially.
- Two colors only: both approaches handle missing colors correctly.
- Input validation: the LeetCode problem guarantees only 0/1/2.

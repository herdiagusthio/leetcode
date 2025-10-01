# Explanation — Sort Colors

This document explains the two approaches included in `solution.go` for the "Sort Colors" problem (LeetCode).

Problem restatement
-------------------
Given an array containing only 0, 1, and 2, reorder it in-place so that all 0s come first, then 1s, then 2s.

Approach 1 — Dutch National Flag (DNF)
-----------------------------------
Idea
~~~~
Maintain three pointers:

- low: boundary between processed 0s and others (next position to place 0)
- mid: current index being inspected
- high: boundary between processed 2s and the remainder (next position to place 2)

Algorithm
~~~~~~~~~
Start with low = 0, mid = 0, high = n-1. While mid <= high:

- If nums[mid] == 0: swap nums[low] and nums[mid]; low++; mid++
- If nums[mid] == 1: mid++
- If nums[mid] == 2: swap nums[mid] and nums[high]; high--

Why it works
~~~~~~~~~~~~
Each swap places at least one element in its final region; mid moves forward when we've handled a position, and high shrinks when we've moved a 2 out. This processes each element at most once.

Complexity
~~~~~~~~~~
- Time: O(n) — single pass
- Space: O(1) — constant pointers only

Approach 2 — Count-array overwrite
----------------------------------
Idea
~~~~
Count how many 0s, 1s, and 2s appear, then overwrite the input array with that many 0s, followed by 1s, then 2s.

Algorithm
~~~~~~~~~
1. One pass to compute counts[0], counts[1], counts[2].
2. Overwrite the array in order using the counts.

Complexity
~~~~~~~~~~
- Time: O(n) — two passes (counting + overwrite)
- Space: O(1) — fixed-size count array

Edge cases and notes
--------------------
- Empty slice: no-op for both approaches.
- Single color: handled trivially.
- Two colors only: both approaches handle missing colors correctly.
- Input validation: the LeetCode problem guarantees only 0/1/2; if your code may see other values, define behavior (we advance mid in DNF's default case to avoid infinite loops during debugging).

Which to use?
--------------
- Use DNF (`sortColors`) when you want the canonical, single-pass in-place algorithm.
- Use count-array (`sortColorsCount`) when you prefer a simpler two-pass implementation that is also easy to reason about.

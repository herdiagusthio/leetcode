# Sort Colors

LeetCode: https://leetcode.com/problems/sort-colors/

Problem
-------
Given an array nums with n objects colored red, white, or blue, sort them in-place so that objects of the same color are adjacent, with the colors in the order red, white, and blue.

We use integers 0, 1, and 2 to represent the color red, white, and blue respectively. You must solve this without using a library sort.

Examples
--------
- Input: [2,0,2,1,1,0]
  Output: [0,0,1,1,2,2]

- Input: [2,0,0,2,0]
  Output: [0,0,0,2,2]

Approaches included in this folder
---------------------------------
- Dutch National Flag (DNF) — `sortColors` (recommended): single-pass, in-place, O(n) time and O(1) extra space. Canonical solution.
- Count-array overwrite — `sortColorsCount`: count occurrences of 0/1/2, then overwrite the slice. Two passes but still O(n) time and O(1) extra space.

Files
-----
- `solution.go` — implementations for `sortColors` (DNF) and `sortColorsCount` (count-array), plus a small `main()` demonstration.
- `solution_test.go` — table-driven tests that exercise both implementations.

Complexity
----------
- Time: O(n)
- Space: O(1) extra

Notes
-----
- Inputs are guaranteed to be 0, 1, or 2 per the problem constraints. The DNF implementation includes a defensive `default` branch that advances the mid pointer for unexpected values (useful during debugging); LeetCode inputs will not trigger this.
- To compare performance empirically, add a benchmark in `*_test.go`.

Constraints
-----------
- n == nums.length
- 1 <= n <= 300
- nums[i] is either 0, 1, or 2.

# Container With Most Water

LeetCode problem: Container With Most Water

Link: https://leetcode.com/problems/container-with-most-water/

You are given an integer array `height` of length `n`. There are `n` vertical lines drawn such that the two endpoints of the `i`th line are `(i, 0)` and `(i, height[i])`.

Find two lines that together with the x-axis form a container, such that the container contains the most water. Return the maximum amount of water a container can store. You may not slant the container.

## Solution
Use the two-pointer technique:
- Initialize `left = 0` and `right = n-1`.
- While `left < right`, compute `area = (right-left) * min(height[left], height[right])` and update the best result.
- Move the pointer that points to the smaller height inward, since moving the taller one cannot produce a larger area with the same width.

This runs in O(n) time and uses O(1) extra space.

## How to run
From this folder:

```bash
go test -v
```

To run the example `main` in this package:

```bash
go run .
```

## Files
- `solution.go` — implementation and example `main`.
- `solution_test.go` — unit tests.
- `EXPLANATION.md` — step-by-step explanation of the approach.

# Sort Colors

LeetCode problem: Sort Colors

Link: https://leetcode.com/problems/sort-colors/

Given an array nums with n objects colored red, white, or blue, sort them in-place so that objects of the same color are adjacent, with the colors in the order red, white, and blue. We use integers 0, 1, and 2 to represent the color red, white, and blue respectively. You must solve this without using a library sort.

## Solution
The canonical solution is the **Dutch National Flag (DNF)** algorithm, which sorts the array in a single pass using O(1) extra space.

This folder contains two implementations:
- `sortColors` (DNF): single-pass, in-place, O(n) time.
- `sortColorsCount`: two-pass count-and-overwrite, O(n) time.

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
- `solution.go` — implementations and example `main`.
- `solution_test.go` — unit tests.
- `EXPLANATION.md` — step-by-step explanation of the approach.

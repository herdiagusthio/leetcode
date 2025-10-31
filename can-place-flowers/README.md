LeetCode: https://leetcode.com/problems/can-place-flowers/

## Problem

You have a long flowerbed in which some of the plots are planted, and some are not. However, flowers cannot be planted in adjacent plots.

Given an integer array `flowerbed` containing `0`'s and `1`'s, where `0` means empty and `1` means not empty, and an integer `n`, return `true` if `n` new flowers can be planted in the flowerbed without violating the no-adjacent-flowers rule and `false` otherwise.

## Example

Input: `flowerbed = [1,0,0,0,1]`, `n = 1`

Output: `true`

Input: `flowerbed = [1,0,0,0,1]`, `n = 2`

Output: `false`

## Constraints

- `1 <= flowerbed.length <= 2 * 10^4`
- `flowerbed[i]` is `0` or `1`.
- There are no two adjacent flowers in `flowerbed`.
- `0 <= n <= flowerbed.length`

## How to run

Run the sample program or tests for this folder:

```bash
cd can-place-flowers
go run .        # run the sample main
go test         # run unit tests
```

## Files

- `solution.go` — implementation and small demo `main`.
- `solution_test.go` — unit tests.
- `EXPLANATION.md` — explanation and notes.
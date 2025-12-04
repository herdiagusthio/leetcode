# Best Time to Buy and Sell Stock

LeetCode problem: Best Time to Buy and Sell Stock

Link: https://leetcode.com/problems/best-time-to-buy-and-sell-stock/

Given an array `prices` where `prices[i]` is the price of a given stock on day `i`, return the maximum profit you can achieve from one transaction (buy once and sell once). If you cannot achieve any profit, return 0.

## Solution
Scan once, track the minimum price seen so far and the best profit achievable at each step. This approach runs in O(n) time and uses O(1) space.

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

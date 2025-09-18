# Best Time to Buy and Sell Stock

LeetCode: https://leetcode.com/problems/best-time-to-buy-and-sell-stock/

## Problem
Given an array `prices` where `prices[i]` is the price of a given stock on day `i`,
return the maximum profit you can achieve from one transaction (buy once and sell once).
If you cannot achieve any profit, return 0.

## Solution (short)
Scan once, track the minimum price seen so far and the best profit achievable at each step.

## Complexity
- Time: O(n)
- Space: O(1)

## How to run

```bash
go run solution.go
```

## Run tests

```bash
go test
```

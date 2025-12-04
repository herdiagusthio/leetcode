# Explanation — Best Time to Buy and Sell Stock

## Problem
Given an array `prices` where `prices[i]` is the price of a given stock on day `i`, return the maximum profit you can achieve from one transaction (buy once and sell once). If you cannot achieve any profit, return 0.

## Idea
We want the largest difference `prices[j] - prices[i]` such that `j > i`. Instead of checking all pairs (O(n^2)), scan once keeping the smallest price seen so far (the best day to buy before the current day). For each day, compute the profit if we sold that day using the smallest prior price, and update the best profit seen.

This reduces the problem to an O(n) single-pass algorithm using O(1) extra space.

## Algorithm (step-by-step)
1. Edge-case check: `if len(prices) < 2 { return 0 }` — if there are fewer than two days you cannot complete a buy+sell pair.
2. Initialize tracking variables:
   - `minPrice := prices[0]` — the smallest price observed so far (candidate buy price).
   - `maxProfit := 0` — the best profit we've observed.
3. Loop from day 1 to the end:
   - `price := prices[i]` — price on current day (candidate sell day).
   - If `price < minPrice`, update `minPrice = price` (we found a better buy day).
   - Else compute `profit := price - minPrice` and, if `profit > maxProfit`, update `maxProfit = profit`.
4. After the loop return `maxProfit`.

## Complexity
- Time: O(n) — single pass through `prices`.
- Space: O(1) — only two ints used regardless of input size.

## Alternatives
- Brute force: Check all pairs `(i, j)` where `i < j`. This is O(n^2) and too slow for large inputs.
- Kadane's Algorithm: The problem can be transformed into finding the maximum subarray sum of the differences between adjacent prices.

## Implementation notes
- The solution uses `int` for simplicity.
- Empty array or single-element: returns 0.
- Monotonic decreasing prices: returns 0 (no profitable transaction).

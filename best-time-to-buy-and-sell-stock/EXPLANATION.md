
# Explanation — Best Time to Buy and Sell Stock

This document provides a step-by-step explanation of the `maxProfit` implementation in `solution.go`.

## Contract (inputs / outputs)
- Input: `prices []int` — slice of non-negative integers representing stock prices by day.
- Output: `int` — the maximum profit obtainable by one buy followed by one sell (buy day < sell day). Returns `0` when no positive profit is possible or when `len(prices) < 2`.

## High-level idea
We want the largest difference `prices[j] - prices[i]` such that `j > i`. Instead of checking all pairs (O(n^2)), scan once keeping the smallest price seen so far (the best day to buy before the current day). For each day, compute the profit if we sold that day using the smallest prior price, and update the best profit seen.

This reduces the problem to an O(n) single-pass algorithm using O(1) extra space.

## Step-by-step walkthrough of the code
1. Edge-case check: `if len(prices) < 2 { return 0 }` — if there are fewer than two days you cannot complete a buy+sell pair.

2. Initialize tracking variables:
	- `minPrice := prices[0]` — the smallest price observed so far (candidate buy price).
	- `maxProfit := 0` — the best profit we've observed.

3. Loop from day 1 to the end (index-based loop avoids slice header allocations):
	- `price := prices[i]` — price on current day (candidate sell day).
	- If `price < minPrice`, update `minPrice = price` (we found a better buy day).
	- Else compute `profit := price - minPrice` and, if `profit > maxProfit`, update `maxProfit = profit`.

4. After the loop return `maxProfit`.

## Example walkthrough (prices = [7,1,5,3,6,4])
- Start: minPrice = 7, maxProfit = 0
- Day 1 (price=1): price < minPrice → minPrice = 1
- Day 2 (price=5): profit = 5 - 1 = 4 → maxProfit = 4
- Day 3 (price=3): profit = 3 - 1 = 2 → maxProfit stays 4
- Day 4 (price=6): profit = 6 - 1 = 5 → maxProfit = 5
- Day 5 (price=4): profit = 4 - 1 = 3 → maxProfit stays 5
- Return 5 (buy at 1, sell at 6)

## Complexity
- Time complexity: O(n) — single pass through `prices`.
- Space complexity: O(1) — only two ints used regardless of input size.

## Edge cases and subtle points
- Empty array or single-element: returns 0.
- Monotonic decreasing prices: returns 0 (no profitable transaction).

## Small improvements and alternatives
- Micro-optimization: an index loop (`for i := 1; i < len(prices); i++`) avoids forming `prices[1:]` and is slightly faster and allocation-free; this is what the implementation uses.
- If you needed to support multiple transactions (buy/sell multiple times), use another strategy (sum positive differences) — not applicable here.

## Tests
The `solution_test.go` file included in the folder covers typical cases: example, descending prices, small gains, empty input, single-day input, and two-day profit.

## Final note
This is the standard and optimal solution for this problem: correct, simple, and efficient.


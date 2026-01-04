# Best Time to Buy and Sell Stock - Explanation

## Problem Summary
Given an array `prices` where `prices[i]` is the price of a stock on day `i`, return the maximum profit you can achieve from one transaction (buy once, sell once). If no profit is possible, return 0.

## Approach
This solution uses a **single-pass greedy algorithm** with minimum price tracking:

1. **Initialize Trackers**: Start with the first price as minimum and zero profit
2. **Iterate Through Prices**: For each day starting from day 1:
   - If current price is lower than minimum, update the minimum (better buy opportunity)
   - Otherwise, calculate profit if we sell today and update maximum profit if better
3. **Return Result**: The maximum profit found

## Why This Works
The key insight is that to maximize profit, we need the largest difference `prices[j] - prices[i]` where `j > i`:
- **We must buy before we sell**, so we track the minimum price seen so far as the best buy candidate
- **For each potential sell day**, we calculate profit using the best (lowest) buy price before it
- **Greedy approach is optimal** because:
  - The best buy price is always the minimum seen before the current day
  - We don't need to look ahead - if a better sell day exists later, we'll find it when we get there
  - By tracking the running minimum, we ensure we're always considering the best possible buy for each sell

## Time Complexity
**O(n)** where n is the length of the prices array
- Single pass through the array
- Each element is processed exactly once
- Constant time operations (comparison, arithmetic) for each element
- No nested loops or additional iterations

## Space Complexity
**O(1)** - constant space
- Only two integer variables (`minPrice`, `maxProfit`)
- No data structures that grow with input size
- No recursion stack

## Edge Cases
1. **Less than 2 prices**: No transaction possible, return 0
2. **Decreasing prices**: No profit possible (buy high, sell low), return 0
3. **Increasing prices**: Maximum profit is (last price - first price)
4. **Single peak**: Buy at lowest before peak, sell at peak
5. **Multiple peaks**: Algorithm finds global maximum profit
6. **All same price**: No profit, return 0

## Worked Example
For `prices = [7,1,5,3,6,4]`:

```
Initial state:
minPrice = 7
maxProfit = 0

Day 1 (price = 1):
1 < 7 → Update minPrice = 1
maxProfit = 0

Day 2 (price = 5):
5 > 1 → Calculate profit = 5 - 1 = 4
maxProfit = max(0, 4) = 4

Day 3 (price = 3):
3 > 1 → Calculate profit = 3 - 1 = 2
maxProfit = max(4, 2) = 4 (no change)

Day 4 (price = 6):
6 > 1 → Calculate profit = 6 - 1 = 5
maxProfit = max(4, 5) = 5

Day 5 (price = 4):
4 > 1 → Calculate profit = 4 - 1 = 3
maxProfit = max(5, 3) = 5 (no change)

Result: 5 (Buy at price 1, sell at price 6)
```

## Visual Representation
```
Prices:  7   1   5   3   6   4
Day:     0   1   2   3   4   5
         
         ^   ^           ^
        bad  BUY        SELL

Min tracking:  7 → 1 → 1 → 1 → 1 → 1
Profit calc:   - → - → 4 → 2 → 5 → 3
Max profit:    0 → 0 → 4 → 4 → 5 → 5

Optimal: Buy on day 1 (price=1), sell on day 4 (price=6)
Profit = 6 - 1 = 5
```

## Example: No Profit Possible
For `prices = [7,6,4,3,1]`:

```
Day 0: minPrice = 7, maxProfit = 0
Day 1: 6 < 7 → minPrice = 6, maxProfit = 0
Day 2: 4 < 6 → minPrice = 4, maxProfit = 0
Day 3: 3 < 4 → minPrice = 3, maxProfit = 0
Day 4: 1 < 3 → minPrice = 1, maxProfit = 0

Result: 0 (prices only decrease, no profit possible)
```

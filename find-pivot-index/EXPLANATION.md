# Find Pivot Index - Explanation

## Problem Summary
Given an array of integers `nums`, find the leftmost pivot index where the sum of all numbers strictly to the left of the index equals the sum of all numbers strictly to the right. Return -1 if no such index exists.

## Approach
This solution uses a **two-pass algorithm** with mathematical optimization:

1. **Calculate Total Sum**: First pass computes the sum of all elements
2. **Find Pivot**: Second pass iterates through each index:
   - Track cumulative left sum
   - Calculate right sum using: `right = totalSum - left - nums[i]`
   - If left equals right, return the index
   - Update left sum for next iteration
3. **Return -1**: If no pivot found after checking all indices

## Why This Works
The key insight is avoiding redundant calculations:
- Instead of calculating left and right sums separately for each index, we use the relationship:
  - `left + nums[i] + right = totalSum`
  - Therefore: `right = totalSum - left - nums[i]`
- This allows us to compute the right sum in O(1) time at each index
- We maintain a running left sum as we iterate, also O(1) per index

## Time Complexity
**O(n)** where n is the length of the nums array
- First pass: O(n) to calculate total sum
- Second pass: O(n) to find pivot index
- Overall: O(n) + O(n) = O(n)

## Space Complexity
**O(1)** - We only use constant extra space:
- Three integer variables (`left`, `right`, `totalSum`)
- No data structures that grow with input size

## Edge Cases
1. **Single element** (e.g., [5]): Left sum = 0, right sum = 0, so index 0 is pivot
2. **Pivot at start** (e.g., [0,1,-1]): Left sum = 0 by definition
3. **Pivot at end** (e.g., [-1,1,0]): Right sum = 0 by definition
4. **No pivot** (e.g., [1,2,3]): Return -1
5. **Multiple pivots**: Return the leftmost one (first found)
6. **Negative numbers**: Algorithm works correctly with any integer values
7. **All zeros**: First index (0) is the pivot

## Worked Example
For `nums = [1,7,3,6,5,6]`:

```
Step 1: Calculate totalSum
totalSum = 1 + 7 + 3 + 6 + 5 + 6 = 28

Step 2: Find pivot
left = 0

i = 0, nums[0] = 1
right = 28 - 0 - 1 = 27
left (0) ≠ right (27)
left = 0 + 1 = 1

i = 1, nums[1] = 7
right = 28 - 1 - 7 = 20
left (1) ≠ right (20)
left = 1 + 7 = 8

i = 2, nums[2] = 3
right = 28 - 8 - 3 = 17
left (8) ≠ right (17)
left = 8 + 3 = 11

i = 3, nums[3] = 6
right = 28 - 11 - 6 = 11
left (11) == right (11) ✓
Return 3
```

Verification:
- Left sum: nums[0] + nums[1] + nums[2] = 1 + 7 + 3 = 11
- Right sum: nums[4] + nums[5] = 5 + 6 = 11

Result: **3**

## Example: No Pivot
For `nums = [1,2,3]`:

```
totalSum = 1 + 2 + 3 = 6

i = 0, nums[0] = 1
left = 0, right = 6 - 0 - 1 = 5
0 ≠ 5

i = 1, nums[1] = 2
left = 1, right = 6 - 1 - 2 = 3
1 ≠ 3

i = 2, nums[2] = 3
left = 3, right = 6 - 3 - 3 = 0
3 ≠ 0

No pivot found
```

Result: **-1**

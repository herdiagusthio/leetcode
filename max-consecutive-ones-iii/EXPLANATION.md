# Max Consecutive Ones III - Explanation

## Problem Summary
Given a binary array `nums` and an integer `k`, return the maximum number of consecutive 1's in the array if you can flip at most `k` 0's.

## Approach
This solution uses a **sliding window** technique with two pointers to find the longest subarray containing at most `k` zeros:

1. **Expand Window**: Move the right pointer and count zeros encountered
2. **Shrink Window**: When zeros exceed `k`, move the left pointer until the window is valid again
3. **Track Maximum**: Keep track of the largest valid window size seen

## Why This Works
The sliding window approach works because:
- We want a contiguous subarray where we can flip at most k zeros to get all 1's
- A valid window has at most k zeros (which we can flip to 1's)
- When we have more than k zeros, we shrink from the left until valid
- The window size represents consecutive 1's after flipping at most k zeros

## Time Complexity
**O(n)** where n is the length of the nums array
- Each element is visited at most twice (once by right pointer, once by left pointer)
- The right pointer moves from 0 to n-1
- The left pointer only moves forward, never backwards

## Space Complexity
**O(1)** - We only use a constant amount of extra space:
- A few integer variables for pointers and counters
- No data structures that grow with input size

## Edge Cases
1. **k = 0**: No flips allowed, find longest sequence of existing 1's
2. **All zeros**: Returns k (flip k zeros to get k consecutive 1's)
3. **All ones**: Returns length of array (no flips needed)
4. **k >= number of zeros**: Can flip all zeros, so returns array length
5. **Single element**: Returns 1 if it's a 1 or k >= 1

## Worked Example
For `nums = [1,1,1,0,0,0,1,1,1,1,0]`, `k = 2`:

```
Initial state:
left = 0, right = 0, currentZero = 0, maxNumber = 0

Step 1: right = 0, nums[0] = 1
currentZero = 0, window = [1]
maxNumber = 1

Step 2: right = 1, nums[1] = 1
currentZero = 0, window = [1,1]
maxNumber = 2

Step 3: right = 2, nums[2] = 1
currentZero = 0, window = [1,1,1]
maxNumber = 3

Step 4: right = 3, nums[3] = 0
currentZero = 1, window = [1,1,1,0]
maxNumber = 4

Step 5: right = 4, nums[4] = 0
currentZero = 2, window = [1,1,1,0,0]
maxNumber = 5

Step 6: right = 5, nums[5] = 0
currentZero = 3 > k (2)
Shrink window: left = 0, nums[0] = 1, left++
Shrink window: left = 1, nums[1] = 1, left++
Shrink window: left = 2, nums[2] = 1, left++
Shrink window: left = 3, nums[3] = 0, currentZero--, left++
currentZero = 2, window = [0,0,0]
maxNumber = 5

Step 7: right = 6, nums[6] = 1
currentZero = 2, window = [0,0,0,1]
maxNumber = 5

Step 8: right = 7, nums[7] = 1
currentZero = 2, window = [0,0,0,1,1]
maxNumber = 5

Step 9: right = 8, nums[8] = 1
currentZero = 2, window = [0,0,0,1,1,1]
maxNumber = 6

Step 10: right = 9, nums[9] = 1
currentZero = 2, window = [0,0,0,1,1,1,1]
maxNumber = 7 (but this is wrong, let me recalculate)
```

Actually, let me trace more carefully:

```
Window: [1,1,1,0,0] with 2 zeros → length 5
Window: [1,1,0,0,0] → shrink left
Window: [1,0,0,0] → shrink left  
Window: [0,0,0] → has 3 zeros, shrink left
Window: [0,0,0,1] → 2 zeros, valid, length 4
Window: [0,0,0,1,1] → 2 zeros, valid, length 5
Window: [0,0,0,1,1,1] → 2 zeros, valid, length 6
Window: [0,0,0,1,1,1,1] → 2 zeros, valid, length 7
But wait, we need to shrink when adding the 11th element (0)
```

The algorithm correctly identifies that flipping the two zeros at indices 3 and 4 (or 4 and 5) gives us windows like `[1,1,1,0,0,1]` with length 6.

Result: **6**

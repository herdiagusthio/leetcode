# Longest Subarray of 1's After Deleting One Element - Explanation

## Problem Summary
Given a binary array `nums`, you must delete exactly one element from it. Return the size of the longest non-empty subarray containing only 1's in the resulting array. Return 0 if there is no such subarray.

## Approach
This solution uses a **sliding window** technique to find the longest subarray with at most 1 zero:

1. **Expand Window**: Move the right pointer through the array, counting zeros
2. **Shrink Window**: When zeros exceed 1, move the left pointer until we have at most 1 zero
3. **Track Maximum**: Keep track of the largest window size, calculated as `right - left`

## Why This Works
The key insight is that:
- We **must** delete exactly one element
- Deleting a zero allows adjacent 1's to connect
- A window with at most 1 zero represents the optimal subarray after deleting that zero (or deleting a 1 if all are 1's)
- The length is `right - left` (not `right - left + 1`) because we're accounting for the mandatory deletion

**Special case**: If the array contains only 1's, we still must delete one, so the answer is `length - 1`.

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
1. **All ones** (e.g., [1,1,1]): Must delete one, so returns length - 1 = 2
2. **All zeros** (e.g., [0,0,0,0]): After deleting one zero, no subarray of 1's exists, returns 0
3. **Single element**: After deleting it, empty array remains, returns 0
4. **One zero in array**: Delete it to connect all 1's
5. **Multiple zeros**: Find longest subarray containing at most 1 zero

## Worked Example
For `nums = [0,1,1,1,0,1,1,0,1]`:

```
Goal: Find longest subarray with at most 1 zero

Initial state:
left = 0, right = 0, zeroCount = 0, result = 0

Step 1: right = 0, nums[0] = 0
zeroCount = 1, window = [0]
potentialLength = 0 - 0 = 0
result = 0

Step 2: right = 1, nums[1] = 1
zeroCount = 1, window = [0,1]
potentialLength = 1 - 0 = 1
result = 1

Step 3: right = 2, nums[2] = 1
zeroCount = 1, window = [0,1,1]
potentialLength = 2 - 0 = 2
result = 2

Step 4: right = 3, nums[3] = 1
zeroCount = 1, window = [0,1,1,1]
potentialLength = 3 - 0 = 3
result = 3

Step 5: right = 4, nums[4] = 0
zeroCount = 2 (exceeds 1!)
Shrink: left = 0, nums[0] = 0, zeroCount--, left++
Now: zeroCount = 1, window = [1,1,1,0]
potentialLength = 4 - 1 = 3
result = 3

Step 6: right = 5, nums[5] = 1
zeroCount = 1, window = [1,1,1,0,1]
potentialLength = 5 - 1 = 4
result = 4

Step 7: right = 6, nums[6] = 1
zeroCount = 1, window = [1,1,1,0,1,1]
potentialLength = 6 - 1 = 5
result = 5

Step 8: right = 7, nums[7] = 0
zeroCount = 2 (exceeds 1!)
Shrink: left = 1, nums[1] = 1, left++
Shrink: left = 2, nums[2] = 1, left++
Shrink: left = 3, nums[3] = 1, left++
Shrink: left = 4, nums[4] = 0, zeroCount--, left++
Now: zeroCount = 1, window = [1,1,0]
potentialLength = 7 - 5 = 2
result = 5

Step 9: right = 8, nums[8] = 1
zeroCount = 1, window = [1,1,0,1]
potentialLength = 8 - 5 = 3
result = 5
```

The longest window has length 5, which means after deleting the zero at position 4, we get the subarray [1,1,1,1,1].

Result: **5**

# Maximum Average Subarray I - Explanation

## Problem Summary
Given an integer array `nums` of length `n` and an integer `k`, find a contiguous subarray of length `k` that has the maximum average value and return this value.

## Approach
This solution uses a **sliding window technique**:

1. **Compute Initial Window**: Calculate the sum of the first `k` elements
2. **Track Maximum**: Set this as the initial maximum sum
3. **Slide the Window**: For each position from `k` to `n-1`:
   - Add the new element entering the window (right side)
   - Subtract the element leaving the window (left side)
   - Update maximum sum if current sum is larger
4. **Return Average**: Divide the maximum sum by `k` to get the maximum average

## Why This Works
The key insight is that maximum average equals maximum sum for fixed-length subarrays:
- Since we're dividing all sums by the same constant `k`, the ordering is preserved
- The subarray with the maximum sum will also have the maximum average
- Sliding window efficiently computes all k-length subarray sums in one pass

**Why sliding window is efficient:**
- Instead of recalculating the entire sum for each subarray (O(k) per window)
- We update the sum incrementally by adding one element and removing one (O(1) per window)
- This transforms an O(n*k) brute force into an O(n) solution

## Time Complexity
**O(n)** where n is the length of the array
- Compute initial window sum: O(k)
- Slide window (n-k) times with O(1) operations each: O(n-k)
- Total: O(k) + O(n-k) = O(n)

## Space Complexity
**O(1)** - constant space
- Only storing a few variables: current sum, maximum sum, loop counters
- No data structures that grow with input size

## Edge Cases
1. **k = 1**: Returns the maximum element in the array
2. **k = n**: Returns the average of the entire array
3. **All negative numbers**: Correctly finds the "maximum" (least negative) average
4. **All same values**: Returns that value as the average
5. **k = 0 or k > n**: Problem constraints guarantee 1 ≤ k ≤ n

## Worked Example
For `nums = [1, 12, -5, -6, 50, 3]`, `k = 4`:

```
Initial window [1, 12, -5, -6]:
  sum = 1 + 12 + (-5) + (-6) = 2
  maxSum = 2

Slide to [12, -5, -6, 50] (remove 1, add 50):
  sum = 2 - 1 + 50 = 51
  maxSum = 51 ✓ (new maximum)

Slide to [-5, -6, 50, 3] (remove 12, add 3):
  sum = 51 - 12 + 3 = 42
  maxSum = 51 (no change)

Maximum sum = 51
Maximum average = 51 / 4 = 12.75 ✓
```

## Visual Representation

For `nums = [1, 12, -5, -6, 50, 3]`, `k = 4`:

```
Index:    0   1    2   3   4   5
Value:   [1, 12, -5, -6, 50,  3]

Window 1: [1, 12, -5, -6]          sum = 2
          └──────────────┘

Window 2:    [12, -5, -6, 50]     sum = 51 ← maximum!
             └──────────────┘

Window 3:        [-5, -6, 50, 3]  sum = 42
                 └──────────────┘

Sliding window process:
Step 1: Calculate sum of first k=4 elements
Step 2: Slide right by removing leftmost, adding rightmost
Step 3: Track maximum sum throughout

Result: 51/4 = 12.75
```

## Step-by-Step Visualization

```
nums = [1, 12, -5, -6, 50, 3], k = 4

┌─────────────────────────────┐
│ Window 1: indices 0-3       │
│ [1, 12, -5, -6] = 2         │
│ maxSum = 2                  │
└─────────────────────────────┘
           ↓ slide right
┌─────────────────────────────┐
│ Window 2: indices 1-4       │
│ [12, -5, -6, 50] = 51       │
│ maxSum = 51 (updated!)      │
└─────────────────────────────┘
           ↓ slide right
┌─────────────────────────────┐
│ Window 3: indices 2-5       │
│ [-5, -6, 50, 3] = 42        │
│ maxSum = 51 (unchanged)     │
└─────────────────────────────┘

Final answer: 51 ÷ 4 = 12.75
```

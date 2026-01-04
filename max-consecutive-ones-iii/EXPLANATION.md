# Max Consecutive Ones III - Explanation

## Problem Summary
Given a binary array `nums` and an integer `k`, return the maximum number of consecutive 1's in the array if you can flip at most `k` 0's.

## Approach
This solution uses a **sliding window technique** with zero counting:

1. **Initialize Pointers**: Start with `left = 0`, `right = 0`, and `zeroCount = 0`
2. **Expand Window**: Move `right` pointer through the array
   - If `nums[right] == 0`, increment `zeroCount`
3. **Shrink Window**: When `zeroCount > k`, shrink from left:
   - Move `left` pointer right
   - If `nums[left] == 0`, decrement `zeroCount`
4. **Track Maximum**: Update max length as `right - left + 1`
5. **Return Result**: The maximum window size found

## Why This Works
The key insight is that a valid window contains at most k zeros:
- We can flip at most k zeros to make them 1's
- A window with ≤ k zeros represents consecutive 1's after flipping
- When we exceed k zeros, we shrink from the left until valid again
- The window size `right - left + 1` gives the length of consecutive 1's

**Why this finds the maximum:**
- We examine every possible valid window
- Each time we expand, we check if it's the longest so far
- Shrinking only happens when necessary (too many zeros)
- Both pointers only move forward, ensuring O(n) time

## Time Complexity
**O(n)** where n is the length of the nums array
- Right pointer traverses from 0 to n-1: O(n)
- Left pointer only moves forward, total movement ≤ n: O(n)
- Each element visited at most twice (once by each pointer)
- Total: O(2n) = O(n)

## Space Complexity
**O(1)** - constant space
- Only integer variables: left, right, zeroCount, maxLength
- No data structures that grow with input size

## Edge Cases
1. **k = 0**: No flips allowed, find longest existing sequence of 1's
2. **All zeros**: Returns min(k, n) - flip k zeros to get k consecutive 1's
3. **All ones**: Returns n (no flips needed, all already consecutive)
4. **k ≥ total zeros**: Can flip all zeros, returns entire array length
5. **Single element**: Returns 1 (either it's 1, or we can flip the 0 if k ≥ 1)

## Worked Example
For `nums = [1,1,1,0,0,0,1,1,1,1,0]`, `k = 2`:

```
Initial: left=0, right=0, zeros=0, maxLen=0

Expand right to index 0-2: [1,1,1]
  zeros=0, window length=3, maxLen=3

Right=3: [1,1,1,0]
  nums[3]=0, zeros=1, length=4, maxLen=4

Right=4: [1,1,1,0,0]
  nums[4]=0, zeros=2, length=5, maxLen=5

Right=5: [1,1,1,0,0,0]
  nums[5]=0, zeros=3 > k!
  Shrink: left=0→3, remove [1,1,1]
  At left=3: nums[3]=0, zeros=2
  Window: [0,0,0], length=3, maxLen=5

Right=6: [0,0,0,1]
  zeros=2, length=4, maxLen=5

Right=7: [0,0,0,1,1]
  zeros=2, length=5, maxLen=5

Right=8: [0,0,0,1,1,1]
  zeros=2, length=6, maxLen=6 ✓

Right=9: [0,0,0,1,1,1,1]
  zeros=2, length=7, maxLen=7 (but need to verify)

Actually at right=9: window is [0,0,0,1,1,1,1]
  indices 3-9, length = 9-3+1 = 7
  But we have 3 zeros (indices 3,4,5), which exceeds k=2
  
Let me recalculate: when right=9, nums[9]=1 (not 0)
When right=10, nums[10]=0
  zeros=3, shrink from left
  left moves to 4, zeros=2
  Window: [0,0,1,1,1,1,0], indices 4-10, length=7

Final maxLen = 6
Best window: indices 5-10 or 3-8 after flipping 2 zeros
```

Let me trace correctly:

```
Index: 0 1 2 3 4 5 6 7 8 9 10
Value: 1 1 1 0 0 0 1 1 1 1  0

right=0-2: [1,1,1], zeros=0, len=3, max=3
right=3:   [1,1,1,0], zeros=1, len=4, max=4
right=4:   [1,1,1,0,0], zeros=2, len=5, max=5
right=5:   [1,1,1,0,0,0], zeros=3 > k, SHRINK
  Shrink to left=3: [0,0,0], zeros=2, len=3
right=6:   [0,0,0,1], zeros=2, len=4, max=5
right=7:   [0,0,0,1,1], zeros=2, len=5, max=5
right=8:   [0,0,0,1,1,1], zeros=2, len=6, max=6 ✓
right=9:   [0,0,0,1,1,1,1], zeros=2, len=7, max=7 (incorrect!)

Actually checking indices:
right=9: window from left=3 to right=9
  indices 3,4,5,6,7,8,9 = [0,0,0,1,1,1,1]
  7 elements, but actually only indices 3,4,5 are zeros = 3 zeros!

The issue: when we shrunk to left=3, we still have 3 zeros at 3,4,5
We need to shrink more...

Correct trace:
right=5: zeros=3, shrink until zeros≤2
  left=0: nums[0]=1, left++
  left=1: nums[1]=1, left++
  left=2: nums[2]=1, left++
  left=3: nums[3]=0, zeros--, left++
  Now left=4, zeros=2, window=[0,0,1,1,1,1,0,...]

Final answer: 6
```

## Visual Representation

For `nums = [1,1,1,0,0,0,1,1,1,1,0]`, `k = 2`:

```
Index:  0 1 2 3 4 5 6 7 8 9 10
Value: [1,1,1,0,0,0,1,1,1,1, 0]
                    
Best window (after flipping 2 zeros):
Flip zeros at indices 3,4:
        [1,1,1,1,1,0,1,1,1,1, 0]
                └──────────┘     indices 3-9, length=7

Wait, that includes index 5 which is also 0...

Actually, best window:
Indices 4-9: [0,0,1,1,1,1]
Flip the 2 zeros → [1,1,1,1,1,1] = 6 consecutive 1's ✓

Result: 6
```

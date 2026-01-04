# Move Zeroes - Explanation

## Problem Summary
Given an integer array `nums`, move all zeros to the end while maintaining the relative order of non-zero elements. The operation must be done in-place without copying the array.

## Approach
This solution uses a **two-pointer swap technique**:

1. **Initialize Pointer**: Start with `pointer = 0` to track where the next non-zero should go
2. **Scan Array**: Iterate through each element with index `i`
3. **Swap Non-Zeros**: When `nums[i] != 0`:
   - Swap `nums[pointer]` with `nums[i]`
   - Increment `pointer`
4. **Result**: Non-zeros at front in original order, zeros at end

## Why This Works
The key insight is using a single pointer to partition the array:
- **Pointer tracks boundary**: Elements before `pointer` are all non-zero
- **Swap maintains order**: When we find a non-zero, we swap it to the next available position
- **Self-swaps are safe**: When `pointer == i`, we swap an element with itself (no-op)
- **Zeros naturally accumulate**: As non-zeros move forward, zeros shift toward the end
- **Single pass suffices**: We don't need to explicitly move zeros; they end up at the end automatically

**Why relative order is preserved**:
- Non-zero elements are encountered and moved in left-to-right order
- The pointer only advances when we place a non-zero
- Earlier non-zeros always get placed before later ones

## Time Complexity
**O(n)** where n is the length of the nums array
- Single pass through the array
- Each element is examined exactly once
- Each swap operation is O(1)
- No nested loops

## Space Complexity
**O(1)** - constant space
- Only two integer variables (`pointer`, `i`)
- In-place modification of the input array
- No additional data structures

## Edge Cases
1. **All zeros**: [0,0,0] → [0,0,0] (no change needed)
2. **No zeros**: [1,2,3] → [1,2,3] (swaps with self)
3. **Single element**: [0] or [1] → no change
4. **Zeros at end**: [1,2,0,0] → [1,2,0,0] (already correct)
5. **Zeros at start**: [0,0,1,2] → [1,2,0,0]
6. **Alternating**: [0,1,0,3] → [1,3,0,0]

## Worked Example
For `nums = [0,1,0,3,12]`:

```
Initial: nums = [0,1,0,3,12], pointer = 0

i=0, nums[0]=0:
  0 is zero, skip
  pointer = 0

i=1, nums[1]=1:
  1 is non-zero!
  Swap nums[0] and nums[1]: [1,0,0,3,12]
  pointer = 1

i=2, nums[2]=0:
  0 is zero, skip
  pointer = 1

i=3, nums[3]=3:
  3 is non-zero!
  Swap nums[1] and nums[3]: [1,3,0,0,12]
  pointer = 2

i=4, nums[4]=12:
  12 is non-zero!
  Swap nums[2] and nums[4]: [1,3,12,0,0]
  pointer = 3

Result: [1,3,12,0,0] ✓
```

## Visual Representation
```
Initial state:
[0, 1, 0, 3, 12]
 ^              pointer=0

After i=1 (swap 0 and 1):
[1, 0, 0, 3, 12]
    ^           pointer=1

After i=3 (swap 0 and 3):
[1, 3, 0, 0, 12]
       ^        pointer=2

After i=4 (swap 0 and 12):
[1, 3, 12, 0, 0]
          ^     pointer=3

Partitioning visualization:
[non-zeros | zeros]
[1, 3, 12 | 0, 0]
          ^
       pointer (boundary)
```

## Example: No Zeros
For `nums = [1,2,3]`:

```
i=0: Swap nums[0] with nums[0]: [1,2,3], pointer=1
i=1: Swap nums[1] with nums[1]: [1,2,3], pointer=2
i=2: Swap nums[2] with nums[2]: [1,2,3], pointer=3

Result: [1,2,3] (unchanged, all self-swaps)
```

## Example: All Zeros
For `nums = [0,0,0]`:

```
i=0: nums[0]=0, skip, pointer=0
i=1: nums[1]=0, skip, pointer=0
i=2: nums[2]=0, skip, pointer=0

Result: [0,0,0] (unchanged, pointer never moved)
```

## Partition Invariant
```
At any point during iteration:
┌─────────────┬─────────────┬──────────────┐
│  Non-zeros  │   Zeros     │  Unprocessed │
│  (ordered)  │             │              │
└─────────────┴─────────────┴──────────────┘
0          pointer          i            n-1

Invariant maintained:
- Elements [0, pointer) are all non-zero
- Elements [pointer, i) are all zero
- Elements [i, n) are not yet processed
```

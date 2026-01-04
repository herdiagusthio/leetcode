# Can Place Flowers - Explanation

## Problem Summary
Given a flowerbed array of 0s (empty) and 1s (flower planted), determine if you can plant `n` additional flowers without any two flowers being adjacent to each other.

## Approach
This solution uses a **greedy single-pass algorithm**:

1. **Iterate Through Flowerbed**: Check each position from left to right
2. **Check Planting Conditions**: For each empty spot, verify:
   - Current position is 0 (empty)
   - Left neighbor is 0 or out of bounds
   - Right neighbor is 0 or out of bounds
3. **Plant Greedily**: If conditions met:
   - Increment planted counter
   - Skip next position (can't plant adjacent)
4. **Early Exit**: Return true if we've planted n flowers
5. **Return Result**: Check if n flowers were planted

## Why This Works
The key insight is that the greedy approach is optimal:
- **If we can plant a flower at a position, we should** - delaying never helps
- **Planting early maximizes future opportunities** - each planted flower creates a constraint, but planting left-to-right ensures we don't miss opportunities
- **Skipping the next position** after planting prevents checking an impossible location
- **No backtracking needed** - if a position is valid, using it immediately is always optimal or equal to any other strategy

## Time Complexity
**O(n)** where n is the length of the flowerbed
- Single pass through the array
- Each position is checked at most once (we skip after planting)
- Constant time operations for each position
- Early exit optimization when n flowers are planted

## Space Complexity
**O(1)** - constant space
- Only a few integer variables (`count`, `i`, `length`)
- No data structures that grow with input size
- The flowerbed is not modified (or modification doesn't count toward space complexity)

## Edge Cases
1. **n = 0**: Return true immediately (no flowers to plant)
2. **Single empty spot** ([0]): Can plant 1 flower
3. **All empty**: Can plant ⌊(length + 1) / 2⌋ flowers
4. **All filled**: Can plant 0 flowers
5. **Adjacent flowers**: [1,0,0,0,1] → can plant 1 in middle
6. **Ends of array**: Special handling for first and last positions

## Worked Example
For `flowerbed = [1,0,0,0,1]`, `n = 1`:

```
Initial: count = 0, need to plant 1 flower

Index 0: flowerbed[0] = 1 (already has flower)
→ Skip

Index 1: flowerbed[1] = 0
Check left (index 0): flowerbed[0] = 1 (has flower)
→ Cannot plant (adjacent to existing flower)

Index 2: flowerbed[2] = 0
Check left (index 1): flowerbed[1] = 0 ✓
Check right (index 3): flowerbed[3] = 0 ✓
→ Can plant! count = 1
→ Skip to index 4 (can't plant adjacent)

Index 4: flowerbed[4] = 1 (already has flower)
→ Skip

count (1) >= n (1) ✓

Result: true
```

## Visual Representation
```
Before: [1, 0, 0, 0, 1]  n = 1
         ^        x  ^
      flower   can   flower
               plant
               here!

After:  [1, 0, 1, 0, 1]
         ^     ^     ^
      existing NEW existing

Can plant 1 flower ✓
```

## Example: Cannot Plant Enough
For `flowerbed = [1,0,0,0,1]`, `n = 2`:

```
Index 0: Has flower, skip
Index 1: Adjacent to flower at index 0, skip
Index 2: Can plant! count = 1, skip to 4
Index 4: Has flower, skip

count (1) < n (2)

Result: false (can only plant 1, but need 2)
```

## Example: All Empty
For `flowerbed = [0,0,0,0,0]`, `n = 3`:

```
Index 0: Empty, no left neighbor, right is empty
→ Plant! count = 1, skip to 2

Index 2: Empty, left is empty, right is empty
→ Plant! count = 2, skip to 4

Index 4: Empty, left is empty, no right neighbor
→ Plant! count = 3, done!

Result: true
```

Visual:
```
[0, 0, 0, 0, 0]
 ↓     ↓     ↓
[1, 0, 1, 0, 1]  ← Can plant 3 flowers
```

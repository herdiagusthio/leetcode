# Container With Most Water - Explanation

## Problem Summary
Given an array `height` where each element represents the height of a vertical line at that position, find two lines that together with the x-axis form a container that can hold the most water. Return the maximum area.

## Approach
This solution uses a **two-pointer technique**:

1. **Initialize Pointers**: Place left pointer at start (0) and right pointer at end (n-1)
2. **Calculate Area**: For current pair, area = `width × min(height[left], height[right])`
   - Width is the distance between pointers: `right - left`
   - Height is limited by the shorter line: `min(height[left], height[right])`
3. **Update Maximum**: Track the largest area found
4. **Move Pointer**: Move the pointer pointing to the shorter line inward
   - If `height[left] < height[right]`, increment left
   - Otherwise, decrement right
5. **Repeat**: Continue until pointers meet

## Why This Works
The key insight is understanding why moving the shorter line's pointer is optimal:
- **Area formula**: `area = width × min_height`
- **Width always decreases** as pointers move inward
- **To potentially increase area**, we need to increase the minimum height
- **Moving the taller line's pointer** guarantees the area will decrease:
  - Width decreases
  - Min height stays same or decreases (limited by the shorter line we didn't move)
- **Moving the shorter line's pointer** gives a chance to increase area:
  - Width decreases (unavoidable)
  - But min height might increase if we find a taller line
- **We never miss the optimal solution** because:
  - We start with maximum width
  - We only discard configurations that cannot be better than what we've seen

## Time Complexity
**O(n)** where n is the length of the height array
- Each pointer moves at most n positions
- Left pointer moves from 0 toward right
- Right pointer moves from n-1 toward left
- Combined movement: at most n steps
- Each iteration: O(1) operations (min, multiplication, comparison)

## Space Complexity
**O(1)** - constant space
- Only a few integer variables (`i`, `j`, `result`, `area`)
- No data structures that grow with input size
- No recursion stack

## Edge Cases
1. **Two elements**: Return area between them
2. **All same height**: Area is `height × (n-1)`
3. **Decreasing heights**: Best area likely near the start
4. **Increasing heights**: Best area likely near the end
5. **One very tall line**: Area limited by other shorter lines
6. **Zero heights**: Area is 0

## Worked Example
For `height = [1,8,6,2,5,4,8,3,7]`:

```
Indices:  0 1 2 3 4 5 6 7 8
Heights:  1 8 6 2 5 4 8 3 7

Initial: left=0, right=8, maxArea=0

Step 1: left=0 (h=1), right=8 (h=7)
area = (8-0) × min(1,7) = 8 × 1 = 8
maxArea = 8
height[0]=1 < height[8]=7 → move left++

Step 2: left=1 (h=8), right=8 (h=7)
area = (8-1) × min(8,7) = 7 × 7 = 49
maxArea = 49
height[1]=8 > height[8]=7 → move right--

Step 3: left=1 (h=8), right=7 (h=3)
area = (7-1) × min(8,3) = 6 × 3 = 18
maxArea = 49 (no change)
height[1]=8 > height[7]=3 → move right--

Step 4: left=1 (h=8), right=6 (h=8)
area = (6-1) × min(8,8) = 5 × 8 = 40
maxArea = 49 (no change)
height[1]=8 = height[6]=8 → move left++ (tie-breaker)

Step 5: left=2 (h=6), right=6 (h=8)
area = (6-2) × min(6,8) = 4 × 6 = 24
maxArea = 49 (no change)
height[2]=6 < height[6]=8 → move left++

Step 6: left=3 (h=2), right=6 (h=8)
area = (6-3) × min(2,8) = 3 × 2 = 6
maxArea = 49 (no change)
height[3]=2 < height[6]=8 → move left++

Step 7: left=4 (h=5), right=6 (h=8)
area = (6-4) × min(5,8) = 2 × 5 = 10
maxArea = 49 (no change)
height[4]=5 < height[6]=8 → move left++

Step 8: left=5 (h=4), right=6 (h=8)
area = (6-5) × min(4,8) = 1 × 4 = 4
maxArea = 49 (no change)
left=5, right=6 → left++ would make left >= right, stop

Result: 49
```

## Visual Representation
```
Height: [1, 8, 6, 2, 5, 4, 8, 3, 7]
Index:   0  1  2  3  4  5  6  7  8

         |              |
         |              |     |
         |  |           |     |
         |  |     |     |     |
         |  |     |  |  |     |
         |  |     |  |  |  |  |
         |  |  |  |  |  |  |  |
      |  |  |  |  |  |  |  |  |
      ----------------------------- (x-axis)
      0  1  2  3  4  5  6  7  8

Best container: between index 1 and index 8
      *--[  water area  ]--*
      1                    8
Height: min(8,7) = 7
Width: 8 - 1 = 7
Area: 7 × 7 = 49
```

## Why Other Pairs Are Worse
```
Pair (0,8): width=8, height=min(1,7)=1, area=8
            ↓ Too short on left side

Pair (1,6): width=5, height=min(8,8)=8, area=40
            ↓ Width too narrow

Pair (1,8): width=7, height=min(8,7)=7, area=49 ✓ BEST
            ↓ Good balance of width and height
```

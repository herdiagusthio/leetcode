# Product of Array Except Self - Explanation

## Problem Summary
Given an integer array `nums`, return an array `answer` where `answer[i]` equals the product of all elements in `nums` except `nums[i]`. The solution must run in O(n) time without using division.

## Approach
This solution uses a **two-pass prefix-suffix product technique**:

1. **First Pass (Left to Right)**: Build prefix products
   - For each index `i`, store the product of all elements to its left
   - Use running variable to accumulate products
2. **Second Pass (Right to Left)**: Multiply by suffix products
   - For each index `i`, multiply by the product of all elements to its right
   - Use running variable to accumulate products
3. **Result**: Each `answer[i]` = (product of left elements) × (product of right elements)

## Why This Works
The key insight is that the product excluding `nums[i]` equals left product × right product:
- **Product except self** = (all elements before i) × (all elements after i)
- **Two passes compute both sides**: 
  - First pass: accumulate products from left
  - Second pass: accumulate products from right
- **No division needed**: We never compute total product or divide
- **Handles zeros naturally**: 
  - One zero: only that position has non-zero answer
  - Multiple zeros: all answers are zero
- **Space optimization**: Use output array to store intermediate results

**Why we don't need division**:
```
Traditional approach: total / nums[i]
❌ Fails with zeros
❌ Division operation disallowed

Our approach: left_product × right_product
✓ Works with zeros
✓ No division needed
```

## Time Complexity
**O(n)** where n is the length of the nums array
- First pass: O(n) to compute prefix products
- Second pass: O(n) to compute and apply suffix products
- Total: O(n) + O(n) = O(n)
- Each element is visited exactly twice
- All operations per element are O(1)

## Space Complexity
**O(1)** extra space (excluding output array)
- Output array `answer` of size n is required by problem
- Only two additional variables: `leftProd` and `rightProd`
- No additional arrays or data structures
- The output array is used to store intermediate results (prefix products)

## Edge Cases
1. **Two elements**: [2,3] → [3,2]
2. **Contains zero**: [1,2,0,4] → [0,0,8,0]
3. **Multiple zeros**: [0,0,1] → [0,0,0]
4. **All ones**: [1,1,1] → [1,1,1]
5. **Negative numbers**: [-1,2,-3] → [-6,-3,2]
6. **Large numbers**: Products fit in 32-bit integers (per problem guarantee)

## Worked Example
For `nums = [1,2,3,4]`:

```
Step 1: Build prefix products (left to right)
leftProd = 1 (identity for multiplication)

i=0: answer[0] = leftProd = 1 (no elements to left)
     leftProd = 1 * nums[0] = 1 * 1 = 1

i=1: answer[1] = leftProd = 1 (product of [1])
     leftProd = 1 * nums[1] = 1 * 2 = 2

i=2: answer[2] = leftProd = 2 (product of [1,2])
     leftProd = 2 * nums[2] = 2 * 3 = 6

i=3: answer[3] = leftProd = 6 (product of [1,2,3])
     leftProd = 6 * nums[3] = 6 * 4 = 24

After first pass: answer = [1, 1, 2, 6]

Step 2: Multiply by suffix products (right to left)
rightProd = 1 (identity for multiplication)

i=3: answer[3] *= rightProd = 6 * 1 = 6 (no elements to right)
     rightProd = 1 * nums[3] = 1 * 4 = 4

i=2: answer[2] *= rightProd = 2 * 4 = 8 (product of [4])
     rightProd = 4 * nums[2] = 4 * 3 = 12

i=1: answer[1] *= rightProd = 1 * 12 = 12 (product of [3,4])
     rightProd = 12 * nums[1] = 12 * 2 = 24

i=0: answer[0] *= rightProd = 1 * 24 = 24 (product of [2,3,4])
     rightProd = 24 * nums[0] = 24 * 1 = 24

Final answer: [24, 12, 8, 6]
```

Verification:
- answer[0] = 2×3×4 = 24 ✓
- answer[1] = 1×3×4 = 12 ✓
- answer[2] = 1×2×4 = 8 ✓
- answer[3] = 1×2×3 = 6 ✓

## Visual Representation
```
nums:    [1,  2,  3,  4]
          ↓   ↓   ↓   ↓

For each index, we need product of others:
Index 0: [ ,  2,  3,  4] → 2×3×4 = 24
Index 1: [1,  ,  3,  4] → 1×3×4 = 12
Index 2: [1,  2,  ,  4] → 1×2×4 = 8
Index 3: [1,  2,  3,  ] → 1×2×3 = 6

Two-pass visualization:
Pass 1 (Prefix - Left to Right):
answer[0] = 1         (nothing to left)
answer[1] = 1         (1)
answer[2] = 1×2       (1,2)
answer[3] = 1×2×3     (1,2,3)

Pass 2 (Suffix - Right to Left):
answer[3] *= 1        (nothing to right)
answer[2] *= 4        (4)
answer[1] *= 3×4      (3,4)
answer[0] *= 2×3×4    (2,3,4)

Result: [24, 12, 8, 6]
```

## Example with Zero
For `nums = [1,2,0,4]`:

```
Pass 1 (prefix):
answer = [1, 1, 2, 0]

Pass 2 (suffix):
i=3: answer[3] = 0 * 1 = 0
     rightProd = 4
i=2: answer[2] = 2 * 4 = 8
     rightProd = 0 (contains zero!)
i=1: answer[1] = 1 * 0 = 0
     rightProd = 0
i=0: answer[0] = 1 * 0 = 0

Result: [0, 0, 8, 0]
```

Only index 2 (the zero's position) has non-zero answer ✓

## Pattern Visualization
```
For nums = [a, b, c, d]:

answer[0] = [    b × c × d]
answer[1] = [a ×     c × d]
answer[2] = [a × b ×     d]
answer[3] = [a × b × c    ]

Left products:  [1, a, a×b, a×b×c]
Right products: [b×c×d, c×d, d, 1]

Multiply corresponding elements:
answer[i] = leftProduct[i] × rightProduct[i]
```

# Max Number of K-Sum Pairs - Explanation

## Problem Summary
Given an integer array `nums` and an integer `k`, you can perform operations where you pick two numbers whose sum equals `k` and remove them from the array. Return the maximum number of such operations you can perform.

## Approach
This solution uses a **hash map with greedy complement matching**:

1. **Initialize Hash Map**: Create a map to track the frequency of numbers we've seen
2. **Process Each Number**: For each element `v` in the array:
   - Calculate complement: `diff = k - v`
   - Check if complement exists in the map
3. **Form Pairs**: If complement is available (count > 0):
   - Increment result counter
   - Decrement the complement's count
4. **Store for Later**: If no complement available:
   - Increment count for current number in the map
5. **Return Result**: The total number of pairs formed

## Why This Works
The key insight is that we can greedily match each number with its complement:
- When we see a number, we immediately check if its complement was seen before
- If yes, we form a pair right away (increment result, decrement complement count)
- If no, we store it for potential future matching
- This greedy approach works because using a complement immediately cannot reduce the total number of pairs we can form later
- Each number is used at most once (we decrement counts when forming pairs)

**Why greedy is optimal:**
- Any valid pairing can be reordered to match our greedy order
- Using an available complement now vs later doesn't affect other potential pairs
- The map ensures we never use a number more times than it appears

## Time Complexity
**O(n)** where n is the length of the array
- Single pass through the array: O(n)
- Hash map operations (lookup, insert, update): O(1) average case
- Total: O(n)

## Space Complexity
**O(u)** where u is the number of distinct values in nums
- Hash map stores at most one entry per unique number
- In worst case (all numbers distinct): O(n)
- In best case (all numbers same): O(1)

## Edge Cases
1. **Self-complement (k/2)**: Works correctly when `v == diff` (e.g., k=4, v=2) because we track counts and require two instances
2. **Negative numbers**: Hash map approach handles negative values naturally
3. **Zero**: Works correctly with k=0 or nums containing 0
4. **Odd pairs**: If one number can't be paired, it remains in the map unused
5. **Empty array or single element**: Returns 0 (no pairs possible)
6. **Duplicates**: Correctly handles multiple occurrences via frequency counting

## Worked Example
For `nums = [1,2,3,4]`, `k = 5`:

```
Initial: counts = {}, result = 0

Process v=1:
  diff = 5-1 = 4
  counts[4] = 0 (not available)
  counts[1]++ → counts = {1:1}

Process v=2:
  diff = 5-2 = 3
  counts[3] = 0 (not available)
  counts[2]++ → counts = {1:1, 2:1}

Process v=3:
  diff = 5-3 = 2
  counts[2] = 1 ✓ (available!)
  result++ → result = 1
  counts[2]-- → counts = {1:1, 2:0}
  Pair formed: (2,3)

Process v=4:
  diff = 5-4 = 1
  counts[1] = 1 ✓ (available!)
  result++ → result = 2
  counts[1]-- → counts = {1:0, 2:0}
  Pair formed: (1,4)

Final result: 2 ✓
Pairs: (2,3) and (1,4)
```

## Complex Example
For `nums = [3,1,3,4,3]`, `k = 6`:

```
Initial: counts = {}, result = 0

Process v=3:
  diff = 3, no complement
  counts = {3:1}

Process v=1:
  diff = 5, no complement
  counts = {3:1, 1:1}

Process v=3:
  diff = 3, counts[3] = 1 ✓
  result = 1, counts[3]-- → counts = {3:0, 1:1}
  Pair: (3,3)

Process v=4:
  diff = 2, no complement
  counts = {3:0, 1:1, 4:1}

Process v=3:
  diff = 3, counts[3] = 0 (not available)
  counts = {3:1, 1:1, 4:1}

Final result: 1 ✓
Pair: (3,3) - one element remains unpaired
```

## Visual Representation

For `nums = [1,2,3,4]`, `k = 5`:
```
Array:  [1,  2,  3,  4]
         └────────┘  └─┘
         pair (1,4)   pair (2,3) formed when we see 3

Step-by-step:
1 → store in map
2 → store in map
3 → finds 2 in map → pair (2,3) ✓
4 → finds 1 in map → pair (1,4) ✓

Result: 2 pairs
```

For `nums = [3,1,3,4,3]`, `k = 6`:
```
Array:  [3,  1,  3,  4,  3]
         └────┘              
         pair (3,3)
         
Only one valid pair possible: (3,3)
Remaining: 1, 4, 3 (cannot form pairs with k=6)
```

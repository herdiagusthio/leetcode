# Contains Duplicate - Explanation

## Problem Summary
Given an integer array `nums`, return `true` if any value appears at least twice in the array, otherwise return `false`.

## Approach
This solution uses a **hash set tracking** approach:

1. **Initialize Hash Set**: Create an empty set to track seen numbers
2. **Iterate Through Array**: For each number:
   - Check if it's already in the set
   - If yes, return true immediately (found duplicate)
   - If no, add it to the set
3. **No Duplicates Found**: If we finish the loop, return false

## Why This Works
The key insight is using a hash set for O(1) lookups:
- **Hash set stores unique values** - attempting to add a duplicate tells us we've seen it before
- **Early return optimization** - we stop as soon as we find the first duplicate
- **Trade space for time** - we use O(n) space to achieve O(n) time instead of O(n²) brute force
- **No need to count** - we only care if something appears more than once, not how many times

## Time Complexity
**O(n)** where n is the length of the nums array
- In the worst case (no duplicates), we iterate through all n elements
- Each hash set insertion: O(1) average case
- Each hash set lookup: O(1) average case
- Total: O(n) average case
- Worst case with hash collisions: O(n²), but very rare with good hash function

## Space Complexity
**O(n)** in the worst case
- Hash set stores up to n unique elements (when no duplicates)
- Best case: O(1) if duplicate found immediately
- Average case: O(k) where k is number of unique elements before finding duplicate

## Edge Cases
1. **Empty array**: No elements means no duplicates, return false
2. **Single element**: Cannot have duplicate, return false
3. **Two identical elements**: [1, 1] → return true
4. **All unique**: [1, 2, 3, 4, 5] → return false
5. **Duplicate at end**: Must check all elements before finding it
6. **Negative numbers**: Works with any integer values
7. **Zero duplicates**: [0, 0] → return true

## Worked Example
For `nums = [1,2,3,1]`:

```
Initial: seen = {}

Index 0, num = 1:
1 not in seen
seen = {1}

Index 1, num = 2:
2 not in seen
seen = {1, 2}

Index 2, num = 3:
3 not in seen
seen = {1, 2, 3}

Index 3, num = 1:
1 IS in seen ✓
Return true immediately
```

## Visual Representation
```
nums: [1, 2, 3, 1]
       ^  ^  ^  ^
       |  |  |  |
seen:  1  1,2 1,2,3
                  └── 1 already exists! DUPLICATE FOUND

Timeline:
Step 1: Check 1 → Add to set
Step 2: Check 2 → Add to set
Step 3: Check 3 → Add to set
Step 4: Check 1 → Already in set → Return true
```

## Example: No Duplicates
For `nums = [1,2,3,4]`:

```
Index 0: Add 1, seen = {1}
Index 1: Add 2, seen = {1,2}
Index 2: Add 3, seen = {1,2,3}
Index 3: Add 4, seen = {1,2,3,4}

Loop completed, no duplicates found
Result: false
```

## Example: Immediate Duplicate
For `nums = [1,1]`:

```
Index 0: Add 1, seen = {1}
Index 1: Check 1 → in seen → Return true

Result: true (early exit after 2 checks)
```
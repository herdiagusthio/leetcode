# Unique Number of Occurrences - Explanation

## Problem Summary
Given an array of integers `arr`, return true if the number of occurrences of each value in the array is unique, or false otherwise.

## Approach
This solution uses a **two-pass hash map strategy**:

1. **Count Occurrences**: Create a map to count how many times each value appears
2. **Check Uniqueness**: Use a second map (or set) to track which occurrence counts we've seen
3. **Early Exit**: If we encounter a duplicate occurrence count, return false immediately
4. **Return True**: If all occurrence counts are unique, return true

## Why This Works
The key insight is separating the problem into two distinct phases:
- **Phase 1**: Count frequencies - this tells us how many times each value appears
- **Phase 2**: Verify uniqueness - check if any two values have the same occurrence count

Using a boolean map (`map[int]bool`) for tracking is more efficient than counting duplicates, as we only need to know if we've seen an occurrence count before, not how many times.

## Time Complexity
**O(n)** where n is the length of the array
- First loop: O(n) to count all occurrences
- Second loop: O(k) where k is number of unique values ≤ n
- Overall: O(n) + O(k) = O(n)
- Map operations (insert, lookup) are O(1) average case

## Space Complexity
**O(n)** in worst case
- First map: O(k) where k = number of unique values ≤ n
- Second map: O(k) for tracking unique counts
- Overall: O(k) + O(k) = O(2k) = O(k) ≤ O(n)
- Best case: O(1) when all elements are the same

## Edge Cases
1. **Single element**: Always returns true (one value with one occurrence)
2. **All same values**: Returns true (only one occurrence count)
3. **All unique values**: Returns false (all have count of 1)
4. **Two elements**: Returns false (both have count of 1)
5. **Negative numbers**: Algorithm works with any integer values
6. **Zeros**: Treated like any other value

## Worked Example
For `arr = [1,2,2,1,1,3]`:

```
Step 1: Count occurrences
Process each element:
1 → occurrences[1] = 1
2 → occurrences[2] = 1
2 → occurrences[2] = 2
1 → occurrences[1] = 2
1 → occurrences[1] = 3
3 → occurrences[3] = 1

occurrences map: {1: 3, 2: 2, 3: 1}

Step 2: Check uniqueness of counts
Check count 3:
  seen[3] = false → mark seen[3] = true

Check count 2:
  seen[2] = false → mark seen[2] = true

Check count 1:
  seen[1] = false → mark seen[1] = true

All counts are unique!
Result: true
```

## Example: Not Unique
For `arr = [1,2]`:

```
Step 1: Count occurrences
occurrences map: {1: 1, 2: 1}

Step 2: Check uniqueness of counts
Check count 1 (for value 1):
  seen[1] = false → mark seen[1] = true

Check count 1 (for value 2):
  seen[1] = true → DUPLICATE FOUND!
  Return false immediately

Result: false
```

## Example: Complex Case
For `arr = [-3,0,1,-3,1,1,1,-3,10,0]`:

```
Step 1: Count occurrences
-3 appears 3 times
0 appears 2 times
1 appears 4 times
10 appears 1 time

occurrences map: {-3: 3, 0: 2, 1: 4, 10: 1}

Step 2: Check uniqueness of counts
Counts to check: [3, 2, 4, 1]
- 3: unique ✓
- 2: unique ✓
- 4: unique ✓
- 1: unique ✓

All occurrence counts are different!
Result: true
```

## Optimization Details

**Why `map[int]bool` is better than `map[int]int`:**

1. **Memory**: bool (1 byte) vs int (8 bytes) = 87.5% less memory per entry
2. **Semantics**: We only need "exists" vs "doesn't exist", not a count
3. **Clarity**: Code intent is clearer - checking membership, not counting
4. **Performance**: Slightly faster since we don't increment unnecessarily

**Original approach (less optimal):**
```go
trackerSet[v]++
if trackerSet[v] > 1 {
    return false
}
```

**Optimized approach:**
```go
if seen[count] {
    return false
}
seen[count] = true
```

The optimized version is cleaner, more efficient, and idiomatic Go!

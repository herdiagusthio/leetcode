# Longest Consecutive Sequence - Explanation

## Problem Summary
Given an unsorted array of integers `nums`, return the length of the longest sequence of consecutive integers. The sequence must consist of consecutive numbers, but they don't need to be contiguous in the original array.

## Approach
This solution uses a **hash set with intelligent sequence detection**:

1. **Build Hash Set**: Insert all numbers into a set for O(1) lookups
2. **Identify Sequence Starts**: For each number, check if it's the start of a sequence
   - A number is a sequence start if `(number - 1)` doesn't exist in the set
3. **Expand Sequences**: From each start, count consecutive numbers by checking `number+1`, `number+2`, etc.
4. **Track Maximum**: Keep track of the longest sequence found

## Why This Works
The key insight is to only expand sequences from their starting point:
- **Avoid redundant work**: Each sequence is counted exactly once (from its start)
- **Hash set provides O(1) lookup**: Can quickly check if next number exists
- **Smart start detection**: `n-1` not in set means `n` starts a new sequence
- **Greedy expansion**: Once we find a start, we expand as far as possible

**Why this is O(n)**:
- Each number is added to the set once: O(n)
- Each number is visited as a potential start: O(n)
- Each number is expanded at most once: O(n) total expansion
- Even though expansion is in a nested loop, each number participates in expansion exactly once

## Time Complexity
**O(n)** where n is the length of the nums array
- Building the set: O(n)
- Iterating through set: O(n)
- Expansion appears nested but is O(n) total:
  - Each number is part of at most one sequence expansion
  - Total expansions across all sequences = n
- Overall: O(n) + O(n) + O(n) = O(n)

## Space Complexity
**O(n)** for the hash set
- Set stores all unique numbers from the array
- In worst case (all unique), stores n elements
- In best case (many duplicates), stores fewer

## Edge Cases
1. **Empty array**: Return 0
2. **Single element**: Return 1
3. **All duplicates**: [1,1,1,1] → return 1
4. **All consecutive**: [1,2,3,4,5] → return 5
5. **Multiple sequences**: [1,2,4,5,6] → return 3 (4,5,6)
6. **Negative numbers**: Works with any integers
7. **Unsorted**: [100,4,200,1,3,2] → return 4 (1,2,3,4)

## Worked Example
For `nums = [100,4,200,1,3,2]`:

```
Step 1: Build set
set = {100, 4, 200, 1, 3, 2}

Step 2: Check each number as potential sequence start

Check 100:
  Is 99 in set? No ✓ (100 is a sequence start)
  Expand: 100 → 101? No
  Sequence length: 1

Check 4:
  Is 3 in set? Yes (4 is NOT a sequence start, skip)

Check 200:
  Is 199 in set? No ✓ (200 is a sequence start)
  Expand: 200 → 201? No
  Sequence length: 1

Check 1:
  Is 0 in set? No ✓ (1 is a sequence start)
  Expand: 1 → 2? Yes
          2 → 3? Yes
          3 → 4? Yes
          4 → 5? No
  Sequence length: 4 (1,2,3,4)
  longest = 4

Check 3:
  Is 2 in set? Yes (3 is NOT a sequence start, skip)

Check 2:
  Is 1 in set? Yes (2 is NOT a sequence start, skip)

Result: 4 (sequence: 1,2,3,4)
```

## Visual Representation
```
nums = [100, 4, 200, 1, 3, 2]

Sequences identified:
Sequence 1: [1 → 2 → 3 → 4]  length = 4 ✓
            ^start      
            (0 not in set)

Sequence 2: [100]  length = 1
            ^start
            (99 not in set)

Sequence 3: [200]  length = 1
            ^start
            (199 not in set)

Numbers 2, 3, 4 are skipped as starts because their 
predecessors (1, 2, 3) exist in the set.

Maximum length: 4
```

## Why Skip Non-Starts
```
When we encounter 4:
  Check: Is 3 in set? Yes
  → 4 is NOT a sequence start
  → Skip (will be counted when we start from 1)

This prevents counting the same sequence multiple times:
❌ Starting from 1: [1,2,3,4] ← Count this
✓ Starting from 2: [2,3,4]   ← Skip (part of above)
✓ Starting from 3: [3,4]     ← Skip (part of above)
✓ Starting from 4: [4]       ← Skip (part of above)
```

## Expansion Example
```
Starting from 1:
set = {1, 2, 3, 4, 100, 200}

current = 1, streak = 1
Is 2 in set? Yes → current = 2, streak = 2
Is 3 in set? Yes → current = 3, streak = 3
Is 4 in set? Yes → current = 4, streak = 4
Is 5 in set? No  → Stop

Final streak: 4
```

## Complexity Analysis Visual
```
For array: [1, 2, 3, 4, 5]

Build set: Visit each element once
[1] → set
[2] → set
[3] → set  } O(n)
[4] → set
[5] → set

Check starts:
1: Is 0 in set? No → START
   Expand: 1→2→3→4→5 (5 checks)
2: Is 1 in set? Yes → SKIP
3: Is 2 in set? Yes → SKIP  } O(n) total
4: Is 3 in set? Yes → SKIP    (each number checked
5: Is 4 in set? Yes → SKIP     as start once,
                               expanded once)
Total operations: n (build) + n (check starts) + n (expand) = O(n)
```

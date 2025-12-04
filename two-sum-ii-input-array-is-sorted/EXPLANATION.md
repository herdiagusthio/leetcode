# Explanation — Two Sum II (Input Array Is Sorted)

## Problem
Given a 1-indexed array of integers `numbers` that is already sorted in non-decreasing order, find two numbers such that they add up to a specific `target` number. Return the indices of the two numbers (1-based) as an integer array `[index1, index2]`.

## Idea
Because the input array `numbers` is sorted in non-decreasing order, we can maintain two indices:
- `left` starting at 0 (the first element)
- `right` starting at n-1 (the last element)

At each step compute `sum = numbers[left] + numbers[right]`:
- If `sum == target` we found the pair. Return `[left+1, right+1]` (the problem uses 1-based indices).
- If `sum < target` the current sum is too small; increment `left` to increase the sum.
- If `sum > target` the current sum is too large; decrement `right` to decrease the sum.

## Algorithm (step-by-step)
1. Initialize `left = 0` and `right = len(numbers) - 1`.
2. Loop while `left < right`:
   - Calculate `sum = numbers[left] + numbers[right]`.
   - If `sum == target`, return `[]int{left + 1, right + 1}`.
   - If `sum < target`, increment `left`.
   - If `sum > target`, decrement `right`.
3. If no pair is found (though the problem guarantees one), return `nil`.

## Complexity
- Time: O(n) — single pass using two pointers.
- Space: O(1) — constant extra space.

## Alternatives
- **Hash Map**: If the array were unsorted, use a hash map (value -> index) to find complements. This uses O(n) extra space.
- **Binary Search**: For each element, binary search for the complement. This is O(n log n).

## Implementation notes
- The problem guarantees exactly one solution and n >= 2, so the algorithm will always find and return a valid pair.
- If you run this on invalid input (unsorted array or no solution), the function returns `nil`.

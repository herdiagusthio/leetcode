# Explanation — Two Sum II (Input Array Is Sorted)

This document explains the two-pointer solution used in `solution.go` for the LeetCode problem "Two Sum II - Input Array Is Sorted".

High-level idea
----------------
Because the input array `numbers` is sorted in non-decreasing order, we can maintain two indices:

- `left` starting at 0 (the first element)
- `right` starting at n-1 (the last element)

At each step compute `sum = numbers[left] + numbers[right]`:

- If `sum == target` we found the pair. Return `[left+1, right+1]` (the problem uses 1-based indices).
- If `sum < target` the current sum is too small; increment `left` to increase the sum.
- If `sum > target` the current sum is too large; decrement `right` to decrease the sum.

Why it works
------------
When the array is sorted, moving `left` rightwards increases the sum (or keeps it same if duplicates), and moving `right` leftwards decreases the sum. The two-pointer movement explores the space of possible pairs efficiently and each element is examined at most once.

Step-by-step example
--------------------
Input: numbers = [2,7,11,15], target = 9

1. left=0 (2), right=3 (15): sum=17 > 9 -> right--
2. left=0 (2), right=2 (11): sum=13 > 9 -> right--
3. left=0 (2), right=1 (7): sum=9 == target -> return [1,2]

Complexity
----------
- Time: O(n) — single pass using two pointers.
- Space: O(1) — constant extra space.

Edge cases
----------
- The problem guarantees exactly one solution and n >= 2, so the algorithm will always find and return a valid pair.
- If you run this on invalid input (unsorted array or no solution), the function returns `nil` in our implementation. For production code you might return an error or panic depending on your needs.

Follow-ups
----------
- If the array were unsorted and you needed an O(n) solution, use a hash map (value -> index) to find complements, but that uses O(n) extra space.
- If you need to enumerate all unique pairs that sum to target in a sorted array (instead of returning one), the two-pointer pattern adapts by moving pointers inward and skipping duplicates.

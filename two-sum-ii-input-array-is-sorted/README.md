# Two Sum II - Input Array Is Sorted

LeetCode: https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/

Problem
-------
Given a 1-indexed array of integers `numbers` that is already sorted in non-decreasing order, find two numbers such that they add up to a specific `target` number. Return the indices of the two numbers (1-based) as an integer array `[index1, index2]` of length 2.

The tests guarantee exactly one solution and you may not use the same element twice. Your solution must use only constant extra space.

Examples
--------
- Input: numbers = [2,7,11,15], target = 9
  Output: [1,2]

- Input: numbers = [2,3,4], target = 6
  Output: [1,3]

- Input: numbers = [-1,0], target = -1
  Output: [1,2]

Approach (this folder)
----------------------
This package implements the classic two-pointer approach for sorted arrays:

1. Maintain two pointers, `left` at start and `right` at end.
2. Compute sum = numbers[left] + numbers[right].
   - If sum == target: return [left+1, right+1] (1-based indices).
   - If sum < target: advance left++ to increase sum.
   - If sum > target: decrement right-- to decrease sum.

This is O(n) time and O(1) extra space.

Files
-----
- `solution.go` — two-pointer `twoSum` function and small demo `main()`.
- `solution_test.go` — table-driven unit tests.

Complexity
----------
- Time: O(n)
- Space: O(1)

Notes
-----
- The function returns `nil` if no solution is found, but per problem guarantees a solution exists.
- This implementation assumes the input array is sorted; for unsorted arrays a different approach (hash table) would be needed, but that would use extra space.

Constraints
-----------
- 2 <= numbers.length <= 3 * 10^4
- -1000 <= numbers[i] <= 1000
- numbers is sorted in non-decreasing order
- -1000 <= target <= 1000

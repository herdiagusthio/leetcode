# Container With Most Water

LeetCode: https://leetcode.com/problems/container-with-most-water/

You are given an integer array `height` of length `n`. There are `n` vertical lines drawn such that the two endpoints of the `i`th line are `(i, 0)` and `(i, height[i])`.

Find two lines that together with the x-axis form a container, such that the container contains the most water. Return the maximum amount of water a container can store. You may not slant the container.

## Examples

- Example 1
  - Input: `height = [1,8,6,2,5,4,8,3,7]`
  - Output: `49`

- Example 2
  - Input: `height = [1,1]`
  - Output: `1`

## Constraints

- `n == height.length`
- `2 <= n <= 10^5`
- `0 <= height[i] <= 10^4`

## Approach (brief)

Use the two-pointer technique:

- Initialize `left = 0` and `right = n-1`.
- While `left < right`, compute `area = (right-left) * min(height[left], height[right])` and update the best result.
- Move the pointer that points to the smaller height inward, since moving the taller one cannot produce a larger area with the same width.

This runs in O(n) time and uses O(1) extra space.

## Complexity

- Time: O(n)
- Space: O(1)

## Run / Test

From the repository root:

```bash
# Run the example main in this folder
go run ./container-with-most-water

# Run tests for this package
go test ./container-with-most-water
```

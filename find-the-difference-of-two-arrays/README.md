# Find the Difference of Two Arrays

**LeetCode problem:**  
https://leetcode.com/problems/find-the-difference-of-two-arrays/

Given two 0-indexed integer arrays `nums1` and `nums2`, return a list answer of size 2 where:

- `answer[0]` is a list of all distinct integers in `nums1` which are not present in `nums2`
- `answer[1]` is a list of all distinct integers in `nums2` which are not present in `nums1`

Note that the integers in the lists may be returned in any order.

**Example 1:**
```
Input: nums1 = [1,2,3], nums2 = [2,4,6]
Output: [[1,3],[4,6]]
Explanation:
For nums1, nums1[1] = 2 is present at index 0 of nums2, whereas nums1[0] = 1 and nums1[2] = 3 are not present in nums2. Therefore, answer[0] = [1,3].
For nums2, nums2[0] = 2 is present at index 1 of nums1, whereas nums2[1] = 4 and nums2[2] = 6 are not present in nums1. Therefore, answer[1] = [4,6].
```

**Example 2:**
```
Input: nums1 = [1,2,3,3], nums2 = [1,1,2,2]
Output: [[3],[]]
Explanation:
For nums1, nums1[2] and nums1[3] are not present in nums2. Since nums1[2] == nums1[3], their value is only included once and answer[0] = [3].
Every integer in nums2 is present in nums1. Therefore, answer[1] = [].
```

**Constraints:**
- `1 <= nums1.length, nums2.length <= 1000`
- `-1000 <= nums1[i], nums2[i] <= 1000`

## Solution

The solution uses a single hash map with state tracking to efficiently find distinct elements. Each number is mapped to a state: 0 (only in nums1), 1 (only in nums2), or 2 (in both). This approach achieves optimal time complexity with better memory efficiency and cache locality compared to using two separate sets.

**Time Complexity:** O(n + m) where n = len(nums1) and m = len(nums2)  
**Space Complexity:** O(n + m)

See [EXPLANATION.md](./EXPLANATION.md) for a detailed breakdown of the algorithm.

## How to run

```bash
# Run tests
go test -v

# Run the solution with example inputs
go run .
```

## Files

- `solution.go` - Main solution implementation with single-map state tracking
- `solution_test.go` - Table-driven tests covering examples and edge cases
- `EXPLANATION.md` - Detailed explanation of the algorithm and approach
- `README.md` - This file
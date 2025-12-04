# Product of Array Except Self

LeetCode problem: Product of Array Except Self

Link: https://leetcode.com/problems/product-of-array-except-self/

Given an integer array `nums`, return an array `answer` such that `answer[i]` is equal to the product of all the elements of `nums` except `nums[i]`.

## Solution
This package implements the standard O(n) solution without using division: compute prefix products for each index, then multiply by suffix products in a second pass. The implementation uses constant extra space (ignoring the output array) and runs in linear time.

## How to run
From this folder:

```bash
go test -v
```

To run the example `main` in this package:

```bash
go run .
```

## Files
- `solution.go` — implementation and example `main`.
- `solution_test.go` — unit tests.
- `EXPLANATION.md` — step-by-step explanation of the approach.
Leetcode problem: Product of Array Except Self

link : https://leetcode.com/problems/product-of-array-except-self/description/

Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].

The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.

You must write an algorithm that runs in O(n) time and without using the division operation.

 

Example 1:

Input: nums = [1,2,3,4]
Output: [24,12,8,6]
Example 2:

Input: nums = [-1,1,0,-3,3]
Output: [0,0,9,0,0]
 

Constraints:

2 <= nums.length <= 105
-30 <= nums[i] <= 30
The input is generated such that answer[i] is guaranteed to fit in a 32-bit integer.
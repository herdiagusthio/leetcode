Partition Array According to Given Pivot
======================================

LeetCode: https://leetcode.com/problems/partition-array-according-to-given-pivot

Problem
-------

Given a 0-indexed integer array `nums` and an integer `pivot`, rearrange `nums` so that:

- Every element less than `pivot` appears before every element greater than `pivot`.
- Every element equal to `pivot` appears in between the elements less than and greater than `pivot`.
- The relative order of the elements less than `pivot` and the elements greater than `pivot` is maintained.

Return `nums` after the rearrangement.

Short solution
--------------

A simple, stable approach is to collect elements into three buckets while scanning the array:

- `left` for elements < pivot
- `mid` for elements == pivot
- `right` for elements > pivot

Then concatenate `left + mid + right`. This preserves the relative order inside each partition.

This folder contains a few variants:

- `pivotArray` — straightforward append-based buckets (clear and stable).
- `pivotArrayPrealloc` — same append logic but preallocates capacities for fewer reallocations.
- `pivotArrayCount` — two-pass approach that preallocates the exact resulting slice and fills it (best for minimizing allocations).

Complexity
----------

- Time: O(n) — single pass (or two passes for the preallocate-fill variant).
- Space: O(n) extra — buckets or one preallocated result slice.

Example
-------

Input:

```go
nums = []int{9,12,5,10,14,3,10}
pivot = 10
```

Output:

```
[9,5,3,10,10,12,14]
```

How to run
----------

- Run the sample program in this folder:

```bash
cd partition-array-according-to-given-pivot
go run .
```

- Run the package tests:

```bash
go test ./partition-array-according-to-given-pivot
```

Or run all tests in the repo:

```bash
go test ./...
```

Files in this folder
--------------------

- `solution.go` — implementations and example `main`.
- `solution_test.go` — unit tests.
- `EXPLANATION.md` — detailed explanation and tradeoffs.

Notes
-----

- Use `pivotArray` for readability. For performance-sensitive use, prefer `pivotArrayCount`.
Partition Array According to Given Pivot
======================================

LeetCode: https://leetcode.com/problems/partition-array-according-to-given-pivot

Problem
-------

Given a 0-indexed integer array `nums` and an integer `pivot`, rearrange `nums` so that:

- Every element less than `pivot` appears before every element greater than `pivot`.
- Every element equal to `pivot` appears in between the elements less than and greater than `pivot`.
- The relative order of the elements less than `pivot` and the elements greater than `pivot` is maintained.

Return `nums` after the rearrangement.

Short solution
--------------

A simple, stable approach is to collect elements into three buckets while scanning the array:

- `left` for elements < pivot
- `mid` for elements == pivot
- `right` for elements > pivot

Then concatenate `left + mid + right`. This preserves the relative order inside each partition.

This folder contains a few variants:

- `pivotArray` — straightforward append-based buckets (clear and stable).
- `pivotArrayPrealloc` — same append logic but preallocates capacities for fewer reallocations.
- `pivotArrayCount` — two-pass approach that preallocates the exact resulting slice and fills it (best for minimizing allocations).

Complexity
----------

- Time: O(n) — single pass (or two passes for the preallocate-fill variant).
- Space: O(n) extra — buckets or one preallocated result slice.

Example
-------

Input:

```go
nums = []int{9,12,5,10,14,3,10}
pivot = 10
```

Output:

```
[9,5,3,10,10,12,14]
```

How to run
----------

- Run the sample program in this folder:

```bash
cd partition-array-according-to-given-pivot
go run .
```

- Run the package tests:

```bash
go test ./partition-array-according-to-given-pivot
```

Or run all tests in the repo:

```bash
go test ./...
```

Files in this folder
--------------------

- `solution.go` — implementations and example `main`.
- `solution_test.go` — unit tests.
- `EXPLANATION.md` — detailed explanation and tradeoffs.

Notes
-----

- Use `pivotArray` for readability. For performance-sensitive use, prefer `pivotArrayCount`.
Partition Array According to Given Pivot
======================================

LeetCode: https://leetcode.com/problems/partition-array-according-to-given-pivot

Problem
-------

Given a 0-indexed integer array `nums` and an integer `pivot`, rearrange `nums` so that:

- Every element less than `pivot` appears before every element greater than `pivot`.
- Every element equal to `pivot` appears in between the elements less than and greater than `pivot`.
- The relative order of the elements less than `pivot` and the elements greater than `pivot` is maintained.

Return `nums` after the rearrangement.

Short solution
--------------

A simple, stable approach is to collect elements into three buckets while scanning the array:

- `left` for elements < pivot
- `mid` for elements == pivot
- `right` for elements > pivot

Then concatenate `left + mid + right`. This preserves the relative order inside each partition.

This folder contains a few variants:

- `pivotArray` — straightforward append-based buckets (clear and stable).
- `pivotArrayPrealloc` — same append logic but preallocates capacities for fewer reallocations.
- `pivotArrayCount` — two-pass approach that preallocates the exact resulting slice and fills it (best for minimizing allocations).

Complexity
----------

- Time: O(n) — single pass (or two passes for the preallocate-fill variant).
- Space: O(n) extra — buckets or one preallocated result slice.

Example
-------

Input:

```go
nums = []int{9,12,5,10,14,3,10}
pivot = 10
```

Output:

```
[9,5,3,10,10,12,14]
```

How to run
----------

- Run the sample program in this folder:

```bash
cd partition-array-according-to-given-pivot
go run .
```

- Run the package tests:

```bash
go test ./partition-array-according-to-given-pivot
```

Or run all tests in the repo:

```bash
go test ./...
```

Files in this folder
--------------------

- `solution.go` — implementations and example `main`.
- `solution_test.go` — unit tests.
- `EXPLANATION.md` — detailed explanation and tradeoffs.

Notes
-----

- Use `pivotArray` for readability. For performance-sensitive use, prefer `pivotArrayCount`.
Leetcode: https://leetcode.com/problems/partition-array-according-to-given-pivot


You are given a 0-indexed integer array nums and an integer pivot. Rearrange nums such that the following conditions are satisfied:

Every element less than pivot appears before every element greater than pivot.
Every element equal to pivot appears in between the elements less than and greater than pivot.
The relative order of the elements less than pivot and the elements greater than pivot is maintained.
More formally, consider every pi, pj where pi is the new position of the ith element and pj is the new position of the jth element. If i < j and both elements are smaller (or larger) than pivot, then pi < pj.
Return nums after the rearrangement.

 

Example 1:

Input: nums = [9,12,5,10,14,3,10], pivot = 10
Output: [9,5,3,10,10,12,14]
Explanation: 
The elements 9, 5, and 3 are less than the pivot so they are on the left side of the array.
The elements 12 and 14 are greater than the pivot so they are on the right side of the array.
The relative ordering of the elements less than and greater than pivot is also maintained. [9, 5, 3] and [12, 14] are the respective orderings.
Example 2:

Input: nums = [-3,4,3,2], pivot = 2
Output: [-3,2,4,3]
Explanation: 
The element -3 is less than the pivot so it is on the left side of the array.
The elements 4 and 3 are greater than the pivot so they are on the right side of the array.
The relative ordering of the elements less than and greater than pivot is also maintained. [-3] and [4, 3] are the respective orderings.
 

Constraints:

1 <= nums.length <= 105
-106 <= nums[i] <= 106
pivot equals to an element of nums.
# Partition List

LeetCode problem: Partition List

Link: https://leetcode.com/problems/partition-list/

Given the head of a singly linked list and an integer x, partition the list so that all nodes with values less than x come before nodes with values greater than or equal to x. Preserve the original relative order of nodes in each partition.

## Solution
Scan the list once and keep two tails (or two lists): `less` for values < x and `greater` for values >= x. Append each node to the corresponding tail. After the traversal, concatenate `less` then `greater`.

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
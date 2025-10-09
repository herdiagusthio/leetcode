## Partition List

LeetCode problem: https://leetcode.com/problems/partition-list/

## Problem

Given the head of a singly linked list and an integer x, partition the list so that all nodes with values less than x come before nodes with values greater than or equal to x. Preserve the original relative order of nodes in each partition.

## Solution (short)

Scan the list once and keep two tails (or two lists): `less` for values < x and `greater` for values >= x. Append each node to the corresponding tail. After the traversal, concatenate `less` then `greater`.

This folder contains two implementations:

- `partition` — pointer-based solution that uses dummy heads (O(n) time, O(1) extra space).
- `partitionArray` — collects nodes into slices and reconnects them (O(n) time, O(n) extra space).

## Complexity

- Time: O(n) — single pass.
- Space: O(1) extra for `partition`; O(n) extra for `partitionArray`.

## Example

Input:

```go
head = [1,4,3,2,5,2], x = 3
```

Output:

```
[1,2,2,4,3,5]
```

## How to run

- Run the sample program (prints example output):

```bash
cd partition-list
go run .
```

- Run tests for this package only:

```bash
go test ./partition-list
```

- Run all tests in the repository:

```bash
go test ./...
```

If you run without Go modules (GOPATH):

```bash
cd partition-list
GO111MODULE=off go test
```

## Files in this folder

- `solution.go` — implementations and small example `main`.
- `solution_test.go` — unit tests.
- `EXPLANATION.md` — detailed explanation and notes.

## Notes

- The `partition` implementation preserves relative order inside each partition (stable).
- If you want the example runner removed or moved into a `cmd/` folder, I can change that.
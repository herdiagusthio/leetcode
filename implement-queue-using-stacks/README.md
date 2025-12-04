# Implement Queue using Stacks

LeetCode problem: Implement Queue using Stacks

Link: https://leetcode.com/problems/implement-queue-using-stacks/

Implement a first in first out (FIFO) queue using only two stacks. The implemented queue should support all the functions of a normal queue (push, peek, pop, and empty).

## Solution
Use two stacks (`in` and `out`). Push to `in`. For pop/peek, if `out` is empty, move all elements from `in` to `out`. This reverses the order so `out` behaves like a queue.

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
- `solution.go` — implementation.
- `solution_test.go` — unit tests.
- `EXPLANATION.md` — step-by-step explanation of the approach.
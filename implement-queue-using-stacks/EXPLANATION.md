# Implement Queue using Stacks — Explanation

## Problem
Implement a FIFO queue using only two stacks. Support push, pop, peek, and empty operations using only stack operations.

## High-level idea
Use two stacks (`in` and `out`) to simulate queue behavior:
- Push: Always push to the `in` stack.
- Pop/Peek: If `out` is empty, move all elements from `in` to `out` (reversing order). Then pop/peek from `out`.
- Empty: True if both stacks are empty.

This ensures the oldest element is always on top of `out` for queue-like access.

## Approach
1. **Push(x):** `in = append(in, x)`
2. **Pop():** If `out` is empty, move all from `in` to `out`. Pop from `out`.
3. **Peek():** If `out` is empty, move all from `in` to `out`. Return top of `out`.
4. **Empty():** Return `len(in) == 0 && len(out) == 0`

## Complexity
- Time: Amortized O(1) per operation (each element is moved at most once between stacks).
- Space: O(n) for n elements in the queue.

## Edge cases
- Multiple pushes before any pop/peek
- Alternating push/pop
- Empty queue

## Notes
- Only stack operations are used (append, pop from end, len).
- This is the canonical solution for this problem and is efficient for all valid LeetCode test cases.

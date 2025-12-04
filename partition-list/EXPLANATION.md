# Explanation — Partition List

## Problem
Given the head of a singly linked list and a value x, partition the list so that all nodes with values less than x come before nodes with values greater than or equal to x. You must preserve the original relative order of the nodes in each of the two partitions.

## Idea
We split the list into two lists while traversing the original list once:
- "less" list contains nodes with value < x (order preserved)
- "greater" list contains nodes with value >= x (order preserved)

At the end of the traversal we connect the tail of the "less" list to the head of the "greater" list, and terminate the combined list.

## Algorithm (step-by-step)
1. Create two dummy nodes: `lessHead` and `greaterHead`. Maintain two tails: `lessTail` and `greaterTail`.
2. Iterate the input list with a pointer `curr`:
   - If `curr.Val < x`: attach `curr` to `lessTail.Next` and move `lessTail`.
   - Else: attach `curr` to `greaterTail.Next` and move `greaterTail`.
   - Move `curr` to `curr.Next`.
3. After the loop, terminate the greater list with `greaterTail.Next = nil` to avoid cycles.
4. Connect `lessTail.Next = greaterHead.Next` and return `lessHead.Next` as the new head.

## Complexity
- Time: O(n) — single pass where n is number of nodes.
- Space: O(1) — extra space (we only use a few pointers/dummy nodes).

## Alternatives
- **Array collect-then-reconnect**: Traverse the list and collect pointers to nodes into two slices, then reconnect. This is O(n) extra space.

## Implementation notes
- It's critical to terminate the greater list (set `tail.Next = nil`) after reconnection to avoid cycles.
- The dummy-node pattern makes corner cases (empty partitions) easy to handle.

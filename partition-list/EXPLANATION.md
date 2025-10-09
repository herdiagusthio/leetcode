Partition List (LeetCode)
=========================

Problem
-------
Given the head of a singly linked list and a value x, partition the list so that all nodes with values less than x come before nodes with values greater than or equal to x.

You must preserve the original relative order of the nodes in each of the two partitions.

High-level idea
----------------
We split the list into two lists while traversing the original list once:
- "less" list contains nodes with value < x (order preserved)
- "greater" list contains nodes with value >= x (order preserved)

At the end of the traversal we connect the tail of the "less" list to the head of the "greater" list, and terminate the combined list.

Approach 1 — Dummy-pointer, O(1) extra space (stable)
-------------------------------------------------------
Algorithm outline:
1. Create two dummy nodes: lessHead and greaterHead. Maintain two tails: lessTail and greaterTail.
2. Iterate the input list with a pointer curr:
   - If curr.Val < x: attach curr to lessTail.Next and move lessTail.
   - Else: attach curr to greaterTail.Next and move greaterTail.
   - Move curr to curr.Next.
3. After the loop, terminate the greater list with greaterTail.Next = nil to avoid cycles.
4. Connect lessTail.Next = greaterHead.Next and return lessHead.Next as the new head.

Why this works:
- Every node is appended to exactly one list, preserving original order since we append by visiting nodes in list order.
- Using dummy heads simplifies handling empty partitions and removing special cases.
- Termination of the greater list ensures there are no cycles if original next pointers still point to later nodes.

Complexity:
- Time: O(n), single pass where n is number of nodes.
- Space: O(1) extra space (we only use a few pointers/dummy nodes).

Approach 2 — Array (slice) collect-then-reconnect, O(n) extra space
------------------------------------------------------------------
Algorithm outline:
1. Traverse the list and collect pointers to nodes into two slices: less[] and greater[].
2. Reconnect nodes in the less slice in order, then nodes in the greater slice in order.
3. Terminate the final tail.Next = nil and return the head of the reconstructed list.

Why this works:
- Order preserved by the slices which collect nodes in traversal order.
- Reconnection is straightforward and avoids pointer juggling during the traversal.

Complexity:
- Time: O(n) to collect and O(n) to reconnect => O(n).
- Space: O(n) extra pointers stored in slices.

Edge cases
----------
- Empty list (head == nil) — return nil.
- All nodes less than x — the greater list is empty; algorithm returns original order.
- All nodes greater than or equal to x — the less list is empty; algorithm returns original order.
- Duplicate values equal to x — duplicates go to the greater partition and their order is preserved.

Correctness proofs (informal)
----------------------------
- Stability: Both approaches process nodes in original order and append nodes to the end of the respective partition lists. Therefore, the relative order of nodes within each partition is unchanged.
- Partitioning property: Every node is placed in less if val < x, otherwise in greater. The concatenation of less followed by greater satisfies the required property.

Implementation notes
--------------------
- It's critical to terminate the greater list (set tail.Next = nil) after reconnection to avoid cycles when nodes' original Next pointers pointed into the other partition.
- The dummy-node pattern (lessHead/greaterHead) makes corner cases (empty partitions) easy to handle.

Testing
-------
Include tests for:
- Example cases from the problem statement.
- Nil / single-node lists.
- Lists with duplicates and negative numbers.
- Lists where all nodes fall into one partition.

Summary
-------
- Preferred solution: dummy-pointer partition (Approach 1) — O(n) time, O(1) extra space, stable.
- Alternate solution: collect-then-reconnect (Approach 2) — O(n) time, O(n) extra space, easier to reason about.

Both produce correct, stable partitions when implemented carefully.

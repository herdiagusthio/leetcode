Explanation — Partition Array According to Given Pivot
======================================================

Problem recap
--------------

We must reorder `nums` so that all elements less than `pivot` come first, elements equal to `pivot` in the middle, and elements greater than `pivot` last. The relative order of elements within the "less-than" group and within the "greater-than" group must be preserved (stable partition).

Approaches implemented
----------------------

1) Three buckets (pivotArray)
   - Walk the input once, appending items to three slices: `left`, `mid`, `right`.
   - Return `append(append(left, mid...), right...)`.
   - Pros: Simple, stable, easy to read.
   - Cons: May cause multiple allocations if capacities aren't pre-sized.

2) Preallocated buckets (pivotArrayPrealloc)
   - Same logic as (1) but we preallocate each bucket's capacity using a heuristic (input length or a small growth factor) to reduce reallocations.
   - Slightly more efficient in practice for large inputs.

3) Two-pass preallocate-and-fill (pivotArrayCount)
   - First pass: count numbers in each partition to compute the exact size of the resulting slice.
   - Allocate a single result slice with len == n and use three pointers to fill the result in one more pass.
   - Pros: Minimal allocations (one result slice), predictable memory.
   - Cons: Slightly more code; still O(n) time and O(n) extra space.

Why not in-place DNF?
----------------------

The Dutch National Flag (DNF) algorithm partitions an array in-place into three parts in O(n) time and O(1) extra space. However, it does not preserve the relative order of elements — it's unstable. The LeetCode problem explicitly requires preserving relative order for elements less than and greater than the pivot. Therefore, DNF is not acceptable here unless stability is relaxed. For demonstration, a DNF variant can be implemented but was removed from the default solutions to respect the problem constraints.

Complexity
----------

- Time: O(n) for all stable variants. The two-pass variant does two linear passes but still O(n).
- Space: O(n) extra — either multiple small slices (buckets) or a single preallocated result slice.

Edge cases and tests
--------------------

- Empty input and single-element arrays.
- All elements equal to pivot (result equals input).
- Duplicates around pivot (stability verified in tests).

Notes
-----

- For production code where minimizing allocations matters, prefer `pivotArrayCount`.
- Keep the `pivot` guaranteed to exist in `nums` as the problem states; otherwise, handle missing pivot by treating it as an ordinary value.

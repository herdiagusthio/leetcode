# Explanation — Max Number of K-Sum Pairs

Problem summary
- Given an integer array `nums` and an integer `k`, in one operation you may pick two numbers
  whose sum equals `k` and remove them. Return the maximum number of such operations.

Approach
- Use a hash map (`counts`) to keep track of how many times we've seen each value while
  scanning the array once.
- For each element `v` in `nums`, compute its complement `diff = k - v`.
  - If `counts[diff] > 0`, a previously-seen complement is available: form one pair,
    increment the result and decrement `counts[diff]`.
  - Otherwise, record `v` by incrementing `counts[v]`.
- This greedy approach matches each number with the earliest available complement and
  ensures pairs are disjoint (we decrement counts when we use them).

Why this is correct
- Each time we find a complement we remove both numbers from further consideration by
  decrementing the stored count; thus pairs are disjoint.
- The single-pass greedy matching maximizes the number of pairs because any available
  complement can only be used once; using it immediately cannot reduce the number of
  total pairs that can be formed later.

Complexity
- Time: O(n) — single pass through `nums`.
- Space: O(u) — extra space for the map where `u` is the number of distinct values in `nums`.

Notes and edge cases
- Works when `v == diff` (e.g., `k` even and `v == k/2`) because we count occurrences and
  require two instances to form a pair.
- Works correctly when numbers are zero or negative (map-based approach is general).
- Preallocating the map with `len(nums)` improves performance for large inputs.

Small example
- Input: `nums = [1,2,3,4]`, `k = 5`
  - Process `1`: record `counts[1]=1`
  - Process `2`: record `counts[2]=1`
  - Process `3`: complement `2` exists → form pair (2,3), decrement `counts[2]`
  - Process `4`: complement `1` exists → form pair (1,4), decrement `counts[1]`
  - Result = 2

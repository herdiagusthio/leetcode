## Generate Parentheses (LeetCode)

### Problem

Given n pairs of parentheses, generate all combinations of well-formed parentheses.

### High-level idea

Use DFS/backtracking while tracking two counters: number of open parentheses used and number of close parentheses used. At each step:

- add '(' when open < n
- add ')' when close < open

When the current buffer reaches length 2*n, append it to the results.

### Approaches

1) Backtracking (recommended)

- Simple, readable, and produces exactly the set of valid sequences.
- Implementation detail: either pass a growing slice and use `append` (clear and idiomatic), or write directly into a preallocated byte buffer to reduce allocations.

2) In-place buffer (allocation-optimized)

- Preallocate `buf := make([]byte, 2*n)` and write characters by index.
- Convert `string(buf)` when a solution is complete. This minimizes transient allocations during recursion.

3) Dynamic programming (constructive)

- Possible via Catalan decomposition (combine smaller valid sequences), but usually more complex and less direct than backtracking for this problem.

### Complexity

- Time: O(C_n * n), where C_n is the nth Catalan number (the number of valid sequences). Each sequence length is 2*n.
- Space: O(C_n * n) for the output plus O(n) recursion depth (or O(2*n) for the in-place buffer).

### Edge cases

- n = 0 -> return [""]
- Validate small n (1..3) against known outputs.

### Implementation notes

- For clarity use the recursive append-based approach. If allocation pressure matters, switch to the in-place buffer variant.
- Remember converting a `[]byte` to `string` allocates a new string; that cost is per result and unavoidable.

### Testing

- Use unit tests that compare sets (order-insensitive) for small n. The repository's `solution_test.go` already covers n=0..3.

### Summary

- Preferred: backtracking/DFS for readability.
- Optimize with an in-place buffer if necessary.

# Summary Ranges

LeetCode problem: Summary Ranges

Given a sorted, unique integer array, return the smallest sorted list of ranges that cover all the numbers in the array. Each range `[a,b]` should be output as:

- `"a->b"` if `a != b`
- `"a"` if `a == b`

Link: https://leetcode.com/problems/summary-ranges/

## Solution
This folder contains an O(n) solution that scans the array once, groups consecutive numbers into ranges, and formats them as strings using `strconv.Itoa`.

## How to run
From this folder:

```bash
go test -v
```

Benchmarks:

```bash
go test -bench . -benchmem
```

## Files
- `solution.go` — implementation using a map-free single pass.
- `solution_test.go` — unit tests and benchmarks comparing `fmt.Sprintf` vs `strconv` variants.
- `EXPLANATION.md` — step-by-step reasoning and complexity analysis.
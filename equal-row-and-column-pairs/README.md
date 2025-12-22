# Equal Row and Column Pairs

LeetCode problem: Equal Row and Column Pairs

Link: https://leetcode.com/problems/equal-row-and-column-pairs/description/

Given a 0-indexed n x n integer matrix grid, return the number of pairs (ri, cj) such that row ri and column cj are equal.

A row and column pair is considered equal if they contain the same elements in the same order (i.e., an equal array).

## Solution
We convert each row into a string key and store the frequency in a map. Then, we iterate through each column, convert it to the same string key format, and check if it exists in the map. If it does, we add the frequency to our count.

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
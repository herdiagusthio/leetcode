# String Compression

LeetCode problem: String Compression

Link: https://leetcode.com/problems/string-compression/description/

Given an array of characters `chars`, compress it using the following algorithm:

Begin with an empty string `s`. For each group of consecutive repeating characters in `chars`:

- If the group's length is 1, append the character to `s`.
- Otherwise, append the character followed by the group's length.

The compressed string `s` should not be returned separately, but instead, be stored in the input character array `chars`. Note that group lengths that are 10 or longer will be split into multiple characters in `chars`.

After you are done modifying the input array, return the new length of the array.

You must write an algorithm that uses only constant extra space.

## Solution
This package implements a solution that iterates through the input array, counts consecutive characters, and builds the compressed string.

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
# Explanation — Determine if Two Strings Are Close

## Problem
Given two strings `word1` and `word2`, determine if they are "close". Two strings are considered close if you can transform one into the other using two operations any number of times:
1. **Swap any two existing characters** (e.g., `abc` → `acb`)
2. **Transform all occurrences of one character into another, and vice versa** (e.g., `aacabb` → `bbcbaa`)

Examples:
- `word1 = "abc", word2 = "bca"` → `true` (can reorder)
- `word1 = "a", word2 = "aa"` → `false` (different lengths)
- `word1 = "cabbba", word2 = "abbccc"` → `true` (can transform)

## Idea
The key insight is that two strings are "close" if and only if they satisfy three conditions:

1. **Same length**: If the strings have different lengths, no operations can make them equal.
2. **Same character set**: Both strings must contain exactly the same set of unique characters. Operation 2 only transforms existing characters, so we can't create new ones.
3. **Same frequency distribution**: When we sort the frequencies of characters in both strings, they must match. Operation 1 allows reordering (doesn't affect frequencies), and Operation 2 allows swapping frequencies between characters.

For example, if `word1` has characters with frequencies `[3, 2, 1]` and `word2` has `[2, 3, 1]`, they have the same frequency distribution when sorted: `[1, 2, 3]`.

## Algorithm (step-by-step)

1. **Check lengths**: If `len(word1) != len(word2)`, return `false`.

2. **Count character frequencies**:
   - Create a hash map `freqWord1` mapping each character in `word1` to its count.
   - Create a hash map `freqWord2` mapping each character in `word2` to its count.

3. **Check character sets**:
   - If the two maps have different sizes, return `false`.
   - For each character in `freqWord1`, check if it exists in `freqWord2`. If not, return `false`.

4. **Extract and sort frequency counts**:
   - Extract all frequency values from `freqWord1` into an array `countsWord1`.
   - Extract all frequency values from `freqWord2` into an array `countsWord2`.
   - Sort both arrays.

5. **Compare sorted frequencies**:
   - Compare the two sorted arrays element by element.
   - If they match, return `true`; otherwise, return `false`.

## Complexity
- **Time**: O(n + m log m)
  - O(n) to iterate through both strings and count frequencies (where n is the string length).
  - O(m log m) to sort the frequency arrays (where m is the number of unique characters, at most 26 for lowercase letters).
  - O(m) to compare the sorted arrays.
  - Overall: O(n + m log m), which simplifies to O(n) when m is constant (26).

- **Space**: O(m)
  - O(m) for the two hash maps storing character frequencies.
  - O(m) for the two arrays storing frequency counts.
  - Overall: O(m), which is O(1) when considering only lowercase English letters (m ≤ 26).

## Alternatives
- **Using fixed-size arrays**: Instead of hash maps, use two arrays of size 26 for lowercase letters (`freq[char-'a']++`). This is slightly faster but assumes the input contains only lowercase English letters.
- **Direct frequency matching**: After extracting frequencies, you could use a multiset or additional hash map to compare frequency counts without sorting, but sorting is simpler and efficient enough for this problem.

## Implementation notes
- The solution uses `map[rune]int` to handle any Unicode characters, though the problem guarantees only lowercase English letters.
- We initialize slices with capacity (`make([]int, 0, capacity)`) to avoid initial zero values when using `append`.
- The character set check is crucial: even if two strings have the same frequency distribution, they're not close if they don't share the same characters (e.g., `"abc"` vs `"def"`).

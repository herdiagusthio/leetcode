# Generate Parentheses - Explanation

## Problem Summary
Given `n` pairs of parentheses, generate all valid (well-formed) combinations of parentheses.

## Approach
This solution uses **backtracking with constraints**:

1. **Track Counts**: Maintain count of open '(' and close ')' parentheses used
2. **Base Case**: When string length reaches `2*n`, add to results
3. **Add Open Paren**: If `open < n`, we can add '(' and recurse
4. **Add Close Paren**: If `close < open`, we can add ')' and recurse
   - This ensures we never have more ')' than '(' at any point
5. **Backtrack**: Explore all valid paths through the decision tree

## Why This Works
The key insights that make this approach correct:
- **Valid parentheses constraint**: At any point, number of ')' cannot exceed number of '('
- **Complete sequences**: We need exactly n open and n close parentheses
- **Backtracking explores all paths**: By trying both options at each step (when valid), we generate all possible combinations
- **Pruning invalid paths**: The conditions `open < n` and `close < open` prevent invalid sequences from being generated
- **No duplicates**: Each path through the decision tree is unique

**Why `close < open` works**:
- If we've used more or equal ')' as '(', adding another ')' creates an invalid sequence
- Example: ")(" is invalid, "()" is valid
- This constraint maintains the "well-formed" property

## Time Complexity
**O(4^n / √n)** which is the nth Catalan number multiplied by the work per result
- Number of valid sequences: C_n = (2n)! / ((n+1)! × n!) ≈ 4^n / (n^(3/2) × √π)
- This is the nth Catalan number
- For each valid sequence, we build a string of length 2n
- Total: O(C_n × n) ≈ O(4^n / √n)

More precisely: O(n × C_n) where C_n is the nth Catalan number

## Space Complexity
**O(n × C_n)** for storing all results
- Output space: C_n strings, each of length 2n
- Recursion depth: O(n) - maximum depth is 2n (length of each string)
- Temporary string buffer: O(n)
- Total dominated by output: O(n × C_n)

## Edge Cases
1. **n = 0**: Returns [""] (empty string)
2. **n = 1**: Returns ["()"]
3. **n = 2**: Returns ["(())", "()()"]
4. **Large n**: Exponential growth in number of combinations

## Worked Example
For `n = 3`:

```
Start: open=0, close=0, current=""

Decision tree (simplified):
                    ""
                    |
            ┌───────┴───────┐
           (                (impossible to add ')')
         o=1,c=0           
            |
      ┌─────┴─────┐
    ((           ()
  o=2,c=0      o=1,c=1
     |            |
  ┌──┴──┐      ┌─┴─┐
 (((   (()    ()(  ()(impossible to add 3rd '(')
o=3,c=0 o=2,c=1 o=2,c=1

Continue expanding...

Valid sequences found:
1. "((()))" - open parentheses: ((( then close: )))
2. "(()())" - ((, close one ), open (, close ))
3. "(())()" - ((, close )), open (, close )
4. "()(())" - (, close ), open ((, close ))
5. "()()()" - (, close ), (, close ), (, close )
```

Complete generation trace for `n=2`:

```
Step-by-step:
1. Start: "" → add '(' → "("
2. "(" → add '(' → "(("
3. "((" → can't add '(' (open=2=n), add ')' → "(()"
4. "(()" → add ')' → "(())" ✓ VALID (length=4=2n)

Backtrack to "((" and try different path...
5. "(" → add ')' → "()"
6. "()" → add '(' → "()("
7. "()(" → add ')' → "()()" ✓ VALID

Final result: ["(())", "()()"]
```

## Visual Representation
Decision tree for `n=2`:

```
                         start=""
                          o=0,c=0
                             |
                            '('
                          o=1,c=0
                        /          \
                      '('          ')'
                   o=2,c=0      o=1,c=1
                      |             |
                     ')'           '('
                   o=2,c=1      o=2,c=1
                      |             |
                     ')'           ')'
                   o=2,c=2      o=2,c=2
                      |             |
                   "(())"        "()()"
                   ✓ VALID      ✓ VALID
```

## Generation Pattern for n=3
```
All valid combinations (5 total):
1. ((()))  - max nesting
2. (()())  - 
3. (())()  - 
4. ()(())  - 
5. ()()()  - no nesting

The pattern follows the Catalan number sequence:
n=1: 1 combination
n=2: 2 combinations
n=3: 5 combinations
n=4: 14 combinations
```

## Constraints That Ensure Validity
```
At each step, to maintain validity:
┌─────────────────────────────────┐
│ Rule 1: open < n                │  Can add '('
│         (haven't used all opens)│
└─────────────────────────────────┘

┌─────────────────────────────────┐
│ Rule 2: close < open            │  Can add ')'
│         (balance maintained)    │
└─────────────────────────────────┘

Example of rule 2 preventing invalid:
  Current: "()"
  open=1, close=1
  Can we add ')'? close(1) < open(1)? NO
  This prevents: "())" which is invalid
```

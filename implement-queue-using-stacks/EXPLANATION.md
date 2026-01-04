# Implement Queue using Stacks - Explanation

## Problem Summary
Implement a FIFO (First-In-First-Out) queue using only two stacks. The queue should support `push`, `pop`, `peek`, and `empty` operations.

## Approach
This solution uses a **two-stack technique** with lazy transfer:

1. **Two Stacks**: Maintain `in` stack (for push) and `out` stack (for pop/peek)
2. **Push Operation**: Always push to the `in` stack
3. **Pop/Peek Operations**: 
   - If `out` stack is empty, transfer all elements from `in` to `out` (reversing order)
   - Pop/peek from the `out` stack
4. **Empty Check**: Both stacks must be empty

## Why This Works
The key insight is using two stacks to reverse the order twice:
- **Stack reverses order**: LIFO means last pushed is first popped
- **First reversal**: Pushing to `in` stack stores elements in reverse order
- **Second reversal**: Moving from `in` to `out` reverses again, restoring FIFO order
- **Lazy transfer optimization**: We only move elements when `out` is empty, amortizing the cost

**Example flow**:
```
Push 1,2,3:  in=[1,2,3], out=[]
Pop:         Move all → in=[], out=[3,2,1]
             Pop from out → get 1 ✓ (FIFO order)
```

**Why this is efficient**:
- Each element is moved at most once from `in` to `out`
- Elements stay in `out` until consumed
- No need to transfer back to `in`

## Time Complexity
**Amortized O(1)** for all operations
- **Push**: O(1) - simple append to `in` stack
- **Pop/Peek**: Amortized O(1)
  - Worst case: O(n) when transferring all elements from `in` to `out`
  - But each element is transferred at most once
  - Over n operations, total transfer cost is O(n), so amortized O(1) per operation
- **Empty**: O(1) - just check lengths of both stacks

## Space Complexity
**O(n)** where n is the number of elements in the queue
- Two stacks combined store exactly n elements
- No additional data structures
- Each element exists in exactly one stack at any time

## Edge Cases
1. **Empty queue**: Pop/peek from empty queue (handle gracefully or per problem spec)
2. **Single element**: Push then pop should work correctly
3. **Alternating push/pop**: Should maintain FIFO order
4. **Many pushes then many pops**: Amortized efficiency shines here
5. **Peek doesn't remove**: Should be able to peek multiple times

## Worked Example
For operations: `push(1), push(2), peek(), pop(), push(3), pop(), pop()`:

```
Operation: push(1)
in = [1], out = []

Operation: push(2)
in = [1, 2], out = []

Operation: peek()
out is empty, transfer from in:
  Pop 2 from in → push to out
  Pop 1 from in → push to out
in = [], out = [2, 1]
Peek top of out → 1 ✓

Operation: pop()
out = [2, 1] (not empty, no transfer)
Pop from out → 1 ✓
in = [], out = [2]

Operation: push(3)
in = [3], out = [2]

Operation: pop()
out = [2] (not empty, no transfer)
Pop from out → 2 ✓
in = [3], out = []

Operation: pop()
out is empty, transfer from in:
  Pop 3 from in → push to out
in = [], out = [3]
Pop from out → 3 ✓
in = [], out = []
```

## Visual Representation
```
Initial state:
in: []    out: []

After push(1), push(2), push(3):
in: [1→2→3]    out: []
     ↑ bottom    top ↑

First pop() triggers transfer:
Step 1: Pop 3 from in, push to out
in: [1→2]       out: [3]

Step 2: Pop 2 from in, push to out
in: [1]         out: [3→2]

Step 3: Pop 1 from in, push to out
in: []          out: [3→2→1]
                      ↑      ↑
                    bottom  top

Now pop from out:
Pop() → 1 (FIFO order ✓)
in: []          out: [3→2]

Pop() → 2
in: []          out: [3]

Pop() → 3
in: []          out: []
```

## Amortized Analysis
```
Sequence of n operations with k pushes:
- Total push operations: k, each O(1) → O(k)
- Total pop operations: n-k
- Elements transferred: at most k (each pushed element transferred once)
- Total transfer cost: O(k)
- Total cost: O(k) + O(k) + O(n-k) = O(n)
- Average per operation: O(n)/n = O(1) ✓ Amortized constant time
```

## Comparison with Alternatives
```
Approach 1: Transfer on every pop (this solution)
✓ Push: O(1)
✓ Pop: Amortized O(1)
✓ Space: O(n)

Approach 2: Transfer on every push
- Push: O(n) - transfer everything each time
✗ Pop: O(1)
- Worse overall

Approach 3: Single stack + rotation
- Expensive rotations
✗ Not efficient
```

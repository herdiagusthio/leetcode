# Asteroid Collision - Explanation

## Problem Summary
Given an array of integers representing asteroids in a row, where the absolute value represents size and the sign represents direction (positive = right, negative = left), determine the state of asteroids after all collisions. When asteroids collide, the smaller one explodes. If they're equal size, both explode. Asteroids moving in the same direction never collide.

## Approach
This solution uses a **stack-based approach** to simulate collisions:

1. **Initialize Stack**: Use an array `result` to act as a stack for surviving asteroids
2. **Process Each Asteroid**: For each asteroid in the input:
   - Check if collision is possible
   - If no collision: add to stack
   - If collision possible: handle collision logic
3. **Collision Detection**: Collision occurs when:
   - Stack is not empty AND
   - Top of stack is moving right (positive) AND
   - Current asteroid is moving left (negative)
4. **Handle Collisions**: Compare sizes and determine outcome:
   - Equal size: both explode (pop stack, skip current)
   - Right wins: current explodes (skip current)
   - Left wins: pop stack and continue checking against remaining asteroids
5. **Return Result**: The stack contains all surviving asteroids

## Why This Works
The key insight is that asteroids can only collide if they're moving toward each other:
- Right-moving asteroid (in stack) + Left-moving asteroid (current) = collision
- Two right-moving or two left-moving asteroids never collide
- A left-moving asteroid can destroy multiple right-moving asteroids in sequence

**Why stack is the right data structure:**
- We process asteroids left to right
- A left-moving asteroid only collides with asteroids to its left (already in stack)
- We can destroy multiple asteroids by repeatedly popping from the stack
- The stack naturally maintains the relative positions of surviving asteroids

**Why the algorithm is correct:**
- Each asteroid is processed exactly once
- All possible collisions are checked when a left-moving asteroid is encountered
- The inner loop continues until the current asteroid is destroyed or all colliding asteroids are removed
- Asteroids are added to the stack in order, preserving their relative positions

## Time Complexity
**O(n)** where n is the number of asteroids
- Each asteroid is added to the stack at most once: O(n)
- Each asteroid is removed from the stack at most once: O(n)
- Total operations: O(n) + O(n) = O(2n) = O(n)
- Each asteroid is processed at most twice (once when added, once when compared/removed)

## Space Complexity
**O(n)** in worst case
- Stack can hold at most n asteroids (when all move in same direction)
- Best case: O(1) when all asteroids destroy each other
- Average case: O(k) where k is the number of surviving asteroids

## Edge Cases
1. **All moving right**: No collisions, all survive (e.g., `[5, 10, 15]` → `[5, 10, 15]`)
2. **All moving left**: No collisions, all survive (e.g., `[-5, -10, -15]` → `[-5, -10, -15]`)
3. **Equal size collision**: Both explode (e.g., `[8, -8]` → `[]`)
4. **Single asteroid**: Always survives (e.g., `[10]` → `[10]`)
5. **Chain collisions**: One asteroid destroys multiple (e.g., `[1, 2, 3, -10]` → `[-10]`)
6. **No collisions**: Left then right (e.g., `[-5, 10]` → `[-5, 10]`)
7. **Multiple survivors**: Complex patterns (e.g., `[-2, -1, 1, 2]` → `[-2, -1, 1, 2]`)

## Worked Example
For `asteroids = [5, 10, -5]`:

```
Initial: stack = []

Step 1: Process asteroid = 5
  Stack is empty
  Add to stack: [5]

Step 2: Process asteroid = 10
  Stack = [5], top = 5 > 0, asteroid = 10 > 0
  No collision (both moving right)
  Add to stack: [5, 10]

Step 3: Process asteroid = -5
  Stack = [5, 10], top = 10 > 0, asteroid = -5 < 0
  Collision detected! ✓
  
  Inner loop iteration 1:
    top = 10, asteroid = -5
    Compare: 10 vs 5 (absolute values)
    10 > 5 → Right wins, current explodes
    Continue outer_loop (skip adding -5)

Final result: [5, 10] ✓
```

## Complex Example
For `asteroids = [10, 2, -5]`:

```
Initial: stack = []

Step 1: asteroid = 10
  Add to stack: [10]

Step 2: asteroid = 2
  No collision (both moving right)
  Add to stack: [10, 2]

Step 3: asteroid = -5
  Collision with top = 2
  
  Inner loop iteration 1:
    top = 2, asteroid = -5
    2 < 5 → Left wins
    Pop stack: [10]
    Continue inner loop
  
  Inner loop iteration 2:
    top = 10, asteroid = -5
    10 > 5 → Right wins
    Continue outer_loop
    
Final result: [10] ✓
```

## Visual Representation

For `asteroids = [5, 10, -5]`:

```
Initial state:
[5, 10, -5]
 →  →   ←    (directions)

Step 1: Add 5
Stack: [5→]

Step 2: Add 10
Stack: [5→, 10→]

Step 3: -5 approaches from right
Stack: [5→, 10→] ← -5

Collision: 10 vs 5
10 > 5, so 10 wins, -5 explodes

Final: [5→, 10→]
```

For `asteroids = [10, 2, -5]`:

```
State:  [10, 2, -5]
         →  →   ←

After adding 10, 2:
Stack: [10→, 2→]

-5 approaches:
Stack: [10→, 2→] ← -5

Collision 1: 2 vs 5
2 < 5, so -5 wins, 2 explodes
Stack: [10→] ← -5

Collision 2: 10 vs 5
10 > 5, so 10 wins, -5 explodes
Stack: [10→]

Final: [10→]
```

For `asteroids = [8, -8]`:

```
State: [8, -8]
        →  ←

Add 8:
Stack: [8→]

-8 approaches:
Stack: [8→] ← -8

Collision: 8 vs 8
Equal size, both explode
Stack: []

Final: [] (all destroyed)
```

## Algorithm Flow Diagram

```
For each asteroid:
    ┌────────────────────────────────────┐
    │ Can collision happen?              │
    │ - Stack not empty                  │
    │ - Top is positive (moving right)   │
    │ - Current is negative (moving left)│
    └─────────┬────────────┬─────────────┘
              │ NO         │ YES
              │            │
              v            v
         Add to stack   Compare sizes
                            │
              ┌─────────────┼─────────────┐
              │             │             │
              v             v             v
         Equal size    Right wins    Left wins
         (both pop)   (skip current) (pop stack,
                                      continue)
```

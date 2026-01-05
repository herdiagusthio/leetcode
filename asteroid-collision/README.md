# Asteroid Collision

**LeetCode problem:**  
https://leetcode.com/problems/asteroid-collision/

We are given an array `asteroids` of integers representing asteroids in a row. The indices of the asteroid in the array represent their relative position in space.

For each asteroid, the absolute value represents its size, and the sign represents its direction (positive meaning right, negative meaning left). Each asteroid moves at the same speed.

Find out the state of the asteroids after all collisions. If two asteroids meet, the smaller one will explode. If both are the same size, both will explode. Two asteroids moving in the same direction will never meet.

**Example 1:**
```
Input: asteroids = [5,10,-5]
Output: [5,10]
Explanation:
The 10 and -5 collide resulting in 10. The 5 and 10 never collide.
```

**Example 2:**
```
Input: asteroids = [8,-8]
Output: []
Explanation:
The 8 and -8 collide exploding each other.
```

**Example 3:**
```
Input: asteroids = [10,2,-5]
Output: [10]
Explanation:
The 2 and -5 collide resulting in -5. The 10 and -5 collide resulting in 10.
```

**Constraints:**
- `2 <= asteroids.length <= 10^4`
- `-1000 <= asteroids[i] <= 1000`
- `asteroids[i] != 0`

## Solution

The solution uses a stack-based approach to simulate asteroid collisions. For each asteroid, we check if it can collide with asteroids already in the stack (collision only occurs when a right-moving asteroid meets a left-moving one). When collisions occur, we compare sizes and handle three cases: equal size (both explode), right wins (current explodes), or left wins (pop stack and continue checking). This approach efficiently handles chain reactions where one asteroid can destroy multiple others.

**Time Complexity:** O(n) where n is the number of asteroids  
**Space Complexity:** O(n) for the stack

See [EXPLANATION.md](./EXPLANATION.md) for a detailed breakdown of the algorithm.

## How to run

```bash
# Run tests
go test -v

# Run the solution with example inputs
go run .
```

## Files

- `solution.go` - Main solution implementation with stack-based approach
- `solution_test.go` - Table-driven tests covering examples and edge cases
- `EXPLANATION.md` - Detailed explanation of the algorithm and approach
- `README.md` - This file

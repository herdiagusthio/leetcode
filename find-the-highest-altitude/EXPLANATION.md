# Find the Highest Altitude - Explanation

## Problem Summary
A biker starts at altitude 0 and travels through n+1 points. Given an array `gain` where `gain[i]` is the net gain in altitude between points i and i+1, return the highest altitude reached during the trip.

## Approach
This solution uses a simple **iterative approach** with cumulative sum tracking:

1. **Initialize**: Start with altitude 0 and highest altitude 0
2. **Iterate**: For each gain value, add it to the current altitude
3. **Track Maximum**: After each update, check if current altitude exceeds the highest seen so far
4. **Return**: The maximum altitude encountered

## Why This Works
The problem is essentially asking for the maximum prefix sum:
- Starting altitude is 0
- Each gain value changes the altitude cumulatively
- We need to track the maximum value reached at any point
- Since we start at 0, the highest altitude is at least 0 (even if all gains are negative)

## Time Complexity
**O(n)** where n is the length of the gain array
- Single pass through the array
- Each element is processed exactly once
- Constant time operations (addition, comparison) for each element

## Space Complexity
**O(1)** - We only use constant extra space:
- Two integer variables (`highestAltitude` and `altitude`)
- No data structures that grow with input size

## Edge Cases
1. **All negative gains**: Highest altitude remains 0 (starting point)
2. **All positive gains**: Highest altitude is the sum of all gains
3. **Single element**: Highest altitude is max(0, gain[0])
4. **Zero gains**: Highest altitude is 0
5. **Peak at start**: First gain is largest positive value
6. **Peak in middle**: Gains go up then down

## Worked Example
For `gain = [-5,1,5,0,-7]`:

```
Starting altitude: 0
Highest so far: 0

Step 1: gain[0] = -5
altitude = 0 + (-5) = -5
highestAltitude = max(0, -5) = 0

Step 2: gain[1] = 1
altitude = -5 + 1 = -4
highestAltitude = max(0, -4) = 0

Step 3: gain[2] = 5
altitude = -4 + 5 = 1
highestAltitude = max(0, 1) = 1

Step 4: gain[3] = 0
altitude = 1 + 0 = 1
highestAltitude = max(1, 1) = 1

Step 5: gain[4] = -7
altitude = 1 + (-7) = -6
highestAltitude = max(1, -6) = 1
```

The altitudes visited were: [0, -5, -4, 1, 1, -6]

Result: **1**

## Alternative Visualization
```
Point:     0    1    2    3    4    5
Altitude:  0   -5   -4    1    1   -6
           ^                   ^
         start            highest (1)
```

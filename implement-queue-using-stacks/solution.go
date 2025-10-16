package main

import (
	"fmt"
)

// MyQueue implements a FIFO queue using two stacks (in and out).
// Push pushes onto in; Pop/Peek move elements to out when needed.
type MyQueue struct {
	in  []int
	out []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (q *MyQueue) Push(x int) {
	q.in = append(q.in, x)
}

// move transfers elements from in to out if out is empty.
func (q *MyQueue) move() {
	if len(q.out) == 0 {
		for len(q.in) > 0 {
			n := len(q.in)
			v := q.in[n-1]
			q.in = q.in[:n-1]
			q.out = append(q.out, v)
		}
	}
}

func (q *MyQueue) Pop() int {
	q.move()
	if len(q.out) == 0 {
		return 0
	}
	n := len(q.out)
	v := q.out[n-1]
	q.out = q.out[:n-1]
	return v
}

func (q *MyQueue) Peek() int {
	q.move()
	if len(q.out) == 0 {
		return 0
	}
	return q.out[len(q.out)-1]
}

func (q *MyQueue) Empty() bool {
	return len(q.in) == 0 && len(q.out) == 0
}

func main() {
	q := Constructor()
	fmt.Println("Push 1")
	q.Push(1)
	fmt.Println("Push 2")
	q.Push(2)
	fmt.Printf("Peek: %d\n", q.Peek())   // Should print 1
	fmt.Printf("Pop: %d\n", q.Pop())     // Should print 1
	fmt.Printf("Empty: %v\n", q.Empty()) // Should print false
	fmt.Printf("Pop: %d\n", q.Pop())     // Should print 2
	fmt.Printf("Empty: %v\n", q.Empty()) // Should print true
}

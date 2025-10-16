package main

import "testing"

func TestMyQueue_Basic(t *testing.T) {
	q := Constructor()
	if !q.Empty() {
		t.Fatalf("new queue should be empty")
	}

	q.Push(1)
	q.Push(2)
	if q.Empty() {
		t.Fatalf("queue should not be empty after pushes")
	}

	if got := q.Peek(); got != 1 {
		t.Fatalf("peek got=%v want=1", got)
	}

	if got := q.Pop(); got != 1 {
		t.Fatalf("pop got=%v want=1", got)
	}

	if got := q.Empty(); got {
		t.Fatalf("queue should not be empty after one pop")
	}

	if got := q.Pop(); got != 2 {
		t.Fatalf("pop got=%v want=2", got)
	}

	if !q.Empty() {
		t.Fatalf("queue should be empty after popping all elements")
	}
}

func TestMyQueue_EdgeCases(t *testing.T) {
	q := Constructor()

	// multiple pushes and pops to exercise move()
	for i := 0; i < 10; i++ {
		q.Push(i)
	}
	for i := 0; i < 5; i++ {
		if got := q.Pop(); got != i {
			t.Fatalf("pop got=%v want=%v", got, i)
		}
	}

	for i := 10; i < 15; i++ {
		q.Push(i)
	}

	// continue popping remaining
	want := []int{5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
	for _, w := range want {
		if got := q.Pop(); got != w {
			t.Fatalf("pop got=%v want=%v", got, w)
		}
	}

	if !q.Empty() {
		t.Fatalf("queue should be empty after all operations")
	}
}

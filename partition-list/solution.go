package main

import "strconv"

// ListNode is a singly-linked list node used by the tests.
type ListNode struct {
    Val  int
    Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
    lessHead := &ListNode{}
    greaterHead := &ListNode{}
    lessTail, greaterTail := lessHead, greaterHead

    curr := head
    for curr != nil {
        if curr.Val < x {
            lessTail.Next = curr
            lessTail = lessTail.Next
        } else {
            greaterTail.Next = curr
            greaterTail = greaterTail.Next
        }
        curr = curr.Next
    }

    // terminate greater list to avoid cycles
    greaterTail.Next = nil
    // connect less list to greater list
    lessTail.Next = greaterHead.Next
    return lessHead.Next
}

// partitionArray is an alternative implementation that collects nodes into
// two slices (less, greater) and then reconnects them. It uses O(n) extra
// space but may be simpler to reason about.
func partitionArray(head *ListNode, x int) *ListNode {
    if head == nil {
        return nil
    }

    var less []*ListNode
    var greater []*ListNode
    for cur := head; cur != nil; cur = cur.Next {
        if cur.Val < x {
            less = append(less, cur)
        } else {
            greater = append(greater, cur)
        }
    }

    // reconnect
    var dummy ListNode
    tail := &dummy
    for _, n := range less {
        tail.Next = n
        tail = tail.Next
    }
    for _, n := range greater {
        tail.Next = n
        tail = tail.Next
    }
    // terminate to avoid cycles
    tail.Next = nil
    return dummy.Next
}

// helper: build a linked list from slice and return head
func buildList(vals []int) *ListNode {
    dummy := &ListNode{}
    tail := dummy
    for _, v := range vals {
        tail.Next = &ListNode{Val: v}
        tail = tail.Next
    }
    return dummy.Next
}

// helper: convert list to slice
func toSlice(head *ListNode) []int {
    out := []int{}
    for cur := head; cur != nil; cur = cur.Next {
        out = append(out, cur.Val)
    }
    return out
}

func main() {
    // Example 1: [1,4,3,2,5,2], x=3 -> [1,2,2,4,3,5]
    l1 := buildList([]int{1, 4, 3, 2, 5, 2})
    res1 := partition(l1, 3)

    // Example 2: [2,1], x=2 -> [1,2]
    l2 := buildList([]int{2, 1})
    res2 := partition(l2, 2)

    println("example1:", fmtSlice(toSlice(res1)))
    println("example2:", fmtSlice(toSlice(res2)))
}

// fmtSlice prints a slice in the usual Go literal style
func fmtSlice(s []int) string {
    if s == nil {
        return "[]"
    }
    out := "["
    for i, v := range s {
        if i > 0 {
            out += ","
        }
        out += fmtInt(v)
    }
    out += "]"
    return out
}

// small int->string helper (avoid importing fmt for minimal change)
func fmtInt(n int) string {
    // use Sprintf-like logic via strconv
    return itoa(n)
}

// itoa simple implementation using strconv
func itoa(n int) string {
    // import via fully-qualified call to avoid changing imports list at top
    return strconv.Itoa(n)
}
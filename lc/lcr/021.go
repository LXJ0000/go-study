package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	h := &ListNode{
		Val:  -1,
		Next: head,
	}
	l, r := h, h
	for i := 1; i < n; i++ {
		r = r.Next
	}
	for r != nil {
		l = l.Next
		r = r.Next
	}
	l.Next = l.Next.Next
	return h.Next
}

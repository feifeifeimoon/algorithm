package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeDuplicateNodes(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	m := make(map[int]struct{})
	m[head.Val] = struct{}{}

	for cur := head; cur.Next != nil; {
		if _, ok := m[cur.Next.Val]; !ok {
			m[cur.Next.Val] = struct{}{}
			cur = cur.Next
			continue
		}
		cur.Next = cur.Next.Next
	}

	return head
}

func main() {
	l := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  3,
					Next: &ListNode{Val: 3},
				},
			},
		},
	}
	removeDuplicateNodes(l)

	for cur := l; cur != nil; cur = cur.Next {
		fmt.Println(cur.Val)
	}

}

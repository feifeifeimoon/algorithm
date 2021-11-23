package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	if head.Val == val {
		return head.Next
	}

	fast, slow := head.Next, head

	for fast.Val != val {
		fast = fast.Next
		slow = slow.Next
	}

	if fast != nil {
		slow.Next = fast.Next
	}

	return head
}

func main() {
	list := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 5,
			Next: &ListNode{
				Val:  1,
				Next: &ListNode{Val: 9},
			},
		},
	}

	ret := deleteNode(list, 4)

	for ret != nil {
		fmt.Printf("%d->", ret.Val)
		ret = ret.Next
	}
}

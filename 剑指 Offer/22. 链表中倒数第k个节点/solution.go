package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	fast, slow := head, head

	for ; k > 0; k-- {
		fast = fast.Next
	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	return slow
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

	ret := getKthFromEnd(list, 4)

	for ret != nil {
		fmt.Printf("%d->", ret.Val)
		ret = ret.Next
	}
}

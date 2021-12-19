package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func kthToLast(head *ListNode, k int) int {
	fast, slow := head, head

	for i := 0; i < k; i++ {
		fast = fast.Next
	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	return slow.Val
}

func main() {
	l := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: &ListNode{Val: 5},
				},
			},
		},
	}
	fmt.Println(kthToLast(l, 1))
	fmt.Println(kthToLast(l, 2))
	fmt.Println(kthToLast(l, 3))
	fmt.Println(kthToLast(l, 4))
	fmt.Println(kthToLast(l, 5))
}

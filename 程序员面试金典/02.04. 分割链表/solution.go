package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	start, end := head, head
	m := make(map[*ListNode]*ListNode)
	for ; end.Next != nil; end = end.Next {
		m[end.Next] = end
	}

	for start != end {
		if start.Val < x {
			start = start.Next
		} else {
			start.Val, end.Val = end.Val, start.Val
			end = m[end]
		}
	}

	return head
}

func main() {
	l := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  2,
					Next: &ListNode{Val: 5, Next: &ListNode{Val: 2}},
				},
			},
		},
	}
	ret := partition(l, 3)

	for ; ret != nil; ret = ret.Next {
		fmt.Printf("%d ", ret.Val)
	}
	fmt.Println()

	l2 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val:  1,
			Next: nil,
		},
	}

	ret = partition(l2, 2)

	for ; ret != nil; ret = ret.Next {
		fmt.Printf("%d ", ret.Val)
	}
}

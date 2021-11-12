package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
	1.递归法
*/
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	node := reverseList(head.Next)

	head.Next.Next = head
	head.Next = nil

	return node
}

/*
	遍历
*/
func reverseList2(head *ListNode) *ListNode {
	var pre, tmp, newHead *ListNode

	cur := head
	for cur != nil {
		tmp = cur.Next
		cur.Next = pre

		if tmp == nil {
			newHead = cur
		}
		pre = cur
		cur = tmp
	}

	return newHead
}

func main() {

	linkList := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  2,
				Next: nil,
			},
		},
	}
	//fmt.Println(reverseList(linkList))
	fmt.Println(reverseList2(linkList))
	fmt.Println(reverseList(nil))
	fmt.Println(reverseList2(nil))
}

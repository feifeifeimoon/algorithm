package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	cur1, cur2, head := l1, l2, &ListNode{0, nil}

	cur := head
	flag := false
	for ; cur1 != nil && cur2 != nil; cur1, cur2 = cur1.Next, cur2.Next {
		r := cur1.Val + cur2.Val
		if flag {
			r += 1
			flag = false
		}

		if r >= 10 {
			flag = true
			r = r % 10
		}

		cur.Next = &ListNode{Val: r, Next: nil}
		cur = cur.Next
	}

	for ; cur1 != nil; cur1 = cur1.Next {
		r := cur1.Val
		if flag {
			r += 1
			flag = false
			if r >= 10 {
				flag = true
				r = r % 10
			}
		}
		cur.Next = &ListNode{Val: r, Next: nil}
		cur = cur.Next
	}

	for ; cur2 != nil; cur2 = cur2.Next {
		r := cur2.Val
		if flag {
			r += 1
			flag = false
			if r >= 10 {
				flag = true
				r = r % 10
			}
		}
		cur.Next = &ListNode{Val: r, Next: nil}
		cur = cur.Next
	}

	if flag {
		cur.Next = &ListNode{Val: 1, Next: nil}
	}

	return head.Next
}

func main() {
	l1 := &ListNode{Val: 1, Next: &ListNode{Val: 5, Next: &ListNode{Val: 3, Next: &ListNode{Val: 7}}}}
	l2 := &ListNode{Val: 1, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 2, Next: &ListNode{Val: 9}}}}}

	ret := addTwoNumbers(l1, l2)
	for ; ret != nil; ret = ret.Next {
		fmt.Println(ret.Val)
	}
}

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil {
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}

		slow = slow.Next
		if fast == slow {
			break
		}
	}

	if fast == nil {
		return nil
	}

	cur := head
	for ; cur != slow; cur, slow = cur.Next, slow.Next {
	}

	return cur
}

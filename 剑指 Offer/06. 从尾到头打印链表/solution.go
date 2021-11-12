package main

import (
	"container/list"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	l := list.New()
	for head != nil {
		l.PushBack(head.Val)
		head = head.Next
	}

	var ret []int
	for l.Len() != 0 {
		t := l.Back()
		ret = append(ret, t.Value.(int))
		l.Remove(t)
	}

	return ret
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
	fmt.Println(reversePrint(linkList))
	fmt.Println(reversePrint(nil))

}

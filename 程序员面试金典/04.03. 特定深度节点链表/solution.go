package main

import "container/list"

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func listOfDepth(tree *TreeNode) []*ListNode {
	if tree == nil {
		return []*ListNode{}
	}
	var ret []*ListNode

	queue := list.New()
	queue.PushBack(tree)

	for queue.Len() > 0 {
		size := queue.Len()

		head := &ListNode{}
		cur := head
		for ; size > 0; size-- {
			elem := queue.Front()
			node := elem.Value.(*TreeNode)
			cur.Next = &ListNode{Val: node.Val}
			cur = cur.Next

			if node.Left != nil {
				queue.PushBack(node.Left)
			}

			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			queue.Remove(elem)
		}

		ret = append(ret, head.Next)
	}
	return ret
}

func main() {
	listOfDepth(&TreeNode{Val: 1})
}

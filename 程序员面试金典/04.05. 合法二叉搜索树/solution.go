package main

import (
	"container/list"
	"math"
)

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	stack := list.New()
	pre := math.MinInt64
	for stack.Len() > 0 || root != nil {
		for root != nil {
			stack.PushBack(root)
			root = root.Left
		}
		elem := stack.Back()
		node := elem.Value.(*TreeNode)

		if node.Val <= pre {
			return false
		}
		pre = node.Val

		root = node.Right
		stack.Remove(elem)
	}

	return true
}

func main() {
	isValidBST(&TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	})
}

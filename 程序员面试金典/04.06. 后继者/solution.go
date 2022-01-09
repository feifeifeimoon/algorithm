package main

import (
	"container/list"
	"fmt"
	"reflect"
)

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	stack := list.New()
	cur := root
	var pre *TreeNode

	for stack.Len() > 0 || cur != nil {
		for cur != nil {
			stack.PushBack(cur)
			cur = cur.Left
		}

		elem := stack.Back()
		node := elem.Value.(*TreeNode)

		if reflect.DeepEqual(pre, p) {
			return node
		}

		pre = node
		cur = node.Right
		stack.Remove(elem)

	}

	return nil
}

func main() {
	fmt.Println(inorderSuccessor(&TreeNode{
		Val:   2,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 3},
	}, &TreeNode{Val: 1}))
}

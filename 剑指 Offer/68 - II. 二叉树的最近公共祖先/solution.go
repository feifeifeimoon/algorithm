package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	// 当前节点为p 或 q直接返回
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	// 证明 p q一个在左子树 一个在右子树 此时当前root节点就是最近公共祖先
	if left != nil && right != nil {
		return root
	}

	if left != nil {
		return left
	}
	return right
}

func main() {

	tree := &TreeNode{
		Val: 6,
		Left: &TreeNode{
			Val:  2,
			Left: &TreeNode{Val: 0},
			Right: &TreeNode{
				Val:   4,
				Left:  &TreeNode{Val: 3},
				Right: &TreeNode{Val: 5},
			},
		},
		Right: &TreeNode{
			Val:   8,
			Left:  &TreeNode{Val: 7},
			Right: &TreeNode{Val: 9},
		},
	}
	fmt.Println(lowestCommonAncestor(tree, &TreeNode{Val: 9}, &TreeNode{Val: 7}))

}

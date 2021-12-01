package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode, depth int) int {
	if root == nil {
		return depth
	}
	depth++

	l := dfs(root.Left, depth)
	r := dfs(root.Right, depth)

	if l > r {
		return l
	}
	return r
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return dfs(root, 0)
}

func main() {
	tree := &TreeNode{
		Val:  3,
		Left: &TreeNode{Val: 9},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}

	fmt.Println(maxDepth(tree))

}

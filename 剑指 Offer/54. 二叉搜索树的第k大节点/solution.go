package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var g_k int
var ret int

func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	dfs(root.Right)
	if g_k == 0 {
		return
	}
	g_k--
	if g_k == 0 {
		ret = root.Val
	}
	dfs(root.Left)

}

func kthLargest(root *TreeNode, k int) int {
	ret = 0
	g_k = k

	dfs(root)

	return ret
}

func main() {
	tree := &TreeNode{
		Val:   3,
		Left:  &TreeNode{Val: 1, Right: &TreeNode{Val: 2}},
		Right: &TreeNode{Val: 4},
	}

	fmt.Println(kthLargest(tree, 1))
}

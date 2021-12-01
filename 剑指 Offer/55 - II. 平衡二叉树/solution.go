package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l := dfs(root.Left)
	r := dfs(root.Right)

	if (r != -1 && l != -1) && math.Abs(float64(l-r)) <= 1 {
		if l > r {
			return l
		}
		return r
	}
	return -1
}

func isBalanced(root *TreeNode) bool {
	return dfs(root) >= 0
}

func main() {

}

package main

import "math"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func isBalanced(root *TreeNode) bool {

	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 1
		}
		left := dfs(node.Left)
		right := dfs(node.Right)

		if left == -1 || right == -1 || math.Abs(float64(left)-float64(right)) > 1 {
			return -1
		}
		big := left
		if right > big {
			big = right
		}

		return big + 1
	}

	return dfs(root) != -1

}

func main() {
	isBalanced(&TreeNode{
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

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSame(A *TreeNode, B *TreeNode) bool {
	if A == nil && B == nil {
		return true
	}

	if A == nil || B == nil || A.Val != B.Val {
		return false
	}

	return isSame(A.Right, B.Left) && isSame(A.Left, B.Right)

}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSame(root.Left, root.Right)
}

func main() {

}

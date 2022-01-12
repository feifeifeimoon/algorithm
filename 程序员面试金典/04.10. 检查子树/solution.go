package main

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func checkSubTree(t1 *TreeNode, t2 *TreeNode) bool {
	if t2 == nil {
		return true
	}

	if t1 == nil {
		return false
	}

	if t1.Val == t2.Val {
		return checkSubTree(t1.Left, t2.Left) && checkSubTree(t1.Right, t2.Right)
	}

	return checkSubTree(t1.Left, t2) || checkSubTree(t1.Right, t2)
}

func main() {
	checkSubTree(&TreeNode{Val: 1}, &TreeNode{Val: 2})
}

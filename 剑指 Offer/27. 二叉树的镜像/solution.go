package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	right := mirrorTree(root.Right)
	left := mirrorTree(root.Left)
	root.Right = left
	root.Left = right

	return root
}

func main() {

}

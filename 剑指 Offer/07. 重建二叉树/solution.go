package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {

	if len(preorder) == 0 {
		return nil
	}

	node := &TreeNode{Val: preorder[0]}

	index := 0
	for i, v := range inorder {
		if v == preorder[0] {
			index = i
		}
	}

	node.Left = buildTree(preorder[1:index+1], inorder[:index])
	node.Right = buildTree(preorder[index+1:], inorder[index+1:])

	return node
}

func main() {

	//preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]

	preorde := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}

	root := buildTree(preorde, inorder)
	fmt.Println(root)
}

package main

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {

	var build func(begin, end int) *TreeNode
	build = func(begin, end int) *TreeNode {
		if begin >= end {
			return nil
		}
		mid := (end + begin) / 2
		node := &TreeNode{Val: nums[mid]}
		node.Left = build(begin, mid)
		node.Right = build(mid+1, end)
		return node
	}

	return build(0, len(nums))
}

func main() {
	sortedArrayToBST([]int{-10, -3, 0, 5, 9})
}

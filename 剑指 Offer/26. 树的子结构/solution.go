package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func is(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return true
	}

	if A == nil || A.Val != B.Val {
		return false
	}
	return is(A.Left, B.Left) && is(A.Right, B.Right)

}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}

	if is(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B) {
		return true
	}
	return false
}

func main() {
	type args struct {
		A *TreeNode
		B *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test1",
			args{
				A: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val:   4,
						Left:  &TreeNode{Val: 1},
						Right: &TreeNode{Val: 2},
					},
					Right: &TreeNode{Val: 5},
				},
				B: &TreeNode{
					Val:   4,
					Left:  &TreeNode{Val: 1},
					Right: nil,
				},
			},
			true,
		},
	}
	for _, tt := range tests {
		if got := isSubStructure(tt.args.A, tt.args.B); got != tt.want {
			fmt.Printf("isSubStructure() = %v, want %v\n", got, tt.want)
		}
	}
}

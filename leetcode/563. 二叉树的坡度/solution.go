package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ret int

func sum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	sumLeft := sum(root.Left)
	sumRight := sum(root.Right)
	//fmt.Println(sumLeft, sumRight, int(math.Abs(float64(sumLeft - sumRight))))

	ret += int(math.Abs(float64(sumLeft - sumRight)))
	return sumLeft + sumRight + root.Val
}

func findTilt(root *TreeNode) int {
	ret = 0
	if root == nil {
		return 0
	}

	sum(root)
	return ret
}

func main() {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test1",
			args{&TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			}},
			1,
		},
		{
			"test2",
			args{&TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 3},
					Right: &TreeNode{Val: 5},
				},
				Right: &TreeNode{
					Val:   9,
					Left:  nil,
					Right: &TreeNode{Val: 7},
				},
			}},
			15,
		},
	}
	for _, tt := range tests {
		if got := findTilt(tt.args.root); got != tt.want {
			fmt.Printf("isSubStructure() = %v, want %v\n", got, tt.want)
		}
	}
}

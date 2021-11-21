package main

import (
	"fmt"
)

type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) int {
	max := 0
	if root == nil {
		return max
	}

	for _, node := range root.Children {
		d := maxDepth(node)
		if d > max {
			max = d
		}
	}

	return max + 1
}

func main() {
	type args struct {
		root *Node
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test1",
			args{root: &Node{
				Val: 1,
				Children: []*Node{{
					Val:      3,
					Children: []*Node{{Val: 5}, {Val: 6}},
				}, {Val: 2}, {Val: 4}},
			}},
			3,
		},
	}
	for _, tt := range tests {
		if got := maxDepth(tt.args.root); got != tt.want {
			fmt.Printf("maxDepth() = %v, want %v\n", got, tt.want)
		}
	}
}

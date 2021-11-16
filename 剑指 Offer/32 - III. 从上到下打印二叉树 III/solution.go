package main

import (
	"container/list"
	"fmt"
	"reflect"
)

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := list.New()
	queue.PushBack(root)

	var ret [][]int

	flag := true
	size := queue.Len()
	for size > 0 {
		var tmp []int
		for ; size > 0; size-- {
			var item *list.Element
			item = queue.Front()

			cur := item.Value.(*TreeNode)
			tmp = append(tmp, cur.Val)

			if cur.Left != nil {
				queue.PushBack(cur.Left)
			}

			if cur.Right != nil {
				queue.PushBack(cur.Right)
			}
			queue.Remove(item)
		}
		if !flag {
			for i, j := 0, len(tmp)-1; i < j; i, j = i+1, j-1 {
				tmp[i], tmp[j] = tmp[j], tmp[i]
			}
		}
		ret = append(ret, tmp)
		flag = !flag

		size = queue.Len()
	}

	return ret
}

func main() {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"test1",
			args{root: &TreeNode{
				Val:  3,
				Left: &TreeNode{Val: 9},
				Right: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val: 15,
					},
					Right: &TreeNode{Val: 7},
				},
			}},
			[][]int{{3}, {20, 9}, {15, 7}},
		},
	}
	for _, tt := range tests {
		if got := levelOrder(tt.args.root); !reflect.DeepEqual(got, tt.want) {
			fmt.Printf("levelOrder() = %v, want %v\n", got, tt.want)
		}
	}
}

package main

import (
	"fmt"
	"reflect"
)

func findRedundantConnection(edges [][]int) []int {
	parent := make([]int, len(edges)+1)
	for i := range edges {
		parent[i] = i
	}

	var find func(i int) int
	find = func(i int) int {
		// 寻找代表节点 代表节点不是自己 接着找
		if parent[i] != i {
			parent[i] = find(parent[i])
		}
		return parent[i]
	}

	for i := range edges {
		first, second := find(edges[i][0]), find(edges[i][1])
		// 代表节点一样 证明这条线是多余的
		if first == second {
			return []int{edges[i][0], edges[i][1]}
		}
		parent[first] = second

	}

	return []int{}
}

func main() {
	type args struct {
		edges [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"test1",
			args{[][]int{{1, 2}, {1, 3}, {2, 3}}},
			[]int{2, 3},
		}, {
			"test2",
			args{[][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}}},
			[]int{1, 4},
		},
	}
	for _, tt := range tests {
		if got := findRedundantConnection(tt.args.edges); !reflect.DeepEqual(got, tt.want) {
			fmt.Printf("%s findRedundantConnection() = %v, want %v\n", tt.name, got, tt.want)
		}
	}
}

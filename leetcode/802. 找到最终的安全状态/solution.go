package main

import (
	"fmt"
	"reflect"
)

// 三色标记
// 0 - 初始颜色
// 1 - 正常标记队列中
// 2 - 确认是安全节点
func eventualSafeNodes(graph [][]int) []int {
	color := make([]int, len(graph))

	var dfs func(i int) bool
	dfs = func(i int) bool {
		if color[i] != 0 {
			// 如果color不是0 证明在队列中 或者已经标记完
			return color[i] == 2
		}
		color[i] = 1
		for _, v := range graph[i] {
			// 如果有一个节点不是安全节点
			if !dfs(v) {
				return false
			}
		}
		color[i] = 2
		return true
	}

	var ret []int
	for i := 0; i < len(graph); i++ {
		if dfs(i) {
			ret = append(ret, i)
		}
	}
	return ret
}

func main() {
	type args struct {
		graph [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"test1",
			args{[][]int{{1, 2}, {2, 3}, {5}, {0}, {5}, {}, {}}},
			[]int{2, 4, 5, 6},
		}, {
			"test1",
			args{[][]int{{}, {0, 2, 3, 4}, {3}, {4}, {}}},
			[]int{0, 1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		if got := eventualSafeNodes(tt.args.graph); !reflect.DeepEqual(got, tt.want) {
			fmt.Printf("%s eventualSafeNodes() = %v, want %v\n", tt.name, got, tt.want)
		}
	}
}

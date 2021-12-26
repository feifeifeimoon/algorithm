package main

import (
	"container/list"
	"fmt"
)

func findWhetherExistsPath(n int, graph [][]int, start int, target int) bool {
	matrix := make([][]int, n)
	for i := range graph {
		matrix[graph[i][0]] = append(matrix[graph[i][0]], graph[i][1])
	}

	walked := make(map[int]struct{})
	queue := list.New()

	queue.PushBack(start)
	for queue.Len() > 0 {
		elem := queue.Back()
		cur := elem.Value.(int)
		walked[cur] = struct{}{}
		for _, next := range matrix[cur] {
			// 直接返回true
			if next == target {
				return true
			}
			// 已经标记走过了
			if _, ok := walked[next]; !ok {
				walked[next] = struct{}{}
				queue.PushBack(next)
			}
		}
		queue.Remove(elem)
	}

	return false
}

func main() {
	type args struct {
		n      int
		graph  [][]int
		start  int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test1",
			args{3, [][]int{{0, 1}, {0, 2}, {1, 2}, {1, 2}}, 0, 2},
			true,
		},
		{
			"test2",
			args{5, [][]int{{0, 1}, {0, 2}, {0, 1}, {1, 2}, {1, 3}, {1, 3}, {2, 3}}, 0, 4},
			false,
		},
	}
	for _, tt := range tests {
		if got := findWhetherExistsPath(tt.args.n, tt.args.graph, tt.args.start, tt.args.target); got != tt.want {
			fmt.Printf("%s findWhetherExistsPath() = %v, want %v\n", tt.name, got, tt.want)
		}
	}
}

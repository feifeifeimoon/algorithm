package main

import (
	"container/list"
	"fmt"
	"reflect"
)

func loudAndRich(richer [][]int, quiet []int) []int {
	// 统计每个节点入度
	in := make([]int, len(quiet))
	// 邻接矩阵
	matrix := make([][]int, len(quiet))
	for _, v := range richer {
		matrix[v[0]] = append(matrix[v[0]], v[1])
		in[v[1]]++
	}

	queue := list.New()
	// 将入度为0的入队
	for i := range in {
		if in[i] == 0 {
			queue.PushBack(i)
		}
	}

	ret := make([]int, len(quiet))
	for i := range ret {
		ret[i] = i
	}
	for queue.Len() > 0 {
		elem := queue.Front()
		val := elem.Value.(int)

		// 遍历所有 钱少的节点
		for _, v := range matrix[val] {
			if quiet[ret[v]] > quiet[ret[val]] {
				ret[v] = ret[val]
			}
			in[v]--
			if in[v] == 0 {
				queue.PushBack(v)
			}

		}
		queue.Remove(elem)
	}

	return ret
}

func main() {
	type args struct {
		richer [][]int
		quiet  []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"test1",
			args{[][]int{{1, 0}, {2, 1}, {3, 1}, {3, 7}, {4, 3}, {5, 3}, {6, 3}}, []int{3, 2, 5, 4, 6, 1, 7, 0}},
			[]int{5, 5, 2, 5, 4, 5, 6, 7},
		},
	}
	for _, tt := range tests {
		if got := loudAndRich(tt.args.richer, tt.args.quiet); !reflect.DeepEqual(got, tt.want) {
			fmt.Printf("%s loudAndRich() = %v, want %v\n", tt.name, got, tt.want)
		}
	}
}

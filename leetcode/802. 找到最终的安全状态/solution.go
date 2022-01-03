package main

import (
	"container/list"
	"fmt"
	"reflect"
)

func eventualSafeNodes(graph [][]int) []int {
	inDeg := make([]int, len(graph))
	matrix := make([][]int, len(graph))
	for i, v := range graph {
		for _, k := range v {
			matrix[k] = append(matrix[k], i)
		}
		inDeg[i] = len(v)
	}

	queue := list.New()
	for i := range inDeg {
		if inDeg[i] == 0 {
			queue.PushBack(i)
		}
	}

	for queue.Len() > 0 {
		elem := queue.Front()
		val := elem.Value.(int)

		for _, v := range matrix[val] {
			inDeg[v]--
			if inDeg[v] == 0 {
				queue.PushBack(v)
			}
		}

		queue.Remove(elem)
	}

	var ret []int
	for i := range inDeg {
		if inDeg[i] == 0 {
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

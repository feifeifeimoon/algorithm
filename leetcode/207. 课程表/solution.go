package main

import (
	"container/list"
	"fmt"
)

func canFinish(numCourses int, prerequisites [][]int) bool {
	// 记录节点的入度
	indegrees := make([]int, numCourses)
	// 邻接矩阵
	matrix := make([][]int, numCourses)
	for _, v := range prerequisites {
		matrix[v[1]] = append(matrix[v[1]], v[0])
		indegrees[v[0]]++
	}

	queue := list.New()
	// 将入度为0的节点入队
	for i, v := range indegrees {
		if v == 0 {
			queue.PushBack(i)
		}
	}
	count := 0
	for queue.Len() > 0 {
		count++
		elem := queue.Front()
		v := elem.Value.(int)

		for _, c := range matrix[v] {
			indegrees[c]--
			if indegrees[c] == 0 {
				queue.PushBack(c)
			}
		}

		queue.Remove(elem)
	}

	return count == numCourses
}

func main() {
	type args struct {
		numCourses    int
		prerequisites [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		//{
		//	"test1",
		//	args{2, [][]int{{1, 0}}},
		//	true,
		//},
		//{
		//	"test2",
		//	args{2, [][]int{{1, 0}, {0, 1}}},
		//	false,
		//},
		{
			"test3",
			args{4, [][]int{{1, 0}, {2, 3}, {3, 1}, {3, 2}}},
			true,
		},
	}
	for _, tt := range tests {
		if got := canFinish(tt.args.numCourses, tt.args.prerequisites); got != tt.want {
			fmt.Printf("%s canFinish() = %v, want %v\n", tt.name, got, tt.want)
		}
	}
}

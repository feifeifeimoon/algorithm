package main

import (
	"container/list"
	"fmt"
	"reflect"
)

func checkIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {
	inDeg := make([]int, numCourses)
	matrix := make([][]int, numCourses)
	for _, v := range prerequisites {
		matrix[v[0]] = append(matrix[v[0]], v[1])
		inDeg[v[1]]++
	}

	queue := list.New()
	for i, v := range inDeg {
		if v == 0 {
			queue.PushBack(i)
		}
	}

	m := make([]map[int]struct{}, numCourses)
	for i := range m {
		m[i] = make(map[int]struct{})
	}

	for queue.Len() > 0 {
		elem := queue.Front()
		val := elem.Value.(int)

		for _, v := range matrix[val] {

			for k := range m[val] {
				m[v][k] = struct{}{}
			}
			m[v][val] = struct{}{}

			inDeg[v]--
			if inDeg[v] == 0 {
				queue.PushBack(v)
			}

		}

		queue.Remove(elem)
	}

	ret := make([]bool, len(queries))

	for i, v := range queries {
		if _, ok := m[v[1]][v[0]]; ok {
			ret[i] = true
		} else {
			ret[i] = false
		}
	}
	return ret
}

func main() {
	type args struct {
		numCourses             int
		prerequisites, queries [][]int
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		{
			"test1",
			args{2, [][]int{{1, 0}}, [][]int{{0, 1}, {1, 0}}},
			[]bool{false, true},
		},
		{
			"test2",
			args{5, [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}}, [][]int{{0, 4}, {4, 0}, {1, 3}, {3, 0}}},
			[]bool{true, false, true, false},
		},
		{
			"test3",
			args{3, [][]int{{1, 2}, {1, 0}, {2, 0}}, [][]int{{1, 0}, {1, 2}}},
			[]bool{true, true},
		},
	}
	for _, tt := range tests {
		if got := checkIfPrerequisite(tt.args.numCourses, tt.args.prerequisites, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
			fmt.Printf("%s checkIfPrerequisite() = %v, want %v\n", tt.name, got, tt.want)
		}
	}
}

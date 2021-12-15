package main

import (
	"fmt"
	"reflect"
)

func findOrder(numCourses int, prerequisites [][]int) []int {
	m := make([][]int, numCourses)
	s := make([]int, numCourses)
	var ret []int

	for i := range prerequisites {
		m[prerequisites[i][0]] = append(m[prerequisites[i][0]], prerequisites[i][1])
	}

	for {
		count := 0
		for i := range m {
			// 已经学过了
			if s[i] != 0 {
				continue
			}
			//不需要选学
			if len(m[i]) == 0 {
				s[i] = 1
				ret = append(ret, i)
				count++
				continue
			}

			canStudy := true
			for j := range m[i] {
				if s[m[i][j]] == 0 {
					canStudy = false
					break
				}
			}

			if canStudy {
				s[i] = 1
				ret = append(ret, i)
				count++
			}
		}

		if count == 0 {
			break
		}
	}

	if len(ret) == numCourses {
		return ret
	}

	return []int{}
}

func main() {
	type args struct {
		numCourses    int
		prerequisites [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"test1",
			args{2, [][]int{{1, 0}}},
			[]int{0, 1},
		},
		{
			"test2",
			args{4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}},
			[]int{0, 2, 1, 3},
		},
	}
	for _, tt := range tests {
		if got := findOrder(tt.args.numCourses, tt.args.prerequisites); !reflect.DeepEqual(got, tt.want) {
			fmt.Printf("%s findOrder() = %v, want %v\n", tt.name, got, tt.want)
		}
	}
}

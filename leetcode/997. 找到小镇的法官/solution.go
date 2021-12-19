package main

import (
	"fmt"
)

func findJudge(n int, trust [][]int) int {
	matrix := make([][]int, n+1)

	for i := range trust {
		matrix[trust[i][0]] = append(matrix[trust[i][0]], trust[i][1])
	}

	candidate := 0
	// 找出谁都不信任的做为候选人
	for i := 1; i < n+1; i++ {
		if len(matrix[i]) == 0 {
			if candidate != 0 {
				return -1
			}
			candidate = i
		}
	}

	for i := 1; i < n+1; i++ {
		if i == candidate {
			continue
		}

		found := false
		for j := range matrix[i] {
			if candidate == matrix[i][j] {
				found = true
				break
			}
		}

		if !found {
			return -1
		}
	}

	return candidate
}

func main() {
	type args struct {
		n     int
		trust [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test1",
			args{2, [][]int{{1, 2}}},
			2,
		},
		{
			"test2",
			args{3, [][]int{{1, 3}, {2, 3}}},
			3,
		},
		{
			"test3",
			args{3, [][]int{{1, 3}, {2, 3}, {3, 1}}},
			-1,
		},
		{
			"test4",
			args{4, [][]int{{1, 3}, {1, 4}, {2, 3}, {2, 4}, {4, 3}}},
			3,
		},
		{
			"test5",
			args{4, [][]int{{1, 3}, {1, 4}, {2, 3}}},
			-1,
		},
	}
	for _, tt := range tests {
		if got := findJudge(tt.args.n, tt.args.trust); got != tt.want {
			fmt.Printf("%s findJudge() = %v, want %v\n", tt.name, got, tt.want)
		}
	}
}

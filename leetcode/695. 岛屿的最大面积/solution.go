package main

import "fmt"

var dx = []int{1, -1, 0, 0}
var dy = []int{0, 0, 1, -1}

func maxAreaOfIsland(grid [][]int) int {
	row, col := len(grid), len(grid[0])
	ret, temp := 0, 0

	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i >= row || j < 0 || j >= col || grid[i][j] != 1 {
			return
		}
		temp++
		grid[i][j] = 2
		for k := range dx {
			dfs(i+dx[k], j+dy[k])
		}
	}

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 0 {
				continue
			}
			temp = 0
			dfs(i, j)
			if temp > ret {
				ret = temp
			}

		}
	}

	return ret
}

func main() {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test1",
			args{[][]int{{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
				{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0}, {0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
				{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0}}},
			6,
		},
		{
			"test2",
			args{[][]int{{0, 0, 0, 0, 0}}},
			0,
		},
		{
			"test3",
			args{[][]int{{1, 1, 1, 1, 1}}},
			5,
		},
	}
	for _, tt := range tests {
		if got := maxAreaOfIsland(tt.args.grid); got != tt.want {
			fmt.Printf("%s maxAreaOfIsland() = %v, want %v\n", tt.name, got, tt.want)
		}
	}
}

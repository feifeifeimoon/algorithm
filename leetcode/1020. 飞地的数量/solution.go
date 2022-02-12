package main

import "fmt"

var dx = []int{1, -1, 0, 0}
var dy = []int{0, 0, 1, -1}

func numEnclaves(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	visit := make([][]bool, m)
	for i := range visit {
		visit[i] = make([]bool, n)
	}

	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == 0 || visit[i][j] == true {
			return
		}
		visit[i][j] = true

		for k := 0; k < len(dx); k++ {
			dfs(i+dx[k], j+dy[k])
		}
	}

	for i := 0; i < m; i++ {
		dfs(i, 0)
		dfs(i, n-1)
	}
	for j := 0; j < n; j++ {
		dfs(0, j)
		dfs(m-1, j)

	}
	ret := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 && !visit[i][j] {
				ret++
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
			args{[][]int{
				{0, 0, 0, 1, 1, 1, 0, 1, 0, 0},
				{1, 1, 0, 0, 0, 1, 0, 1, 1, 1},
				{0, 0, 0, 1, 1, 1, 0, 1, 0, 0},
				{0, 1, 1, 0, 0, 0, 1, 0, 1, 0},
				{0, 1, 1, 1, 1, 1, 0, 0, 1, 0},
				{0, 0, 1, 0, 1, 1, 1, 1, 0, 1},
				{0, 1, 1, 0, 0, 0, 1, 1, 1, 1},
				{0, 0, 1, 0, 0, 1, 0, 1, 0, 1},
				{1, 0, 1, 0, 1, 1, 0, 0, 0, 0},
				{0, 0, 0, 0, 1, 1, 0, 0, 0, 1}}},
			3,
		},
		{
			"test2",
			args{[][]int{
				{0, 0, 0, 0, 0},
				{1, 1, 1, 0, 0},
				{0, 0, 1, 0, 0},
				{0, 1, 1, 0, 0},
				{0, 0, 0, 1, 0},
				{0, 0, 0, 0, 0}}},
			1,
		},
		//{
		//	"test3",
		//	args{[][]int{{1, 1, 1, 1, 1}}},
		//	5,
		//},
	}
	for _, tt := range tests {
		if got := numEnclaves(tt.args.grid); got != tt.want {
			fmt.Printf("%s numEnclaves() = %v, want %v\n", tt.name, got, tt.want)
		}
	}
}

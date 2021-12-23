package main

import (
	"fmt"
)

var dx = []int{1, -1, 0, 0}
var dy = []int{0, 0, 1, -1}

// calArea 计算每一块岛屿的大小 并且将值修改为index
func calArea(grid [][]int, i, j, index int) int {
	area, row, col := 0, len(grid), len(grid[0])

	var dfs func(i, j int)
	dfs = func(i, j int) {
		// 越界或者为0 退出
		if i < 0 || i >= row || j < 0 || j >= col || grid[i][j] != 1 {
			return
		}
		area++
		grid[i][j] = index
		for k := range dx {
			dfs(i+dx[k], j+dy[k])
		}
	}
	dfs(i, j)
	return area
}

func largestIsland(grid [][]int) int {
	row, col := len(grid), len(grid[0])

	// 将每块岛屿编号 0,1 不用 从2开始
	area := make([]int, row*col+2)
	cur := 2
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				area[cur] = calArea(grid, i, j, cur)
				cur++
			}
		}
	}

	// 合并岛屿大小，判断将i,j变为1后 能连接上多少岛屿
	mergeArea := func(i, j int) int {
		ret := 1
		// 使用set去重
		set := make(map[int]struct{})

		for k := range dx {
			c, r := i+dx[k], j+dy[k]
			if c < 0 || c >= row || r < 0 || r >= col {
				continue
			}
			set[grid[c][r]] = struct{}{}
		}

		for k := range set {
			ret += area[k]
		}
		return ret
	}

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] != 0 {
				continue
			}
			// merge area
			area[cur] = mergeArea(i, j)
			cur++
		}
	}

	ret := area[2]
	for i := 2; i < len(area) && area[i] != 0; i++ {
		if area[i] > ret {
			ret = area[i]
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
			args{[][]int{{1, 0}, {0, 1}}},
			3,
		},
		{
			"test2",
			args{[][]int{{1, 1}, {1, 0}}},
			4,
		},
		{
			"test3",
			args{[][]int{{1, 1}, {1, 1}}},
			4,
		},
	}
	for _, tt := range tests {
		if got := largestIsland(tt.args.grid); got != tt.want {
			fmt.Printf("%s largestIsland() = %v, want %v\n", tt.name, got, tt.want)
		}
	}
}

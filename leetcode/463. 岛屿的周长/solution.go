package main

import (
	"fmt"
)

var dx = []int{1, -1, 0, 0}
var dy = []int{0, 0, 1, -1}

func islandPerimeter(grid [][]int) int {
	row, col := len(grid), len(grid[0])
	count := 0

	cal := func(i, j int) int {
		if i < 0 || i >= row || j < 0 || j >= col || grid[i][j] == 0 {
			return 0
		}
		return -1
	}

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 0 {
				continue
			}
			max := 4
			for k := range dx {
				max += cal(i+dx[k], j+dy[k])
			}
			count += max
		}
	}
	return count
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
			args{[][]int{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}}},
			16,
		},
		{
			"test2",
			args{[][]int{{1}}},
			4,
		},
		{
			"test3",
			args{[][]int{{1, 0}}},
			4,
		},
	}
	for _, tt := range tests {
		if got := islandPerimeter(tt.args.grid); got != tt.want {
			fmt.Printf("%s islandPerimeter() = %v, want %v\n", tt.name, got, tt.want)
		}
	}
}

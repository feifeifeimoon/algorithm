package main

import (
	"fmt"
)

/*
	从左下角或者右上角开始即可
*/
func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])

	x := 0
	y := m - 1

	for x < n && y >= 0 {
		if matrix[y][x] > target {
			y--
		} else if matrix[y][x] < target {
			x++
		} else {
			return true
		}

	}

	return false
}

func main() {
	type args struct {
		matrix [][]int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test1",
			args{
				matrix: [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}},
				target: 5,
			},
			true,
		},
		{
			"test2",
			args{
				matrix: [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}},
				target: 20,
			},
			false,
		},
		{
			"test3",
			args{
				matrix: [][]int{{1, 1}},
				target: 2,
			},
			false,
		},
	}
	for _, tt := range tests {
		if got := searchMatrix(tt.args.matrix, tt.args.target); got != tt.want {
			fmt.Printf("%s searchMatrix() = %v, want %v\n", tt.name, got, tt.want)
		}
	}
}

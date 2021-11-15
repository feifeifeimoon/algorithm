package main

import (
	"fmt"
)

func findNumberIn2DArray(matrix [][]int, target int) bool {
	row, col := len(matrix)-1, 0

	for row >= 0 && col < len(matrix[0]) {
		if matrix[row][col] < target {
			col++
		} else if matrix[row][col] > target {
			row--
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
				matrix: [][]int{
					{1, 4, 7, 11, 15},
					{2, 5, 8, 12, 19},
					{3, 6, 9, 16, 22},
					{10, 13, 14, 17, 24},
					{18, 21, 23, 26, 30},
				},
				target: 5,
			},
			true,
		},
		{
			"test2",
			args{
				matrix: [][]int{
					{1, 4, 7, 11, 15},
					{2, 5, 8, 12, 19},
					{3, 6, 9, 16, 22},
					{10, 13, 14, 17, 24},
					{18, 21, 23, 26, 30},
				},
				target: 20,
			},
			false,
		},
		{
			"test3",
			args{
				matrix: [][]int{},
				target: 20,
			},
			false,
		},
	}
	for _, tt := range tests {
		if got := findNumberIn2DArray(tt.args.matrix, tt.args.target); got != tt.want {
			fmt.Printf("findNumberIn2DArray() = %v, want %v\n", got, tt.want)
		}
	}
}

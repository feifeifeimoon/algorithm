package main

import (
	"fmt"
	"reflect"
)

func setZeroes(matrix [][]int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}

	row, col := make([]int, len(matrix)), make([]int, len(matrix[0]))

	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 0 {
				row[i], col[j] = 1, 1
			}
		}
	}

	for i := range matrix {
		for j := range matrix[i] {
			if row[i] == 1 || col[j] == 1 {
				matrix[i][j] = 0
			}
		}
	}

}

func main() {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"test1",
			args{[][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}},
			[][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}},
		},
	}
	for _, tt := range tests {
		if setZeroes(tt.args.matrix); !reflect.DeepEqual(tt.args.matrix, tt.want) {
			fmt.Printf("%s setZeroes() = %v, want %v\n", tt.name, tt.args.matrix, tt.want)
		}
	}
}

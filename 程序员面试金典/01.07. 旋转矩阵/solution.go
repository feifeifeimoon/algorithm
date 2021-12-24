package main

import (
	"fmt"
	"reflect"
)

func rotate(matrix [][]int) {
	newMatrix := make([][]int, len(matrix))
	n := len(matrix[0])
	for i := range newMatrix {
		newMatrix[i] = make([]int, n)
	}
	for i := range matrix {
		for j := range matrix[i] {
			newMatrix[j][n-i-1] = matrix[i][j]

		}
	}
	copy(matrix, newMatrix)
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
			args{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}},
			[][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}},
		},
	}
	for _, tt := range tests {
		if rotate(tt.args.matrix); !reflect.DeepEqual(tt.args.matrix, tt.want) {
			fmt.Printf("%s rotate() = %v, want %v\n", tt.name, tt.args.matrix, tt.want)
		}
	}
}

package main

import (
	"fmt"
	"reflect"
)

func luckyNumbers(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	var ret []int

	t := make(map[int]struct{})

	for i := 0; i < m; i++ {
		min := matrix[i][0]
		for j := 1; j < n; j++ {
			if matrix[i][j] < min {
				min = matrix[i][j]
			}
		}
		t[min] = struct{}{}
	}

	for i := 0; i < n; i++ {
		max := matrix[0][i]
		for j := 1; j < m; j++ {
			if matrix[j][i] > max {
				max = matrix[j][i]
			}
		}

		if _, ok := t[max]; ok {
			ret = append(ret, max)
		}
	}

	return ret
}

func main() {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"test1",
			args{matrix: [][]int{{3, 7, 8}, {9, 11, 13}, {15, 16, 17}}},
			[]int{15},
		},
	}
	for _, tt := range tests {
		if got := luckyNumbers(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
			fmt.Printf("%s luckyNumbers() = %v, want %v\n", tt.name, got, tt.want)
		}
	}
}

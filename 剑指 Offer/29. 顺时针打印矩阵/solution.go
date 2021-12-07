package main

import (
	"fmt"
	"reflect"
)

func printOnce(matrix [][]int, row1, col1, row2, col2 int) []int {
	var ret []int

	//  只用打印一行
	if row1 == row2 {
		for cur := col1; cur <= col2; cur++ {
			ret = append(ret, matrix[row1][cur])
		}
		return ret
	} else if col1 == col2 {
		// 只用打印一列
		for cur := row1; cur <= row2; cur++ {
			ret = append(ret, matrix[cur][col1])
		}
		return ret
	}

	// 打印上边一行
	for cur := col1; cur < col2; cur++ {
		ret = append(ret, matrix[row1][cur])
	}

	// 打印右边一列
	for cur := row1; cur < row2; cur++ {
		ret = append(ret, matrix[cur][col2])
	}

	// 打印下边一行
	for cur := col2; cur > col1; cur-- {
		ret = append(ret, matrix[row2][cur])
	}

	// 打印左边一列
	for cur := row2; cur > row1; cur-- {
		ret = append(ret, matrix[cur][col1])
	}

	return ret
}

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	row1, col1 := 0, 0
	row2, col2 := len(matrix)-1, len(matrix[0])-1

	var ret []int

	for row1 <= row2 && col1 <= col2 {
		ret = append(ret, printOnce(matrix, row1, col1, row2, col2)...)
		row1++
		col1++
		row2--
		col2--
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
			args{[][]int{
				{1, 2, 3}, {4, 5, 6}, {7, 8, 9},
			}},
			[]int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			"test2",
			args{[][]int{
				{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12},
			}},
			[]int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
	}
	for _, tt := range tests {
		if got := spiralOrder(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
			fmt.Printf("spiralOrder() = %v, want %v\n", got, tt.want)
		}
	}

}

package main

import (
	"fmt"
	"reflect"
)

func construct2DArray(original []int, m int, n int) [][]int {
	if len(original) != m*n {
		return [][]int{}
	}
	cur := 0
	ret := make([][]int, m)
	for i := range ret {
		ret[i] = append(ret[i], original[cur:cur+n]...)
		cur += n
	}

	return ret
}

func main() {
	type args struct {
		original []int
		m, n     int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"test1",
			args{[]int{1, 2, 3, 4}, 1, 4},
			[][]int{{1, 2, 3, 4}},
		},
	}
	for _, tt := range tests {
		if got := construct2DArray(tt.args.original, tt.args.m, tt.args.n); !reflect.DeepEqual(got, tt.want) {
			fmt.Printf("%s construct2DArray() = %v, want %v\n", tt.name, got, tt.want)
		}
	}
}

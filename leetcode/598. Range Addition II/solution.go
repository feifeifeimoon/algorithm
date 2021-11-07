package main

import (
	"fmt"
)

func maxCount(m int, n int, ops [][]int) int {

	minX, minY := m, n

	for _, op := range ops {

		if op[0] < minX {
			minX = op[0]
		}

		if op[1] < minY {
			minY = op[1]
		}
	}

	return minX * minY
}

func main() {
	type args struct {
		m   int
		n   int
		ops [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test1",
			args{3, 3, [][]int{{2, 2}, {3, 3}}},
			4,
		},
		{
			"test2",
			args{3, 3, [][]int{}},
			9,
		},
	}
	for _, tt := range tests {
		if got := maxCount(tt.args.m, tt.args.n, tt.args.ops); got != tt.want {
			fmt.Printf("maxCount() = %v, want %v\n", got, tt.want)
		}
	}

}

package main

import (
	"fmt"
)

func findRepeatNumber(nums []int) int {
	m := make(map[int]struct{})

	for _, num := range nums {
		if _, ok := m[num]; ok {
			return num
		}
		m[num] = struct{}{}

	}
	return -1
}

func main() {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test1",
			args{[]int{2, 3, 1, 0, 2, 5, 3}},
			2,
		},
		{
			"test2",
			args{[]int{2, 3, 1, 3, 2, 5, 3}},
			3,
		},
	}
	for _, tt := range tests {
		if got := findRepeatNumber(tt.args.nums); got != tt.want {
			fmt.Printf("findRepeatNumber() = %v, want %v\n", got, tt.want)
		}
	}
}

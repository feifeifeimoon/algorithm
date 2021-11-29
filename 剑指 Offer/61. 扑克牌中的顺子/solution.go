package main

import (
	"fmt"
)

func isStraight(nums []int) bool {

	set := make(map[int]struct{})

	max, min := 0, 14
	for _, num := range nums {
		if num == 0 {
			continue
		}
		if _, ok := set[num]; ok {
			return false
		}

		set[num] = struct{}{}

		if num > max {
			max = num
		}
		if num < min {
			min = num
		}

	}

	return !(max-min >= 5)
}

func main() {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test1",
			args{nums: []int{1, 2, 3, 4, 5}},
			true,
		},
		{
			"test2",
			args{[]int{0, 0, 1, 2, 5}},
			true,
		},
		{
			"test3",
			args{[]int{1, 6, 5, 4, 2}},
			false,
		},
	}
	for _, tt := range tests {
		if got := isStraight(tt.args.nums); got != tt.want {
			fmt.Printf("isStraight() = %v, want %v\n", got, tt.want)
		}

	}
}

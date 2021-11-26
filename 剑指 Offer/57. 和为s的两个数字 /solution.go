package main

import (
	"fmt"
	"reflect"
)

func twoSum(nums []int, target int) []int {
	start, end := 0, len(nums)-1

	for start < end {
		count := nums[start] + nums[end]
		if count < target {
			start++
		} else if count > target {
			end--
		} else {
			return []int{nums[start], nums[end]}
		}
	}
	return []int{}
}

func main() {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"test1",
			args{nums: []int{2, 7, 11, 15}, target: 9},
			[]int{2, 7},
		},
		{
			"test2",
			args{nums: []int{2, 7, 11, 15}, target: 18},
			[]int{7, 11},
		},
		{
			"test3",
			args{nums: []int{10, 26, 30, 31, 47, 60}, target: 40},
			[]int{10, 30},
		},
	}
	for _, tt := range tests {
		if got := twoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
			fmt.Printf("twoSum() = %v, want %v\n", got, tt.want)
		}
	}
}

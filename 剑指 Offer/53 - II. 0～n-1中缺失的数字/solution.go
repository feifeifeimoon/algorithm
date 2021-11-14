package main

import (
	"fmt"
)

func missingNumber(nums []int) int {
	begin, end := 0, len(nums)-1

	for begin <= end {
		mid := begin + (end-begin)/2

		if nums[mid] != mid {
			end = mid - 1
		} else {
			begin = mid + 1
		}
	}

	return begin
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
			args{
				nums: []int{0, 1, 3},
			},
			2,
		},
		{
			"test2",
			args{
				nums: []int{0, 1, 2, 3, 4, 5, 6, 7, 9},
			},
			8,
		},
	}
	for _, tt := range tests {
		if got := missingNumber(tt.args.nums); got != tt.want {
			fmt.Printf("missingNumber() = %v, want %v\n", got, tt.want)
		}
	}
}

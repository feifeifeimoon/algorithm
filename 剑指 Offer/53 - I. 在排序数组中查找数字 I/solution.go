package main

import (
	"fmt"
)

func search(nums []int, target int) int {
	begin, end, index := 0, len(nums)-1, -1

	for begin <= end {
		mid := begin + (end-begin)/2

		if nums[mid] < target {
			begin = mid + 1
		} else if nums[mid] > target {
			end = mid - 1
		} else {
			index = mid
			break
		}
	}

	if index == -1 {
		return 0
	}

	count := 0
	for i := index; i <= len(nums)-1 && nums[i] == target; i++ {
		count++
	}
	for i := index - 1; i >= 0 && nums[i] == target; i-- {
		count++
	}

	return count
}

func main() {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test1",
			args{
				nums:   []int{5, 7, 7, 8, 8, 10},
				target: 8,
			},
			2,
		},
		{
			"test2",
			args{
				nums:   []int{5, 7, 7, 8, 8, 10},
				target: 6,
			},
			0,
		},
		{
			"test3",
			args{
				nums:   []int{1},
				target: 1,
			},
			1,
		},
	}
	for _, tt := range tests {
		if got := search(tt.args.nums, tt.args.target); got != tt.want {
			fmt.Printf("search() = %v, want %v\n", got, tt.want)
		}
	}
}

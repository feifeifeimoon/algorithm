package main

import (
	"fmt"
)

func majorityElement(nums []int) int {
	cur, count := nums[0], 1

	for i := 1; i < len(nums); i++ {
		if nums[i] == cur {
			count++
		} else {
			count--
			if count == 0 {
				cur = nums[i]
				count = 1
			}
		}
	}

	return cur
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
			args{[]int{1, 2, 3, 2, 2, 2, 5, 4, 2}},
			2,
		},
	}
	for _, tt := range tests {
		if got := majorityElement(tt.args.nums); tt.want != got {
			fmt.Printf("majorityElement() = %v, want %v", got, tt.want)
		}
	}
}

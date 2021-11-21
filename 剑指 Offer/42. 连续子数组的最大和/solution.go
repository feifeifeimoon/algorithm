package main

import (
	"fmt"
)

func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		if dp[i-1]+nums[i] > nums[i] {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]
		}
	}

	max := dp[0]
	for _, v := range dp {
		if v > max {
			max = v
		}
	}

	return max
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
			args{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}},
			6,
		},
	}
	for _, tt := range tests {
		if got := maxSubArray(tt.args.nums); got != tt.want {
			fmt.Printf("maxSubArray() = %v, want %v", got, tt.want)
		}
	}
}

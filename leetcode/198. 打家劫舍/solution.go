package main

import "fmt"

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	dp := make([]int, len(nums))
	dp[0], dp[1] = nums[0], nums[1]
	if nums[1] < nums[0] {
		dp[1] = nums[0]
	}

	for i := 2; i < len(nums); i++ {
		if dp[i-2]+nums[i] > dp[i-1] {
			dp[i] = dp[i-2] + nums[i]
		} else {
			dp[i] = dp[i-1]
		}
	}

	return dp[len(nums)-1]
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
			args{[]int{1, 2, 3, 1}},
			4,
		},
		{
			"test2",
			args{[]int{2, 7, 9, 3, 1}},
			12,
		},
	}
	for _, tt := range tests {
		if got := rob(tt.args.nums); got != tt.want {
			fmt.Printf("%s rob() = %v, want %v\n", tt.name, got, tt.want)
		}
	}
}

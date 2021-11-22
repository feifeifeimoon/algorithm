package main

import (
	"fmt"
	"strconv"
)

func translateNum(num int) int {
	if num < 10 {
		return 1
	}
	src := strconv.Itoa(num)

	dp := make([]int, len(src)+1)
	dp[0], dp[1] = 1, 1

	//if pre := src[0:2]; pre <= "25" && pre >= "10" {
	//	dp[1] += 1
	//}

	for i := 2; i <= len(src); i++ {
		dp[i] += dp[i-1]

		pre := src[i-2 : i]
		if pre <= "25" && pre >= "10" {
			dp[i] += dp[i-2]
		}

	}

	for i := range dp {
		fmt.Printf("%d ", dp[i])
	}
	fmt.Println()

	return dp[len(dp)-1]
}

func main() {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test1",
			args{12258},
			5,
		},
		{
			"test2",
			args{20},
			2,
		},
		{
			"test3",
			args{50},
			1,
		},
		{
			"test4",
			args{9},
			1,
		},
	}
	for _, tt := range tests {
		if got := translateNum(tt.args.num); got != tt.want {
			fmt.Printf("translateNum() = %v, want %v\n", got, tt.want)
		}
	}
}

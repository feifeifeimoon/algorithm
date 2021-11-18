package main

import "fmt"

// todo DP
func maxProfit(prices []int) int {
	min, profit := prices[0], 0

	for _, v := range prices {
		if v < min {
			min = v
			continue
		}

		if v-min > profit {
			profit = v - min
		}

	}
	return profit
}

func main() {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test1",
			args{[]int{7, 1, 5, 3, 6, 4}},
			5,
		},
		{
			"test2",
			args{[]int{7, 6, 4, 3, 1}},
			0,
		},
	}
	for _, tt := range tests {
		if got := maxProfit(tt.args.prices); got != tt.want {
			fmt.Printf("fib() = %v, want %v\n", got, tt.want)
		}
	}
}

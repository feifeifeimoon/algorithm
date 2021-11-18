package main

import "fmt"

func numWays(n int) int {
	const mod int = 1e9 + 7
	if n <= 1 {
		return 1
	}
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1

	for i := 2; i <= n; i++ {
		dp[i] = (dp[i-1] + dp[i-2]) % mod
	}
	return dp[n]
}

func main() {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test1",
			args{2},
			2,
		},
		{
			"test2",
			args{7},
			21,
		},
		{
			"test3",
			args{0},
			1,
		},
	}
	for _, tt := range tests {
		if got := numWays(tt.args.n); got != tt.want {
			fmt.Printf("fib() = %v, want %v\n", got, tt.want)
		}
	}
}

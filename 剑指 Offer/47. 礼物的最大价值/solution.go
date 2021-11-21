package main

import "fmt"

func maxValue(grid [][]int) int {

	dp := make([][]int, len(grid))
	for i := range dp {
		dp[i] = make([]int, len(grid[0]))
	}

	dp[0][0] = grid[0][0]
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {

			if i == 0 && j == 0 {
				continue
			}

			if i == 0 {
				dp[i][j] = dp[i][j-1] + grid[i][j]
			} else if j == 0 {
				dp[i][j] = dp[i-1][j] + grid[i][j]
			} else {
				if dp[i][j-1] > dp[i-1][j] {
					dp[i][j] = dp[i][j-1] + grid[i][j]
				} else {
					dp[i][j] = dp[i-1][j] + grid[i][j]
				}
			}

		}
	}

	return dp[len(grid)-1][len(grid[0])-1]

}

func main() {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test1",
			args{[][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}},
			12,
		},
	}
	for _, tt := range tests {
		if got := maxValue(tt.args.grid); got != tt.want {
			fmt.Printf("maxValue() = %v, want %v", got, tt.want)
		}
	}
}

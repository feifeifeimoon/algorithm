package main

import (
	"fmt"
)

func maxPower(s string) int {
	if len(s) == 0 {
		return 0
	}
	count, max := 1, 1

	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			count++
			if count > max {
				max = count
			}
		} else {
			count = 1
		}
	}

	return max
}

func main() {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test1",
			args{"leetcode"},
			2,
		},
		{
			"test2",
			args{"abbcccddddeeeeedcba"},
			5,
		},
		{
			"test3",
			args{"triplepillooooow"},
			5,
		},
		{
			"test4",
			args{"hooraaaaaaaaaaay"},
			11,
		},
		{
			"test5",
			args{"tourist"},
			1,
		},
	}
	for _, tt := range tests {
		if got := maxPower(tt.args.s); got != tt.want {
			fmt.Printf("maxPower() = %v, want %v\n", got, tt.want)
		}
	}
}

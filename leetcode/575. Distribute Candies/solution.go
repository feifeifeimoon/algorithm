package main

import (
	"fmt"
)

func distributeCandies(candyType []int) int {

	set := make(map[int]struct{})

	for _, t := range candyType {
		set[t] = struct{}{}
	}

	if len(set) > len(candyType)/2 {
		return len(candyType) / 2
	}

	return len(set)
}

func main() {
	type args struct {
		candyType []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test1",
			args{candyType: []int{1, 1, 2, 2, 3, 3}},
			3,
		},
		{
			"test2",
			args{candyType: []int{1, 1, 2, 3}},
			2,
		},
		{
			"test2",
			args{candyType: []int{6, 6, 6, 6}},
			1,
		},
	}
	for _, tt := range tests {
		if got := distributeCandies(tt.args.candyType); got != tt.want {
			fmt.Printf("distributeCandies() = %v, want %v\n", got, tt.want)
		}
	}
}

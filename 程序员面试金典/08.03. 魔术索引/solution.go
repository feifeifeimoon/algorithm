package main

import (
	"fmt"
)

func findMagicIndex(nums []int) int {
	for i := 0; i < len(nums); i++ {
		if i == nums[i] {
			return i
		}
	}
	return -1
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
			args{nums: []int{0, 2, 3, 4, 5}},
			0,
		},
	}
	for _, tt := range tests {
		if got := findMagicIndex(tt.args.nums); got != tt.want {
			fmt.Printf("%s findMagicIndex() = %v, want %v\n", tt.name, got, tt.want)
		}
	}
}

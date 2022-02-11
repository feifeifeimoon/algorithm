package main

import (
	"fmt"
	"math"
	"reflect"
	"sort"
)

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func minimumDifference(nums []int, k int) int {
	sort.Ints(nums)
	ans := math.MaxInt32
	for i, num := range nums[:len(nums)-k+1] {
		ans = min(ans, nums[i+k-1]-num)
	}
	return ans
}

func main() {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test1",
			args{[]int{9, 4, 1, 7}, 2},
			2,
		},
	}
	for _, tt := range tests {
		if got := minimumDifference(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
			fmt.Printf("%s minimumDifference() = %v, want %v\n", tt.name, got, tt.want)
		}
	}
}

package main

import (
	"fmt"
	"reflect"
)

func singleNumbers(nums []int) []int {
	temp, ret1, ret2 := 0, 0, 0
	for i := range nums {
		temp ^= nums[i]
	}

	b := 0
	for temp != 1 {
		temp = temp >> 1
		b++
	}

	b = 1 << (b)

	for i := range nums {

		if nums[i]&b == 0 {
			ret1 ^= nums[i]
		} else {
			ret2 ^= nums[i]
		}
	}

	return []int{ret1, ret2}
}

func main() {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"test1",
			args{[]int{4, 1, 4, 6}},
			[]int{1, 6},
		},
	}
	for _, tt := range tests {
		if got := singleNumbers(tt.args.nums); !reflect.DeepEqual(tt.want, got) {
			fmt.Printf("singleNumbers() = %v, want %v", got, tt.want)
		}
	}
}

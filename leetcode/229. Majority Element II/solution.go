package main

import (
	"fmt"
	"reflect"
)

func majorityElement(nums []int) []int {
	var first, second, countFirst, countSecond int

	for i := range nums {
		if (countFirst == 0 || nums[i] == first) && nums[i] != second {
			first = nums[i]
			countFirst++
		} else if countSecond == 0 || nums[i] == second {
			second = nums[i]
			countSecond++
		} else {
			countFirst--
			countSecond--
		}

	}

	countFirst, countSecond = 0, 0
	for i := range nums {
		if nums[i] == first {
			countFirst++
		} else if nums[i] == second {
			countSecond++
		}
	}

	var ret []int

	if countFirst > len(nums)/3 {
		ret = append(ret, first)
	}

	if countSecond > len(nums)/3 {
		ret = append(ret, second)
	}

	return ret
}

func main() {
	type args struct {
		digits []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"test1",
			args{digits: []int{3, 2, 3}},
			[]int{3},
		},
		{
			"test2",
			args{digits: []int{1, 1, 1, 3, 3, 2, 2, 2}},
			[]int{1, 2},
		},
		{
			"test3",
			args{digits: []int{1, 2, 3}},
			[]int{},
		},
		{
			"test4",
			args{digits: []int{1}},
			[]int{1},
		},
		{
			"test4",
			args{digits: []int{6, 5, 5}},
			[]int{5},
		},
		{
			"test5",
			args{digits: []int{2, 1, 1, 3, 1, 4, 5, 6}},
			[]int{1},
		},
	}
	for _, tt := range tests {
		if got := majorityElement(tt.args.digits); !reflect.DeepEqual(got, tt.want) {
			fmt.Printf("majorityElement() = %v, want %v\n", got, tt.want)
		}
	}
}

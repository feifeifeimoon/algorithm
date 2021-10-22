package main

import (
	"fmt"
	"reflect"
)

func majorityElement(nums []int) []int {
	var first, second, countFirst, countSecond int

	for i := range nums {
		if countFirst == 0 || nums[i] == first {
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

	//fmt.Println(first, countFirst, second, countSecond)

	countFirst, countSecond = 0, 0
	for i := range nums {
		if nums[i] == first {
			countFirst++
		} else if nums[i] == second {
			countSecond++
		}
	}

	if countSecond > len(nums)/3 {
		return []int{first, second}
	}

	if countFirst > len(nums)/3 {
		return []int{first}
	}

	return []int{}
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
	}
	for _, tt := range tests {
		if got := majorityElement(tt.args.digits); !reflect.DeepEqual(got, tt.want) {
			fmt.Printf("majorityElement() = %v, want %v\n", got, tt.want)
		}
	}
}

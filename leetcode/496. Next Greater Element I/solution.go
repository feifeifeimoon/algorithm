package main

import (
	"fmt"
	"reflect"
)

/*
	第一种：暴力 两层for循环遍历查找 时间复杂度 O(n * n)
	todo() 第二种：关键字 单调栈
*/
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	var ret []int

	for i := range nums1 {
		found, flag := false, false

		for j := range nums2 {
			if nums2[j] == nums1[i] {
				found = true
				continue
			}

			if found && nums2[j] > nums1[i] {
				ret = append(ret, nums2[j])
				flag = true
				break
			}
		}

		if !flag {
			ret = append(ret, -1)
		}

	}

	return ret
}

func main() {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"test1",
			args{
				nums1: []int{4, 1, 2},
				nums2: []int{1, 3, 4, 2},
			},
			[]int{-1, 3, -1},
		},
		{
			"test2",
			args{
				nums1: []int{2, 4},
				nums2: []int{1, 2, 3, 4},
			},
			[]int{3, -1},
		},
	}
	for _, tt := range tests {
		if got := nextGreaterElement(tt.args.nums1, tt.args.nums2); !reflect.DeepEqual(got, tt.want) {
			fmt.Printf("%s nextGreaterElement() = %v, want %v\n", tt.name, got, tt.want)
		}
	}
}

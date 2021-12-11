package main

import (
	"container/list"
	"fmt"
)

func validateStackSequences(pushed []int, popped []int) bool {
	if len(pushed) == 0 {
		return true
	}

	s := list.New()
	s.PushBack(pushed[0])

	next, target := 1, 0

	for target < len(popped) {

		// 判断栈顶的元素和要出栈的元素是否相同
		if s.Len() > 0 {
			elem := s.Back()
			if elem.Value.(int) == popped[target] {
				target++
				s.Remove(elem)
				continue
			}
		}

		// 栈为空 或者 栈顶的元素和要出栈的元素不相同
		// 需要入栈 如果没有元素可入 直接返回失败
		if next < len(pushed) {
			s.PushBack(pushed[next])
			next++
		} else {
			return false
		}

	}

	return true
}

func main() {
	type args struct {
		pushed []int
		popped []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{

		{
			"test1",
			args{[]int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1}},
			true,
		},
		{
			"test2",
			args{[]int{1, 2, 3, 4, 5}, []int{4, 3, 5, 1, 2}},
			false,
		},
	}
	for _, tt := range tests {
		if got := validateStackSequences(tt.args.pushed, tt.args.popped); tt.want != got {
			fmt.Printf("%s validateStackSequences() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

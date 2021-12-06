package main

import (
	"fmt"
)

func verifyPostorder(postorder []int) bool {
	if len(postorder) == 0 || len(postorder) == 1 {
		return true
	}

	end := len(postorder) - 1

	cur := 0

	for postorder[cur] < postorder[end] {
		cur++
	}

	left := cur

	for postorder[cur] > postorder[end] {
		cur++
	}

	return cur == end && verifyPostorder(postorder[:left]) && verifyPostorder(postorder[left:end])
}

func main() {
	type args struct {
		postorder []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test1",
			args{[]int{1, 6, 3, 2, 5}},
			false,
		},
		{
			"test2",
			args{[]int{1, 3, 2, 6, 5}},
			true,
		},
	}
	for _, tt := range tests {
		if got := verifyPostorder(tt.args.postorder); got != tt.want {
			fmt.Printf("verifyPostorder() = %v, want %v\n", got, tt.want)
		}
	}

}

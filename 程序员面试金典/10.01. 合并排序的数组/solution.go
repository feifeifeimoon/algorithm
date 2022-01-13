package main

import (
	"fmt"
	"reflect"
)

func merge(A []int, m int, B []int, n int) {
	curA, curB, cur := m-1, n-1, m+n-1

	for curA >= 0 && curB >= 0 {
		if A[curA] > B[curB] {
			A[cur] = A[curA]
			curA--
		} else {
			A[cur] = B[curB]
			curB--
		}
		cur--
	}

	for curA >= 0 {
		A[cur] = A[curA]
		curA--
		cur--
	}

	for curB >= 0 {
		A[cur] = B[curB]
		curB--
		cur--
	}

}

func main() {
	type args struct {
		A []int
		m int
		B []int
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"test1",
			args{[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3},
			[]int{1, 2, 2, 3, 5, 6},
		},
	}
	for _, tt := range tests {
		if merge(tt.args.A, tt.args.m, tt.args.B, tt.args.n); !reflect.DeepEqual(tt.args.A, tt.want) {
			fmt.Printf("%s merge() = %v, want %v\n", tt.name, tt.args.A, tt.want)
		}
	}
}

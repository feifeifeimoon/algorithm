package main

import (
	"fmt"
	"reflect"
)

func waysToStep(n int) int {

	switch n {
	case 1:
		return 1
	case 2:
		return 2
	case 3:
		return 4
	}
	cur, pre, ppre := 4, 2, 1

	for i := 4; i <= n; i++ {
		temp := (cur + pre + ppre) % 1000000007
		ppre = pre
		pre = cur
		cur = temp
	}
	return cur
}

func main() {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test1",
			args{n: 53},
			971328669,
		},
		{
			"test2",
			args{n: 3},
			4,
		},
		{
			"test3",
			args{n: 5},
			13,
		},
	}
	for _, tt := range tests {
		if got := waysToStep(tt.args.n); !reflect.DeepEqual(got, tt.want) {
			fmt.Printf("%s waysToStep() = %v, want %v\n", tt.name, got, tt.want)
		}
	}
}

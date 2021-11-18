package main

import (
	"fmt"
)

func fib(n int) int {
	const mod int = 1e9 + 7
	if n == 0 || n == 1 {
		return n
	}

	first, second, ret := 0, 0, 1
	for i := 2; i <= n; i++ {
		first, second = second, ret
		ret = (first + second) % mod
	}

	return ret

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
			args{2},
			1,
		},
		{
			"test2",
			args{5},
			5,
		},
		{
			"test3",
			args{95},
			407059028,
		},
	}
	for _, tt := range tests {
		if got := fib(tt.args.n); got != tt.want {
			fmt.Printf("fib() = %v, want %v\n", got, tt.want)
		}
	}
}

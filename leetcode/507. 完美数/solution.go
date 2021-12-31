package main

import (
	"fmt"
	"math"
)

func checkPerfectNumber(num int) bool {
	if num <= 2 {
		return false
	}
	count := 1
	for cur := 2; cur <= int(math.Sqrt(float64(num))); cur++ {
		if num%cur == 0 {
			count += cur
			count += num / cur
		}
	}

	return count == num
}

func main() {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test1",
			args{28},
			true,
		},
		{
			"test2",
			args{6},
			true,
		},
		{
			"test31",
			args{496},
			true,
		},
		{
			"test4",
			args{8128},
			true,
		}, {
			"test5",
			args{2},
			false,
		},
	}
	for _, tt := range tests {
		if got := checkPerfectNumber(tt.args.num); got != tt.want {
			fmt.Printf("%s checkPerfectNumber() = %v, want %v\n", tt.name, got, tt.want)
		}
	}
}

package main

import (
	"fmt"
	"reflect"
)

func swapNumbers(numbers []int) []int {
	numbers[0] ^= numbers[1]
	numbers[1] ^= numbers[0]
	numbers[0] ^= numbers[1]
	return numbers
}

func main() {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"test1",
			args{[]int{1, 2}},
			[]int{2, 1},
		},
		{
			"test2",
			args{[]int{1111, 2222}},
			[]int{2222, 1111},
		},
		{
			"test3",
			args{[]int{12345, 24567}},
			[]int{24567, 12345},
		},
	}
	for _, tt := range tests {
		if got := swapNumbers(tt.args.numbers); !reflect.DeepEqual(got, tt.want) {
			fmt.Printf("%s swapNumbers() = %v, want %v\n", tt.name, got, tt.want)
		}
	}
}

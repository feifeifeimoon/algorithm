package main

import (
	"fmt"
	"reflect"
)

func plusOne(digits []int) []int {

	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] += 1

		if digits[i] == 10 {
			digits[i] = 0

			// the first element
			if i == 0 {
				return append([]int{1}, digits...)
			}
			continue
		}
		break
	}

	return digits
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
			"add",
			args{digits: []int{1, 2, 3}},
			[]int{1, 2, 4},
		},
		{
			"overflow",
			args{digits: []int{1, 2, 9}},
			[]int{1, 3, 0},
		},
		{
			"overflow2",
			args{digits: []int{9, 9, 9}},
			[]int{1, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		if got := plusOne(tt.args.digits); !reflect.DeepEqual(got, tt.want) {
			fmt.Printf("plusOne() = %v, want %v\n", got, tt.want)
		}
	}
}

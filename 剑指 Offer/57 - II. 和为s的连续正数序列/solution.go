package main

import (
	"fmt"
	"reflect"
)

func findContinuousSequence(target int) [][]int {

	begin, end, sum := 1, 1, 0
	var ret [][]int

	for begin <= target/2 {

		if sum < target {
			sum += end
			end++
		} else if sum > target {
			sum -= begin
			begin++
		} else {
			temp := make([]int, end-begin)
			for i := begin; i < end; i++ {
				temp[i-begin] = i
			}
			ret = append(ret, temp)

			sum -= begin
			begin++

		}

	}
	return ret
}

func main() {
	type args struct {
		target int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{

		{
			"test1",
			args{9},
			[][]int{{2, 3, 4}, {4, 5}},
		},
		{
			"test2",
			args{15},
			[][]int{{1, 2, 3, 4, 5}, {4, 5, 6}, {7, 8}},
		},
	}
	for _, tt := range tests {
		if got := findContinuousSequence(tt.args.target); !reflect.DeepEqual(tt.want, got) {
			fmt.Printf("%s findContinuousSequence() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

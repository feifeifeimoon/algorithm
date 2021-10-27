package main

import (
	"fmt"
	"math"
	"reflect"
)

/*
	题目说明 L*W=area 证明area肯定可以被W整除
	又给定条件 W <= L, 所以W最大值就是等于L的时候 此时相当于 W*W=area 所以W的最大值就是area开平方
	找到最大之后 只要从最大值开始依次递减 第一个能被area整除的就是符合条件的
*/
func constructRectangle(area int) []int {
	w := int(math.Sqrt(float64(area)))

	for area%w > 0 {
		w--
	}

	return []int{area / w, w}
}

func main() {
	type args struct {
		area int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"test1",
			args{area: 37},
			[]int{37, 1},
		},
	}
	for _, tt := range tests {
		if got := constructRectangle(tt.args.area); !reflect.DeepEqual(got, tt.want) {
			fmt.Printf("constructRectangle() = %v, want %v\n", got, tt.want)
		}
	}
}

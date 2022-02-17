package main

import (
	"fmt"
	"reflect"
)

func multiply(A [][]int, B [][]int) [][]int {
	mA, nA, _, nB := len(A), len(A[0]), len(B), len(B[0])

	ret := make([][]int, mA)
	for i := 0; i < mA; i++ {
		ret[i] = make([]int, nB)
	}

	for i := 0; i < mA; i++ {
		for j := 0; j < nB; j++ {

			count := 0
			for k := 0; k < nA; k++ {
				count += A[i][k] * B[k][j]
			}
			ret[i][j] = count
		}
	}

	return ret
}

func main() {
	type args struct {
		A, B [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"test1",
			args{A: [][]int{{1, 0, 0}, {-1, 0, 3}}, B: [][]int{{7, 0, 0}, {0, 0, 0}, {0, 0, 1}}},
			[][]int{{7, 0, 0}, {-7, 0, 3}},
		},
	}
	for _, tt := range tests {
		if got := multiply(tt.args.A, tt.args.B); !reflect.DeepEqual(got, tt.want) {
			fmt.Printf("%s multiply() = %v, want %v\n", tt.name, got, tt.want)
		}
	}
}

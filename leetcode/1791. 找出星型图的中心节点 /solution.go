package main

import (
	"fmt"
	"reflect"
)

func findCenter(edges [][]int) int {
	flag := make([]int, len(edges)+2)
	for _, edge := range edges {
		flag[edge[0]]++
		flag[edge[1]]++
	}

	for i := range flag {
		if flag[i] == len(edges) {
			return i
		}
	}
	return 0
}

func main() {
	type args struct {
		edges [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test1",
			args{edges: [][]int{{1, 2}, {2, 3}, {4, 2}}},
			2,
		},
	}
	for _, tt := range tests {
		if got := findCenter(tt.args.edges); !reflect.DeepEqual(got, tt.want) {
			fmt.Printf("%s findCenter() = %v, want %v\n", tt.name, got, tt.want)
		}
	}
}

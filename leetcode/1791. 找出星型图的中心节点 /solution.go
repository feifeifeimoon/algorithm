package main

import (
	"fmt"
	"reflect"
)

func findCenter(edges [][]int) int {
	if edges[0][0] == edges[1][0] || edges[0][0] == edges[1][1] {
		return edges[0][0]
	}
	return edges[0][1]
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

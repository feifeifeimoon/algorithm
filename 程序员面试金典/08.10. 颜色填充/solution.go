package main

import (
	"fmt"
	"reflect"
)

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	dx := []int{1, -1, 0, 0}
	dy := []int{0, 0, 1, -1}
	row, col, color := len(image), len(image[0]), image[sr][sc]

	var dfs func(x, y int)
	dfs = func(x, y int) {
		if x < 0 || x >= row || y < 0 || y >= col || image[x][y] != color || image[x][y] < 0 {
			return
		}
		image[x][y] = -newColor
		for i := range dx {
			dfs(x+dx[i], y+dy[i])
		}
	}

	dfs(sr, sc)
	for i := range image {
		for j := range image[i] {
			if image[i][j] < 0 {
				image[i][j] = newColor
			}
		}
	}

	return image
}

func main() {
	type args struct {
		image            [][]int
		sr, sc, newColor int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"test1",
			args{[][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}, 1, 1, 2},
			[][]int{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}},
		},
	}
	for _, tt := range tests {
		if got := floodFill(tt.args.image, tt.args.sr, tt.args.sc, tt.args.newColor); !reflect.DeepEqual(got, tt.want) {
			fmt.Printf("%s floodFill() = %v, want %v\n", tt.name, got, tt.want)
		}
	}
}

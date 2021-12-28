package main

import (
	"fmt"
	"reflect"
)

var dx = []int{1, -1, 0, 0}
var dy = []int{0, 0, 1, -1}

func solve(board [][]byte) {
	row, col := len(board), len(board[0])

	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i >= row || j < 0 || j >= col || board[i][j] != 'O' {
			return
		}
		board[i][j] = 'A'
		for k := range dx {
			dfs(i+dx[k], j+dy[k])
		}
	}

	for i := 0; i < col; i++ {
		dfs(0, i)
	}

	for i := 0; i < row; i++ {
		dfs(i, col-1)
	}

	for i := 0; i < col; i++ {
		dfs(row-1, i)
	}

	for i := 0; i < row; i++ {
		dfs(i, 0)
	}

	for i := range board {
		for j := range board[i] {
			if board[i][j] == 'A' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}

}

func main() {
	type args struct {
		board [][]byte
	}
	tests := []struct {
		name string
		args args
		want [][]byte
	}{
		{
			"test1",
			args{[][]byte{{'X', 'O', 'X'}, {'O', 'X', 'O'}, {'X', 'O', 'X'}}},
			[][]byte{{'X', 'O', 'X'}, {'O', 'X', 'O'}, {'X', 'O', 'X'}},
		},
	}
	for _, tt := range tests {
		if solve(tt.args.board); !reflect.DeepEqual(tt.args.board, tt.want) {
			fmt.Printf("%s solve() = %v, want %v\n", tt.name, tt.args.board, tt.want)
		}
	}
}

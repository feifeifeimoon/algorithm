package main

import "fmt"

func countBattleships(board [][]byte) int {
	ret := 0
	do := func(i, j int) {
		if i+1 < len(board) && board[i+1][j] == 'X' {
			for ; i+1 < len(board) && board[i+1][j] == 'X'; i++ {
				board[i+1][j] = '.'
			}
			return
		}

		for ; j+1 < len(board[0]) && board[i][j+1] == 'X'; j++ {
			board[i][j+1] = '.'
		}
	}

	for i := range board {
		for j := range board[i] {
			if board[i][j] == 'X' {
				ret++
				do(i, j)
			}
		}
	}

	return ret
}

func main() {
	type args struct {
		board [][]byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test1",
			args{[][]byte{{'X', '.', '.', 'X'}, {'.', '.', '.', 'X'}, {'.', '.', '.', 'X'}}},
			2,
		},
	}
	for _, tt := range tests {
		if got := countBattleships(tt.args.board); got != tt.want {
			fmt.Printf("%s countBattleships() = %v, want %v\n", tt.name, got, tt.want)
		}
	}
}

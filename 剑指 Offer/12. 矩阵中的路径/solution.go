package main

import "fmt"

func exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}

	wordLen := len(word)
	var dfs func(i, j, index int) bool

	dfs = func(i, j, index int) bool {
		if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
			return false
		}

		//fmt.Println(board[i][j], word[index])
		if board[i][j] != word[index] {
			return false
		}

		temp := board[i][j]
		board[i][j] = '0'
		//defer func() {
		//	board[i][j] = temp
		//}()
		index++
		if index == wordLen {
			return true
		}

		ret := dfs(i+1, j, index) || dfs(i-1, j, index) || dfs(i, j+1, index) || dfs(i, j-1, index)
		board[i][j] = temp
		return ret
		//return dfs(i+1, j, index) || dfs(i-1, j, index) || dfs(i, j+1, index) || dfs(i, j-1, index)
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}

	return false
}

func main() {
	type args struct {
		board [][]byte
		word  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		//{
		//	"test1",
		//	args{board: [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, word: "ABCCED"},
		//	true,
		//},
		//{
		//	"test2",
		//	args{board: [][]byte{{'a', 'b'}, {'c', 'd'}}, word: "abcd"},
		//	false,
		//},
		//{
		//	"test3",
		//	args{board: [][]byte{{'a', 'b'}}, word: "ba"},
		//	true,
		//},
		{
			"test4",
			args{board: [][]byte{{'C', 'A', 'A'}, {'A', 'A', 'A'}, {'B', 'C', 'D'}}, word: "AAB"},
			true,
		},
	}
	for _, tt := range tests {
		if got := exist(tt.args.board, tt.args.word); tt.want != got {
			fmt.Printf("exist() = %v, want %v", got, tt.want)
		}
	}
}

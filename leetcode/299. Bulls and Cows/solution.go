package main

import (
	"fmt"
)

func getHint(secret string, guess string) string {

	countA, countB := 0, 0

	var cntS, cntG [10]int

	for i := 0; i < len(secret); i++ {
		if secret[i] == guess[i] {
			countA++
			continue
		}

		cntS[secret[i]-'0']++
		cntG[guess[i]-'0']++

	}

	for i := 0; i < 10; i++ {
		if cntS[i] < cntG[i] {
			countB += cntS[i]
		} else {
			countB += cntG[i]
		}
	}

	return fmt.Sprintf("%dA%dB", countA, countB)
}

func main() {
	type args struct {
		secret string
		guess  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test1",
			args{"1807", "7810"},
			"1A3B",
		},
		{
			"test2",
			args{"1123", "0111"},
			"1A1B",
		},
		{
			"test3",
			args{"1", "0"},
			"0A0B",
		},
		{
			"test4",
			args{"1", "1"},
			"1A0B",
		},
	}
	for _, tt := range tests {
		if got := getHint(tt.args.secret, tt.args.guess); got != tt.want {
			fmt.Printf("getHint() = %v, want %v\n", got, tt.want)
		}
	}
}

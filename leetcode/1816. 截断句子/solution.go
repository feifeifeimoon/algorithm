package main

import "fmt"

func truncateSentence(s string, k int) string {
	index, count := 0, 0
	for index = 0; index < len(s); index++ {
		if s[index] != ' ' {
			continue
		} else {
			count++
			if count == k {
				break
			}
		}

	}

	return s[:index]
}

func main() {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test1",
			args{"Hello how are you Contestant", 4},
			"Hello how are you",
		},
		{
			"test2",
			args{"What is the solution to this problem", 4},
			"What is the solution",
		},
		{
			"test3",
			args{"chopper is not a tanukin", 5},
			"chopper is not a tanukin",
		},
	}
	for _, tt := range tests {
		if got := truncateSentence(tt.args.s, tt.args.k); got != tt.want {
			fmt.Printf("maxPower() = %v, want %v\n", got, tt.want)
		}
	}
}

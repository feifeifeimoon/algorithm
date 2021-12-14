package main

import "fmt"

func canPermutePalindrome(s string) bool {
	count := make(map[byte]int)

	for i := range s {
		if _, ok := count[s[i]]; !ok {
			count[s[i]] = 1
		} else {
			count[s[i]]++
		}
	}

	n := 0
	for _, v := range count {
		if v%2 != 0 {
			n++
			if n > 1 {
				return false
			}
		}
	}

	return true
}

func main() {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test1",
			args{"tactcoa"},
			true,
		},
	}
	for _, tt := range tests {
		if got := canPermutePalindrome(tt.args.s); got != tt.want {
			fmt.Printf("%s canPermutePalindrome() = %v, want %v\n", tt.name, got, tt.want)
		}
	}
}

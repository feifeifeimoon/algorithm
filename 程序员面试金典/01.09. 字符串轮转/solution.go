package main

import (
	"fmt"
	"strings"
)

func isFlipedString(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	return strings.Contains(s2+s2, s1)
}

func main() {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test1",
			args{"waterbottle", "erbottlewat"},
			true,
		},
		{
			"test2",
			args{"aa", "abc"},
			false,
		},
		{
			"test2",
			args{"aa", "abc"},
			false,
		},
	}
	for _, tt := range tests {
		if got := isFlipedString(tt.args.s1, tt.args.s2); got != tt.want {
			fmt.Printf("%s isFlipedString() = %v, want %v\n", tt.name, got, tt.want)
		}
	}
}

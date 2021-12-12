package main

import (
	"fmt"
)

func isUnique(astr string) bool {
	if len(astr) > 26 {
		return false
	}

	ret := 0

	for i := range astr {
		dis := astr[i] - 'a'
		temp := 1 << dis

		if ret&temp == 0 {
			ret |= temp
		} else {
			return false
		}

	}

	return true
}

func main() {
	type args struct {
		astr string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test1",
			args{astr: "abc"},
			true,
		},
	}
	for _, tt := range tests {
		if got := isUnique(tt.args.astr); got != tt.want {
			fmt.Printf("%s isUnique() = %v, want %v\n", tt.name, got, tt.want)
		}
	}
}

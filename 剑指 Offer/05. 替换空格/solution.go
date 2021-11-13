package main

import (
	"fmt"
)

func replaceSpace(s string) string {
	var ret []byte

	for _, c := range s {
		if c == ' ' {
			ret = append(ret, []byte("%20")...)
		} else {
			ret = append(ret, byte(c))
		}
	}

	return string(ret)
}

func main() {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test1",
			args{"a b c"},
			"a%20b%20c",
		},
	}
	for _, tt := range tests {
		if got := replaceSpace(tt.args.s); got != tt.want {
			fmt.Printf("replaceSpace() = %v, want %v\n", got, tt.want)
		}
	}
}

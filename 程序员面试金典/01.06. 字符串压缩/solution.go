package main

import (
	"fmt"
	"strconv"
	"strings"
)

func compressString(S string) string {
	if len(S) == 0 {
		return ""
	}

	builder := strings.Builder{}
	builder.WriteByte(S[0])

	cur := S[0]
	count := 0

	for i := range S {
		if S[i] == cur {
			count++
		} else {
			builder.WriteString(strconv.Itoa(count))
			count = 1
			builder.WriteByte(S[i])
			cur = S[i]
		}
	}
	builder.WriteString(strconv.Itoa(count))

	if builder.Len() >= len(S) {
		return S
	}

	return builder.String()
}

func main() {
	type args struct {
		S string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test1",
			args{"aabcccccaaa"},
			"a2b1c5a3",
		},
		{
			"test2",
			args{"abbccd"},
			"abbccd",
		},
	}
	for _, tt := range tests {
		if got := compressString(tt.args.S); got != tt.want {
			fmt.Printf("%s compressString() = %v, want %v\n", tt.name, got, tt.want)
		}
	}
}

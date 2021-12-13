package main

import (
	"fmt"
	"strings"
)

func replaceSpaces(S string, length int) string {
	builder := strings.Builder{}
	for i := 0; i < length; i++ {
		if S[i] != ' ' {
			builder.WriteByte(S[i])
		} else {
			builder.WriteString("%20")
		}
	}
	return builder.String()
}

func main() {
	type args struct {
		S      string
		length int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test1",
			args{"Mr John Smith    ", 13},
			"Mr%20John%20Smith",
		},
	}
	for _, tt := range tests {
		if got := replaceSpaces(tt.args.S, tt.args.length); got != tt.want {
			fmt.Printf("%s replaceSpaces() = %v, want %v\n", tt.name, got, tt.want)
		}
	}
}

package main

import (
	"fmt"
	"reflect"
	"strings"
)

func findOcurrences(text string, first string, second string) []string {
	var ret []string
	s := strings.Split(text, " ")
	for i := 0; i < len(s)-2; {
		if s[i] == first && s[i+1] == second {
			ret = append(ret, s[i+2])
		}

		i++
	}
	return ret
}

func main() {
	type args struct {
		text, first, second string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"test1",
			args{"alice is a good girl she is a good student", "a", "good"},
			[]string{"girl", "student"},
		},
		{
			"test2",
			args{"we will we will rock you", "we", "will"},
			[]string{"we", "rock"},
		},
	}
	for _, tt := range tests {
		if got := findOcurrences(tt.args.text, tt.args.first, tt.args.second); !reflect.DeepEqual(got, tt.want) {
			fmt.Printf("%s findOcurrences() = %v, want %v\n", tt.name, got, tt.want)
		}
	}
}

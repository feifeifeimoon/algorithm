package main

import (
	"fmt"
	"reflect"
)

var num2str = [][]byte{
	{},
	{},
	{'a', 'b', 'c'},
	{'d', 'e', 'f'},
	{'g', 'h', 'i'},
	{'j', 'k', 'l'},
	{'m', 'n', 'o'},
	{'p', 'q', 'r', 's'},
	{'t', 'u', 'v'},
	{'w', 'x', 'y', 'z'},
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	var ret []string

	var dfs func(index int, str []byte)
	dfs = func(index int, str []byte) {
		if index >= len(digits) {
			ret = append(ret, string(str))
			return
		}
		arr := num2str[digits[index]-'0']
		for i := range arr {
			str[index] = arr[i]
			dfs(index+1, str)
		}
	}
	dfs(0, make([]byte, len(digits)))

	return ret
}

func main() {
	type args struct {
		digits string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"test1",
			args{"23"},
			[]string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
	}
	for _, tt := range tests {
		if got := letterCombinations(tt.args.digits); !reflect.DeepEqual(got, tt.want) {
			fmt.Printf("%s letterCombinations() = %v, want %v\n", tt.name, got, tt.want)
		}
	}
}

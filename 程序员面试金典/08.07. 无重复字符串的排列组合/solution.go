package main

import (
	"fmt"
	"reflect"
)

func permutation(S string) []string {
	var ret []string

	flag := make([]int, len(S))
	// str := make([]byte, len(S))

	var dfs func(cur int, str []byte)

	dfs = func(cur int, str []byte) {
		if cur == len(S) {
			ret = append(ret, string(str))
		}

		for i := range S {
			if flag[i] == 0 {
				flag[i] = 1
				str[cur] = S[i]
				dfs(cur+1, str)
				flag[i] = 0
			}
		}
	}

	dfs(0, make([]byte, len(S)))
	return ret
}

func main() {
	type args struct {
		S string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"test1",
			args{"qwe"},
			[]string{"qwe", "qew", "wqe", "weq", "eqw", "ewq"},
		},
	}
	for _, tt := range tests {
		if got := permutation(tt.args.S); !reflect.DeepEqual(got, tt.want) {
			fmt.Printf("%s permutation() = %v, want %v\n", tt.name, got, tt.want)
		}
	}
}

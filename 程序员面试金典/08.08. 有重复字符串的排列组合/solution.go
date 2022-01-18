package main

import (
	"fmt"
	"reflect"
)

func permutation(S string) []string {
	var ret []string
	flag := make([]int, len(S))

	var dfs func(index int, str []byte)
	dfs = func(index int, str []byte) {
		if index >= len(S) {
			ret = append(ret, string(str))
			return
		}

		set := make(map[byte]struct{})

		for i := range S {
			if flag[i] == 0 {
				if _, ok := set[S[i]]; ok {
					continue
				}
				set[S[i]] = struct{}{}
				flag[i] = 1
				str[index] = S[i]
				dfs(index+1, str)
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
			args{"qqe"},
			[]string{"qqe", "qeq", "eqq"},
		},
	}
	for _, tt := range tests {
		if got := permutation(tt.args.S); !reflect.DeepEqual(got, tt.want) {
			fmt.Printf("%s permutation() = %v, want %v\n", tt.name, got, tt.want)
		}
	}
}

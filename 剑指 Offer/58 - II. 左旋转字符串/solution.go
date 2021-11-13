package main

import "fmt"

/*
通过局部反转+整体反转 达到左旋转的目的。

具体步骤为：
	1.反转区间为前n的子串
	2.反转区间为n到末尾的子串
	3.反转整个字符串
*/

func reverseString(s string) string {
	var ret []byte
	for i := len(s) - 1; i >= 0; i-- {
		ret = append(ret, s[i])
	}
	return string(ret)
}

func reverseLeftWords(s string, n int) string {
	return reverseString(reverseString(s[:n]) + reverseString(s[n:]))
	// or return s[n:] + s[:n]
}

func main() {
	type args struct {
		s string
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test1",
			args{"abcdefg", 2},
			"cdefgab",
		},
		{
			"test21",
			args{"lrloseumgh", 6},
			"umghlrlose",
		},
	}
	for _, tt := range tests {
		if got := reverseLeftWords(tt.args.s, tt.args.n); got != tt.want {
			fmt.Printf("reverseLeftWords() = %v, want %v\n", got, tt.want)
		}
	}
}

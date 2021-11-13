package main

import (
	"fmt"
)

func detectCapitalUse(word string) bool {
	// 判断第一个字符
	if word[0] >= 65 && word[0] <= 90 {

		// 前两个都是大写
		if len(word) > 1 && word[1] >= 65 && word[1] <= 90 {
			// 有一个不是大写就错误
			for i := 1; i < len(word); i++ {
				if word[i] < 65 || word[i] > 90 {
					return false
				}
			}
		} else {
			// 第一个大写 其它都是小写
			for i := 1; i < len(word); i++ {
				if word[i] >= 65 && word[i] <= 90 {
					return false
				}
			}
		}

	} else {
		// 有一个不是小写就错误
		for _, s := range word {
			if s >= 65 && s <= 90 {
				return false
			}

		}
	}

	return true
}

func main() {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test1",
			args{"USA"},
			true,
		},
		{
			"test2",
			args{"FlagG"},
			false,
		},
		{
			"test3",
			args{"hello"},
			true,
		},
		{
			"test4",
			args{"Hello"},
			true,
		},
		{
			"test4",
			args{"a"},
			true,
		},
		{
			"test4",
			args{"H"},
			true,
		},
	}
	for _, tt := range tests {
		if got := detectCapitalUse(tt.args.word); got != tt.want {
			fmt.Printf("detectCapitalUse() = %v, want %v\n", got, tt.want)
		}
	}
}

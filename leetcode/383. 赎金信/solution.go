package main

import (
	"fmt"
)

func canConstruct(ransomNote string, magazine string) bool {
	// 26个字母
	count := [26]int{}
	for _, v := range magazine {
		count[v-'a']++
	}

	for _, v := range ransomNote {
		count[v-'a']--
		if count[v-'a'] < 0 {
			return false
		}
	}

	return true
}

func main() {
	type args struct {
		ransomNote string
		magazine   string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test1",
			args{
				ransomNote: "a",
				magazine:   "b",
			},
			false,
		},
		{
			"test2",
			args{
				ransomNote: "aa",
				magazine:   "ab",
			},
			false,
		},
		{
			"test3",
			args{
				ransomNote: "aa",
				magazine:   "aab",
			},
			true,
		},
	}
	for _, tt := range tests {
		if got := canConstruct(tt.args.ransomNote, tt.args.magazine); !got == tt.want {
			fmt.Printf("[%s] canConstruct() = %v, want %v\n", tt.name, got, tt.want)
		}
	}
}

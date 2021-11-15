package main

import (
	"fmt"
)

func firstUniqChar(s string) byte {
	m := make(map[byte]int)

	for _, b := range s {
		if _, ok := m[byte(b)]; !ok {
			m[byte(b)] = 1
		} else {
			m[byte(b)]++
		}
	}

	for _, b := range s {

		if m[byte(b)] == 1 {
			return byte(b)
		}
	}

	return ' '
}

func main() {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{
			"test1",
			args{"abaccdeff"},
			'b',
		},
		{
			"test2",
			args{""},
			' ',
		},
	}
	for _, tt := range tests {
		if got := firstUniqChar(tt.args.s); got != tt.want {
			fmt.Printf("firstUniqChar() = %v, want %v\n", got, tt.want)
		}
	}
}

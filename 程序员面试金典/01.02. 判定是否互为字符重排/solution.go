package main

import "fmt"

func CheckPermutation(s1 string, s2 string) bool {
	var count [26]int

	for i := range s1 {
		count[s1[i]-'a']++
	}

	for i := range s2 {
		count[s2[i]-'a']--
	}

	for i := range count {
		if count[i] != 0 {
			return false
		}
	}

	return true
}

func main() {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test1",
			args{"abc", "cba"},
			true,
		},
		{
			"test2",
			args{"abc", "bda"},
			false,
		},
	}
	for _, tt := range tests {
		if got := CheckPermutation(tt.args.s1, tt.args.s2); got != tt.want {
			fmt.Printf("%s CheckPermutation() = %v, want %v\n", tt.name, got, tt.want)
		}
	}
}

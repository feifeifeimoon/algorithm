package main

import (
	"fmt"
	"math"
)

func help(first, second string) bool {
	diff, i, j := 0, 0, 0

	for j < len(second) {
		if first[i] == second[j] {
			i++
			j++
		} else {
			i++
			diff++
			if diff > 1 {
				return false
			}
		}
	}

	return true
}

func oneEditAway(first string, second string) bool {
	firstLen, secondLen := len(first), len(second)
	if math.Abs(float64(firstLen-secondLen)) > 1 {
		return false
	}

	if firstLen == secondLen {
		diff := 0
		for i := 0; i < firstLen; i++ {
			if first[i] != second[i] {
				diff++
				if diff > 1 {
					return false
				}
			}
		}
	} else if firstLen > secondLen {
		return help(first, second)
	} else {
		return help(second, first)
	}
	return true
}

func main() {
	type args struct {
		first  string
		second string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test1",
			args{"pale", "ple"},
			true,
		},
		{
			"test2",
			args{"pales", "pal"},
			false,
		},
		{
			"test3",
			args{"abcdefg", "abccefg"},
			true,
		},
		{
			"test4",
			args{"abc", "abcd"},
			true,
		}, {
			"test4",
			args{"teacher", "taches"},
			false,
		},
	}
	for _, tt := range tests {
		if got := oneEditAway(tt.args.first, tt.args.second); got != tt.want {
			fmt.Printf("%s oneEditAway() = %v, want %v\n", tt.name, got, tt.want)
		}
	}
}

package main

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"
)

type iter struct {
	score int
	index int
}

func findRelativeRanks(score []int) []string {
	var desc = [3]string{"Gold Medal", "Silver Medal", "Bronze Medal"}

	r := make([]iter, len(score))
	for i, s := range score {
		r[i].score, r[i].index = s, i
	}

	sort.Slice(r, func(i, j int) bool {
		return r[i].score > r[j].score
	})

	ret := make([]string, len(score))
	for i, v := range r {
		if i < 3 {
			ret[v.index] = desc[i]
		} else {
			ret[v.index] = strconv.Itoa(i + 1)
		}
	}

	return ret

}

func main() {
	type args struct {
		score []int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"test1",
			args{[]int{5, 4, 3, 2, 1}},
			[]string{"Gold Medal", "Silver Medal", "Bronze Medal", "4", "5"},
		},
		{
			"test1",
			args{[]int{10, 3, 8, 9, 4}},
			[]string{"Gold Medal", "5", "Bronze Medal", "Silver Medal", "4"},
		},
	}
	for _, tt := range tests {
		if got := findRelativeRanks(tt.args.score); !reflect.DeepEqual(got, tt.want) {
			fmt.Printf("findRelativeRanks() = %v, want %v\n", got, tt.want)
		}
	}
}

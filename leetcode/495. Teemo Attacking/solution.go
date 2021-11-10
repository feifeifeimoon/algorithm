package main

import (
	"fmt"
)

func findPoisonedDuration(timeSeries []int, duration int) int {
	all := len(timeSeries) * duration

	for i := 0; i < len(timeSeries)-1; i++ {
		temp := timeSeries[i+1] - timeSeries[i]
		if temp < duration {
			all -= duration - temp
		}
	}

	return all
}

func main() {
	type args struct {
		timeSeries []int
		duration   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test1",
			args{
				timeSeries: []int{1, 4},
				duration:   2,
			},
			4,
		},
		{
			"test2",
			args{
				timeSeries: []int{1, 2},
				duration:   2,
			},
			3,
		},
		{
			"test2",
			args{
				timeSeries: []int{1, 2, 3, 4, 5},
				duration:   5,
			},
			9,
		},
	}
	for _, tt := range tests {
		if got := findPoisonedDuration(tt.args.timeSeries, tt.args.duration); got != tt.want {
			fmt.Printf("findPoisonedDuration() = %v, want %v\n", got, tt.want)
		}
	}
}

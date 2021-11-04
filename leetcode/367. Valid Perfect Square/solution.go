package main

import "fmt"

/*
	二分查找
*/
func isPerfectSquare(num int) bool {

	low, high := 1, num

	for low <= high {
		mid := low + (high-low)/2
		temp := mid * mid

		if temp == num {
			return true
		} else if temp < num {
			low = mid + 1
		} else {
			high = mid - 1
		}

	}

	return false
}

func main() {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test1",
			args{4},
			true,
		},
		{
			"test2",
			args{16},
			true,
		},
		{
			"test3",
			args{14},
			false,
		},
		{
			"test4",
			args{1},
			true,
		},
	}
	for _, tt := range tests {
		if got := isPerfectSquare(tt.args.num); got != tt.want {
			fmt.Printf("isPerfectSquare() = %v, want %v\n", got, tt.want)
		}
	}

}

package main

import "fmt"

func hammingWeight(num uint32) int {

	ret := 0
	for num != 0 {

		if (num & 1) == 1 {
			ret++
		}
		num >>= 1
	}
	return ret
}

func main() {
	type args struct {
		num uint32
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test1",
			args{11},
			3,
		},
	}
	for _, tt := range tests {
		if got := hammingWeight(tt.args.num); tt.want != got {
			fmt.Printf("hammingWeight() = %v, want %v", got, tt.want)
		}
	}
}

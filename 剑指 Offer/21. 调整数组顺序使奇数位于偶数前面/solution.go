package main

import (
	"fmt"
)

func swap(nums []int, i, j int) {
	tmp := nums[i]
	nums[i] = nums[j]
	nums[j] = tmp
}

func exchange(nums []int) []int {
	start, end := 0, len(nums)-1

	for start < end {
		if nums[start]%2 == 0 {
			swap(nums, start, end)
			end--
			continue
		} else {
			start++
			continue
		}

	}

	return nums
}

func main() {

	args := []int{1, 2, 3, 4, 5, 7, 9, 6, 8, 11, 20}
	fmt.Println(exchange(args))

}

package main

func minArray(numbers []int) int {

	begin, end := 0, len(numbers)-1

	for begin < end {
		mid := begin + (end-begin)/2

		if numbers[mid] < numbers[end] {
			// 注意 上面判断去掉的相等 这里就不用 mid + 1
			end = mid
		} else if numbers[mid] > numbers[end] {
			begin = mid + 1
		} else {
			end--
		}

	}

	return numbers[begin]
}

func main() {

}

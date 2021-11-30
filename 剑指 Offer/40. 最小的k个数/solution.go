package main

import (
	"fmt"
	"reflect"
)

/*
 构建长度为k的大根堆 k-len(arr) 依次与大根堆的最大元素进行比较 如果小于 就交换位置并且对堆进行调整

 数组转换为堆
	节点i的左子树下标为 2 * i + 1
	节点i的右子树下标为 2 * i + 2
	节点i的副节点下标为 (i - 1) / 2
*/

/*
	heapInsert
	判断当前节点是否小于父节点 如果小 交换位置继续判断 否则推出的
*/
func heapInsert(arr []int, index int) {

	for arr[index] > arr[(index-1)/2] {
		arr[index], arr[(index-1)/2] = arr[(index-1)/2], arr[index]
		index = (index - 1) / 2
	}

}

/*
	heapModify
	对堆进行调整 从堆顶开始 如果当前节点小于左子树和右子树中较大的节点 就交换两个的位置 继续向下比较
*/
func heapModify(arr []int, index, size int) {
	left := index*2 + 1
	for left < size {
		lIndex := left
		// 如果右子树存在 并且右子树的值大于左子树
		if left+1 < size && arr[left+1] > arr[left] {
			lIndex = left + 1
		}

		if arr[index] > arr[lIndex] {
			break
		}

		arr[index], arr[lIndex] = arr[lIndex], arr[index]
		index = lIndex
		left = index*2 + 1

	}

}

func getLeastNumbers(arr []int, k int) []int {
	for i := 1; i < k; i++ {
		heapInsert(arr, i)
	}
	for i := k; i < len(arr); i++ {
		if arr[i] < arr[0] {
			arr[i], arr[0] = arr[0], arr[i]
			heapModify(arr, 0, k)
		}
	}
	return arr[:k]
}

func main() {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"test1",
			args{
				arr: []int{3, 2, 1},
				k:   2,
			},
			[]int{1, 2},
		},
		{
			"test2",
			args{
				arr: []int{3, 4, 5, 6, 7, 8, 1, 2},
				k:   3,
			},
			[]int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		if got := getLeastNumbers(tt.args.arr, tt.args.k); !reflect.DeepEqual(got, tt.want) {
			fmt.Printf("fib() = %v, want %v\n", got, tt.want)
		}
	}
}

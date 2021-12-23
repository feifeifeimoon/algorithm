package main

import "fmt"

type TripleInOne struct {
	stack [3][]int
	index [3]int
	size  int
}

func Constructor(stackSize int) TripleInOne {
	return TripleInOne{
		stack: [3][]int{make([]int, stackSize), make([]int, stackSize), make([]int, stackSize)},
		index: [3]int{-1, -1, -1},
		size:  stackSize,
	}
}

func (t *TripleInOne) Push(stackNum int, value int) {
	pos := t.index[stackNum] + 1
	if pos >= t.size {
		return
	}

	t.stack[stackNum][pos] = value
	t.index[stackNum]++
}

func (t *TripleInOne) Pop(stackNum int) int {
	pos := t.index[stackNum]
	if pos < 0 {
		return -1
	}

	t.index[stackNum]--
	return t.stack[stackNum][pos]
}

func (t *TripleInOne) Peek(stackNum int) int {
	pos := t.index[stackNum]
	if pos < 0 {
		return -1
	}

	return t.stack[stackNum][pos]
}

func (t *TripleInOne) IsEmpty(stackNum int) bool {
	return t.index[stackNum] == -1
}

//["TripleInOne", "peek", "pop", "isEmpty", "push", "pop", "push", "push", "pop", "push", "push", "isEmpty", "pop", "peek", "push", "peek", "isEmpty", "peek", "pop", "push", "isEmpty", "pop", "peek", "peek", "pop", "peek", "isEmpty", "pop", "push", "isEmpty", "peek", "push", "peek", "isEmpty", "isEmpty", "isEmpty", "isEmpty", "peek", "push", "push", "peek", "isEmpty", "peek", "isEmpty", "push", "push", "push", "peek", "peek", "pop", "push", "push", "isEmpty", "peek", "pop", "push", "peek", "peek", "pop", "isEmpty", "pop", "peek", "peek", "push", "push"]
//[[18],           [1],   [2],    [1],      [2, 40], [2],  [0, 44], [1, 40], [0], [1, 54], [0, 42], [0],      [1], [1], [0, 56], [2], [0], [2], [2],   [1, 15], [1],     [1], [0], [2], [0], [0],    [1], [0], [0, 32], [0], [0], [0, 30], [2], [1], [1],        [0], [0], [0], [0, 56], [1, 40], [2], [0], [0], [0],     [2, 11], [0, 16], [0, 3], [2], [0], [2], [0, 39], [0, 63], [1], [2],     [0], [2, 36], [1], [0], [2], [1], [0],    [1], [2], [0, 8], [0, 38]]
func main() {
	s := Constructor(18)
	fmt.Println(s.Peek(1))
	fmt.Println(s.Pop(2))
	fmt.Println(s.IsEmpty(1))
	s.Push(2, 40)
	fmt.Println(s.Pop(2))
	s.Push(0, 44)
	s.Push(1, 40)
	fmt.Println(s.Pop(0))
	s.Push(1, 54)
	s.Push(0, 42)
	fmt.Println(s.IsEmpty(0))
	fmt.Println(s.Pop(1))
	fmt.Println(s.Peek(1))

	//fmt.Println(s.IsEmpty(0), s.IsEmpty(1), s.IsEmpty(2))
	//fmt.Println(s.Peek(0), s.Peek(1), s.Peek(2))
	//s.Push(0, 1)
	//s.Push(0, 2)
	//s.Push(1, 1)
	//s.Push(1, 2)
	//s.Push(1, 3)
	//s.Push(1, 4)
	//fmt.Println(s.Peek(0), s.Peek(1), s.Peek(2))
	//fmt.Println(s.Pop(0), s.Pop(1), s.Pop(2))
	//fmt.Println(s.Peek(0), s.Peek(1), s.Peek(2))
	//fmt.Println(s.Pop(0))
	//fmt.Println(s.Pop(1))
	//fmt.Println(s.Pop(1))
	//fmt.Println(s.Pop(1))

}

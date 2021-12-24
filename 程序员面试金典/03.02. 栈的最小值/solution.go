package main

import (
	"container/list"
	"fmt"
)

type MinStack struct {
	stack    *list.List
	minStack *list.List
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack:    list.New(),
		minStack: list.New(),
	}
}

func (s *MinStack) Push(x int) {
	s.stack.PushBack(x)

	min := s.minStack.Back()
	if min == nil || min.Value.(int) >= x {
		s.minStack.PushBack(x)
	}
}

func (s *MinStack) Pop() {
	if s.stack.Len() == 0 {
		return
	}

	top, min := s.stack.Back(), s.minStack.Back()

	if top.Value.(int) == min.Value.(int) {
		s.minStack.Remove(min)
	}
	s.stack.Remove(top)

}

func (s *MinStack) Top() int {
	return s.stack.Back().Value.(int)
}

func (s *MinStack) GetMin() int {
	return s.minStack.Back().Value.(int)
}

func main() {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	fmt.Println(minStack.GetMin())
	minStack.Pop()
	fmt.Println(minStack.Top())
	fmt.Println(minStack.GetMin())

}

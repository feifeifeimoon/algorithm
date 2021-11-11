package main

import (
	"container/list"
	"fmt"
)

type MinStack struct {
	list    *list.List
	minList *list.List
}

// Constructor initialize your data structure here.
func Constructor() MinStack {
	return MinStack{
		list:    list.New(),
		minList: list.New(),
	}
}

func (s *MinStack) Push(x int) {
	s.list.PushBack(x)

	if s.minList.Len() == 0 {
		s.minList.PushBack(x)
	} else {
		min := s.minList.Back()
		// 比最小的还小 入栈
		if x <= min.Value.(int) {
			s.minList.PushBack(x)
		}
	}
}

func (s *MinStack) Pop() {
	elem := s.list.Back()
	min := s.minList.Back()
	if elem.Value.(int) == min.Value.(int) {
		s.minList.Remove(min)
	}
	s.list.Remove(elem)
}

func (s *MinStack) Top() int {
	return s.list.Back().Value.(int)
}

func (s *MinStack) Min() int {
	return s.minList.Back().Value.(int)
}

func main() {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	fmt.Println(minStack.Min())
	minStack.Pop()
	fmt.Println(minStack.Top())
	fmt.Println(minStack.Min())

}

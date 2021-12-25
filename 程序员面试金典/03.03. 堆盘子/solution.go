package main

import (
	"container/list"
	"fmt"
)

type StackOfPlates struct {
	stack *list.List
	cap   int
}

func Constructor(cap int) StackOfPlates {
	stack := StackOfPlates{
		stack: list.New(),
		cap:   cap,
	}
	//stack.stack.PushBack(list.New())

	return stack
}

func (s *StackOfPlates) Push(val int) {
	// ??? leetcode的用例 cap为0
	if s.cap == 0 {
		return
	}
	elem := s.stack.Back()
	if elem != nil {
		l := elem.Value.(*list.List)
		if l.Len() < s.cap {
			l.PushBack(val)
			return
		}
	}
	l := list.New()
	l.PushBack(val)
	s.stack.PushBack(l)
}

func (s *StackOfPlates) Pop() int {
	elem := s.stack.Back()
	if elem == nil {
		return -1
	}

	l := elem.Value.(*list.List)
	remove := l.Back()
	l.Remove(remove)

	// 如果长度为0就删除整个stack
	if l.Len() == 0 {
		s.stack.Remove(elem)
	}

	return remove.Value.(int)
}

func (s *StackOfPlates) PopAt(index int) int {
	if s.stack.Len() < index+1 {
		return -1
	}
	cur := s.stack.Front()
	for i := 0; i < index; i++ {
		cur = cur.Next()
	}

	l := cur.Value.(*list.List)
	remove := l.Back()
	l.Remove(remove)
	// 如果长度为0就删除整个stack
	if l.Len() == 0 {
		s.stack.Remove(cur)
	}

	return remove.Value.(int)
}

func main() {
	s := Constructor(2)
	s.Push(1)
	s.Push(2)
	s.Push(3)
	fmt.Println(s.PopAt(0))
	fmt.Println(s.PopAt(0))
	fmt.Println(s.PopAt(0))
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
}

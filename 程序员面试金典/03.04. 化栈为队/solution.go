package main

import (
	"container/list"
	"fmt"
)

type MyQueue struct {
	list1, list2 *list.List
}

// Constructor Initialize your data structure here.
func Constructor() MyQueue {
	return MyQueue{
		list1: list.New(),
		list2: list.New(),
	}
}

// Push element x to the back of queue.
func (q *MyQueue) Push(x int) {
	q.list1.PushBack(x)
}

// Pop Removes the element from in front of queue and returns that element.
func (q *MyQueue) Pop() int {
	if q.list2.Len() == 0 {
		// 进行一次两个栈的转变
		for cur := q.list1.Back(); cur != nil; cur = q.list1.Back() {
			q.list2.PushBack(cur.Value.(int))
			q.list1.Remove(cur)
		}
	}

	elem := q.list2.Back()
	q.list2.Remove(elem)
	return elem.Value.(int)
}

// Peek Get the front element
func (q *MyQueue) Peek() int {
	if q.list2.Len() == 0 {
		// 进行一次两个栈的转变
		for cur := q.list1.Back(); cur != nil; cur = q.list1.Back() {
			q.list2.PushBack(cur.Value.(int))
			q.list1.Remove(cur)
		}
	}

	elem := q.list2.Back()
	return elem.Value.(int)
}

// Empty Returns whether the queue is empty
func (q *MyQueue) Empty() bool {
	return q.list1.Len() == 0 && q.list2.Len() == 0
}

func main() {
	queue := Constructor()
	queue.Push(1)
	queue.Push(2)
	fmt.Println(queue.Empty())
	fmt.Println(queue.Peek())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Empty())
	queue.Push(3)
	fmt.Println(queue.Pop())
}

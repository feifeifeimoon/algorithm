package main

import (
	"container/list"
	"fmt"
)

type CQueue struct {
	list1 *list.List
	list2 *list.List
}

func Constructor() CQueue {
	return CQueue{
		list1: list.New(),
		list2: list.New(),
	}
}

func (c *CQueue) AppendTail(value int) {
	c.list1.PushBack(value)
}

func (c *CQueue) DeleteHead() int {
	value := -1

	// move element form list1 to list2
	if c.list2.Len() == 0 {
		for c.list1.Len() != 0 {
			elem := c.list1.Back()
			value = elem.Value.(int)
			c.list2.PushBack(elem.Value)
			c.list1.Remove(elem)
		}
	}

	if c.list2.Len() != 0 {
		elem := c.list2.Back()
		value = elem.Value.(int)
		c.list2.Remove(elem)
	}

	return value
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */

func main() {
	obj := Constructor()
	obj.AppendTail(3)
	fmt.Println(obj.DeleteHead())
	fmt.Println(obj.DeleteHead())
	fmt.Println(obj.DeleteHead())
}

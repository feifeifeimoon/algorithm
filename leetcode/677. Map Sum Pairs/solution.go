package main

import (
	"fmt"
	"strings"
)

type MapSum struct {
	store map[string]int
}

func Constructor() MapSum {
	return MapSum{
		store: make(map[string]int),
	}
}

func (m *MapSum) Insert(key string, val int) {
	m.store[key] = val
}

func (m *MapSum) Sum(prefix string) int {
	ret := 0
	for key, value := range m.store {
		if strings.HasPrefix(key, prefix) {
			ret += value
		}
	}

	return ret
}

func main() {
	m := Constructor()
	m.Insert("apple", 3)
	fmt.Println(m.Sum("ap")) // return 3 (apple = 3)
	m.Insert("app", 2)
	fmt.Println(m.Sum("ap")) // return 5 (apple + app = 3 + 2 = 5)

}

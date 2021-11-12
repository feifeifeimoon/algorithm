package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

var cache map[*Node]*Node

func copyNode(node *Node) *Node {
	if node == nil {
		return nil
	}

	if n, ok := cache[node]; ok {
		return n
	}

	newNode := &Node{
		Val: node.Val,
	}

	cache[node] = newNode

	newNode.Next = copyNode(node.Next)
	newNode.Random = copyNode(node.Random)
	return node
}

func copyRandomList(head *Node) *Node {
	cache = make(map[*Node]*Node)
	return copyNode(head)
}

func main() {

}

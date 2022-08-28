package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head   *Node
	length int
}

func (l *LinkedList) prepend(node *Node) {
	second := l.head
	l.head = node
	l.head.next = second
	l.length++
}

func (l LinkedList) printNode(node Node) {
	//var current Node
	for l.head != nil {
		fmt.Println(l.head)
	}
}

//func main() {
//	mylist := LinkedList{}
//	node1 := &Node{data: 48}
//	node2 := &Node{data: 40}
//	node3 := &Node{data: 16}
//
//	mylist.prepend(node1)
//	mylist.prepend(node2)
//	mylist.prepend(node3)
//
//	fmt.Println(mylist)
//
//}

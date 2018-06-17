package main

import "fmt"
/*
	structure based Stack
*/

type Node struct {
	data int
	next *Node
}


type Stack struct {
	head *Node
}

func (nd *Node)Create(num int) {
	nd.data = num
	nd.next = nil
}
func (stc *Stack) Push(data int) {
	temp := stc.head
	if temp == nil {
		node := new(Node)
		node.Create(data)
		stc.head = node
	} else {
		newNode := new(Node)
		newNode.Create(data)
		newNode.next = temp
		stc.head = newNode
	}
}

func (stc *Stack) Pop() int {
	if stc.Len() > 0 {
		val := stc.head.data
		stc.head = stc.head.next
		return val
	} else {
		return -1
	}
}

func (stc *Stack) Top() int {
	if stc.Len() > 0 {
		return stc.head.data
	} else {
		return -1
	}
}

func (stc *Stack) Len() int {
	temp := stc.head
	count := 0
	for temp != nil {
		count += 1
		temp = temp.next
	}
	return count
}

func (stc *Stack) isEmpty() bool {
	if stc.Len() == 0 {
		return true
	} else {
		return false
	}
}

func (stc *Stack) Print() {
	temp := stc.head
	for temp != nil {
		fmt.Println(temp.data)
		temp = temp.next
	}
}
func main() {
	head := new(Stack)
	head.Push(1)
	head.Push(2)
	head.Push(3)
	fmt.Println("Size of the stack", head.Len())
	head.Push(4)
	head.Push(5)
	head.Print()
	fmt.Println("Size of the stack", head.Top(), head.isEmpty())
}

package main

import (
	"fmt"
)

type Node struct {
	num int
	next *Node
}
//These all are the Property of new node
func (node *Node)GetData() {
	fmt.Println("Current Node :", node.num)
}

func (node *Node)GetNextNode() *Node{
	return node.next
}

func (node *Node)SetData(data int) {
	node.num = data
}

func (node *Node)SetNode(newNode *Node) {
	node.next = newNode
}

/*
	ListNode we can create a list which contain the node and the 
	Codintion for list:
		1. Each node only no details of its front user detail
		2. iF the link break the broke line is assume to be 
			last one.
*/

type LinkedList struct {
	head *Node
}

func (lst *LinkedList) PrintLst() {
	for lst.head.next != nil {
		fmt.Println(lst.head.num)
		lst.head = lst.head.next
	}
}

func (lst *LinkedList) InsertAtFront(num int) {
	head := lst.head
	temp := new(Node)
	temp.SetData(num)
	temp.SetNode(head)
	if head == nil {
		lst.head = temp
	} else {
		lst.head = temp
	}
}
func (lst *LinkedList) InsertAtEnd(num int) {
	temp := new(Node)
	temp.SetData(num)
	temp.SetNode(nil)
	head := lst.head
	if head == nil {
		lst.head = temp
	} else {
		for head.next != nil {
			head = head.next
		}
		head.next = temp
	}
} 

func (lst *LinkedList) Reverse() {
	hold := new(Node)
	temp := new(Node)
	head := lst.head
	//No need to check for empty
	for head != nil {
		head.GetData()
		temp = head.next
		head.next = hold
		hold = head
		head = temp
	}
	lst.head = hold
}

func main() {
	head := new(LinkedList)
	head.InsertAtEnd(5)
	head.InsertAtEnd(4)
	head.InsertAtFront(9)
	head.InsertAtEnd(3)
	head.InsertAtEnd(2)
	head.InsertAtEnd(1)
	head.InsertAtFront(11)
	//head.PrintLst()
	fmt.Println("Reverse the list : ")
	head.Reverse()
	head.PrintLst()
	fmt.Println("vim-go")
}

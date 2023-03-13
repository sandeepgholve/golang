package main

import (
	"fmt"
)

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head *Node
	Tail *Node
}

func (linkedList *LinkedList) Append(value int) {
	// referred with value here
	newNode := Node{Value: value}

	if linkedList.Head == nil {
		linkedList.Head = &newNode
		linkedList.Tail = &newNode
	} else {
		linkedList.Tail.Next = &newNode
		linkedList.Tail = &newNode
	}
}

func (linkedList *LinkedList) Prepend(value int) {
	// referred with reference here
	newNode := &Node{Value: value}

	if linkedList.Head == nil {
		linkedList.Head = newNode
		linkedList.Tail = newNode
	} else {
		newNode.Next = linkedList.Head
		linkedList.Head = newNode
	}
}

func (linkedList *LinkedList) Delete(value int) {
	if linkedList.Head == nil {
		return
	}

	if linkedList.Head.Value == value {
		linkedList.Head = linkedList.Head.Next
		return
	}

	currentNode := linkedList.Head
	for currentNode.Next != nil {
		if currentNode.Next.Value == value {
			currentNode.Next = currentNode.Next.Next
			return
		}
		currentNode = currentNode.Next
	}
}

func (linkedList *LinkedList) Find(value int) *Node {
	currentNode := linkedList.Head

	for currentNode != nil {
		if currentNode.Value == value {
			return currentNode
		}
		currentNode = currentNode.Next
	}
	return nil
}

func main() {
	linkedList := LinkedList{}
	linkedList.Append(1)
	linkedList.Append(2)
	linkedList.Append(3)
	linkedList.Append(4)

	printLinkedList(linkedList)

	//Find node with value 2
	findNode := linkedList.Find(2)
	if findNode != nil {
		fmt.Println("Found node with value 2")
	}

	//Delete node with value 2
	linkedList.Delete(2)
	printLinkedList(linkedList)

	//Find node with value 2
	node := linkedList.Find(2)
	if node == nil {
		fmt.Println("Not able to find node with value 2")
	}
}

func printLinkedList(linkedList LinkedList) {
	currnetNode := linkedList.Head
	for currnetNode != nil {
		fmt.Printf("'%d' --> ", currnetNode.Value)
		currnetNode = currnetNode.Next
	}
	fmt.Printf(" nil\n")
}

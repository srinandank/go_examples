package linkedlists

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head *Node
}

func CreateNode(value int) *Node {
	return &Node{Value: value, Next: nil}
}

func AppendToLinkedList(list *LinkedList, value int) {
	node := CreateNode(value)
	if list.Head == nil {
		list.Head = node
	} else {
		last := list.Head

		for last.Next != nil {
			last = last.Next
		}
		last.Next = node
	}
}

func CreateLinkedList(values []int) LinkedList {
	var linked_list LinkedList
	for _, value := range values {
		AppendToLinkedList(&linked_list, value)
	}
	return linked_list
}

func PrintLinkedList(name string, linked_list *LinkedList) {
	fmt.Printf("%s: ", name)
	last := linked_list.Head
	for last != nil {
		fmt.Printf("%d ", last.Value)
		last = last.Next
	}
	fmt.Println("")
}

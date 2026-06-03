package linkedlists

import "fmt"

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	Head *Node
}

func CreateNode(value int) *Node {
	return &Node{value: value, next: nil}
}

func AppendToLinkedList(list *LinkedList, value int) {
	node := CreateNode(value)
	if list.Head == nil {
		list.Head = node
	} else {
		last := list.Head

		for last.next != nil {
			last = last.next
		}
		last.next = node
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
		fmt.Printf("%d ", last.value)
		last = last.next
	}
	fmt.Println("")
}

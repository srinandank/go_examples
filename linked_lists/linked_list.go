package linkedlists

import (
	"fmt"
	"slices"
)

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

func SortLinkedList(listOfLinkedLists *[]LinkedList) {
	cmpfunc := func(a, b LinkedList) int {
		return a.Head.Value - b.Head.Value
	}
	slices.SortFunc(*listOfLinkedLists, cmpfunc)
}

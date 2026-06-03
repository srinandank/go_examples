package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	Head *Node
}

func createNode(value int) *Node {
	return &Node{value: value, next: nil}
}

func append_to_linked_list(list *LinkedList, value int) {
	node := createNode(value)
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

func createLinkedList(values []int) LinkedList {
	var linked_list LinkedList
	for _, value := range values {
		append_to_linked_list(&linked_list, value)
	}
	return linked_list
}

func printLinkedList(name string, linked_list *LinkedList) {
	fmt.Printf("%s: ", name)
	last := linked_list.Head
	for last != nil {
		fmt.Printf("%d ", last.value)
		last = last.next
	}
	fmt.Println("")
}

func sum_linked_list(list1 *LinkedList, list2 *LinkedList) LinkedList {
	last_first := list1.Head
	last_second := list2.Head
	var new_linked_list LinkedList
	quotient := 0
	for last_first != nil || last_second != nil {
		sum := 0
		if last_first != nil {
			sum += last_first.value
			last_first = last_first.next
		}
		if last_second != nil {
			sum += last_second.value
			last_second = last_second.next
		}
		sum += quotient
		remainder := sum % 10
		quotient = sum / 10
		append_to_linked_list(&new_linked_list, remainder)
	}
	return new_linked_list
}

func main() {
	list1 := createLinkedList([]int{1, 1, 3, 5, 8})
	list2 := createLinkedList([]int{5, 3, 8, 13})

	printLinkedList("First list", &list1)
	printLinkedList("Second list", &list2)

	output_linked_list := sum_linked_list(&list1, &list2)

	printLinkedList("Output List", &output_linked_list)

}

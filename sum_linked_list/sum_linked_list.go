package main

import "github.com/srinandank/go_examples/linked_lists"

func sum_linked_list(list1 *linked_lists.LinkedList, list2 *linked_lists.LinkedList) linked_lists.LinkedList {
	last_first := list1.Head
	last_second := list2.Head
	var new_linked_list linked_lists.LinkedList
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
		linked_lists.AppendToLinkedList(&new_linked_list, remainder)
	}
	return new_linked_list
}

func main() {
	list1 := linked_lists.CreateLinkedList([]int{1, 1, 3, 5, 8})
	list2 := linked_lists.CreateLinkedList([]int{5, 3, 8, 13})

	linked_lists.PrintLinkedList("First list", &list1)
	linked_lists.PrintLinkedList("Second list", &list2)

	output_linked_list := sum_linked_list(&list1, &list2)

	linked_lists.PrintLinkedList("Output List", &output_linked_list)

}

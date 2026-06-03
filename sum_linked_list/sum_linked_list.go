package main

import linkedlists "github.com/srinandank/go_examples/linked_lists"

func sum_linked_list(list1 *linkedlists.LinkedList, list2 *linkedlists.LinkedList) linkedlists.LinkedList {
	last_first := list1.Head
	last_second := list2.Head
	var new_linked_list linkedlists.LinkedList
	quotient := 0
	for last_first != nil || last_second != nil {
		sum := 0
		if last_first != nil {
			sum += last_first.Value
			last_first = last_first.Next
		}
		if last_second != nil {
			sum += last_second.Value
			last_second = last_second.Next
		}
		sum += quotient
		remainder := sum % 10
		quotient = sum / 10
		linkedlists.AppendToLinkedList(&new_linked_list, remainder)
	}
	return new_linked_list
}

func main() {
	list1 := linkedlists.CreateLinkedList([]int{1, 1, 3, 5, 8})
	list2 := linkedlists.CreateLinkedList([]int{5, 3, 8, 13})

	linkedlists.PrintLinkedList("First list", &list1)
	linkedlists.PrintLinkedList("Second list", &list2)

	output_linked_list := sum_linked_list(&list1, &list2)

	linkedlists.PrintLinkedList("Output List", &output_linked_list)

}

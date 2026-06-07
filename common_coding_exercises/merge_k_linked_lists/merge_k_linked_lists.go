package main

import (
	"container/heap"
	"fmt"

	linkedlists "github.com/srinandank/go_examples/common_coding_exercises/linked_lists"
)

type NodeHeap []*linkedlists.Node

// Methods necessary for heap operations: Len(), Less(), Swap()
func (h *NodeHeap) Len() int {
	return len(*h)
}

func (h *NodeHeap) Less(i, j int) bool {
	return (*h)[i].Value <= (*h)[j].Value
}

func (h *NodeHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *NodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func (h *NodeHeap) Push(x interface{}) {
	*h = append(*h, x.(*linkedlists.Node))
}

func core_merger(listOfLinkedLists *[]linkedlists.LinkedList) linkedlists.Node {
	linkedlists.SortLinkedList(listOfLinkedLists)
	var head *linkedlists.Node
	var curNode *linkedlists.Node
	node_heap := &NodeHeap{}
	for _, node := range *listOfLinkedLists {
		heap.Push(node_heap, node.Head)
	}
	for node_heap.Len() > 0 {
		// O(1)
		n := heap.Pop(node_heap).(*linkedlists.Node)
		if n.Next != nil {
			// O(logK)
			heap.Push(node_heap, n.Next)
		}
		if head == nil {
			head = n
			curNode = head
		} else {
			curNode.Next = n
			curNode = curNode.Next
		}
	}
	return *head
}

func merge_linked_list(listOfLinkedLists *[]linkedlists.LinkedList) linkedlists.Node {
	for index, list_linked_lists := range *listOfLinkedLists {
		heading := fmt.Sprintf("Input List %d", index+1)
		linkedlists.PrintLinkedList(heading, &list_linked_lists)
	}
	return core_merger(listOfLinkedLists)
}

func main() {
	list1 := linkedlists.CreateLinkedList([]int{1, 4, 5})
	list2 := linkedlists.CreateLinkedList([]int{1, 3, 4})
	list3 := linkedlists.CreateLinkedList([]int{2, 6})
	list_linked_lists := []linkedlists.LinkedList{list1, list2, list3}
	output_node := merge_linked_list(&list_linked_lists)
	output_linked_list := linkedlists.LinkedList{Head: &output_node}
	linkedlists.PrintLinkedList("Output List", &output_linked_list)
}

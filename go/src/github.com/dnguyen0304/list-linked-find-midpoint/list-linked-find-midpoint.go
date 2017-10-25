// Package main implements an algorithm to find the midpoint of a linked list
// data structure.
//
// # What type of linked list is this? Would it be acceptable to assume the
//   linked list is a singly linked list?
// # How is the linked list represented in memory? Would it be acceptable to
//   assume the linked list is a facade over a collection of nodes?
// # Does the linked list maintain its size?
// # What is the algorithm passed? Would it be acceptable to assume the
//   algorithm is passed the linked list and subsequently a pointer to the
//   head node?
// # What should the algorithm return? Would it be acceptable to assume the
//   algorithm returns a pointer to the midpoint node?
//
// # What type of data does the data structure contain?
// # Am I allowed to use other data structures?
// # Am I allowed to implement other data structures?
// # Does the entire data structure fit into memory?
//
// How do you implement generic containers?
// How do you implement Iterables and Iterators?
//
// TODO(duy): Add error handling for linked lists with 0 nodes.
package main

import (
	"errors"
	"fmt"
	"time"
)

var NoMoreElements = errors.New(
	"There are no more elements over which to iterate.")

// Node is an integer singly linked list node.
type Node struct {
	data int
	// This field must be defined as of type pointer to avoid an invalid
	// recursive type error. This works because the compiler can calculate the
	// size of pointers but not of recursive types.
	next *Node
}

// iterator is an iterator for integer singly linked lists.
//
// Its operations are not thread-safe.
type iterator struct {
	head *Node
	current *Node
}

// Next gets the next element in the iteration.
//
// This returns NoMoreElements when there are no more elements over which to
// iterate. The time complexity is O(1).
func (iterator *iterator) Next() (*Node, error) {
	// Case: start
	if iterator.current == nil {
		iterator.current = iterator.head
		return iterator.current, nil
	}
	// Case: stop
	if iterator.current.next == nil {
		return nil, NoMoreElements
	}
	// Case: iterate
	temporary := iterator.current.next
	iterator.current = iterator.current.next
	return temporary, nil
}

// NewIterator creates an iterator for integer singly linked lists.
//
// The time complexity is O(1).
func NewIterator(head *Node) *iterator {
	return &iterator{head: head, current: nil}
}

// LinkedList is an integer singly linked list.
//
// Its operations are not thread-safe.
type LinkedList struct {
	head *Node
	tail *Node
}

// Add adds a node to the tail of the integer singly linked list.
//
// The time complexity is O(1).
func (list *LinkedList) Add(data int) {
	next := &Node{data: data, next: nil}

	// Case: 0-node linked list
	if list.head == nil {
		list.head = next
	}
	// Case: 0 or 1-node linked list
	if list.tail == nil {
		list.tail = next
	// Case: 2 or more node linked list
	} else {
		list.tail.next = next
		list.tail = list.tail.next
	}
}

// Iterate gets an iterator.
//
// The time complexity is O(1).
func (list *LinkedList) Iterate() *iterator {
	return NewIterator(list.head)
}

// NewLinkedList creates an integer singly linked list of the given size.
//
// The time complexity is O(n), where n is the number of elements in the
// singly linked list.
func NewLinkedList(length int) *LinkedList {
	if length < 0 {
		return nil
	}

	list := &LinkedList{head: nil, tail: nil}

	for i := 0; i < length; i++ {
		list.Add(i)
	}

	return list
}

// FindMidpoint finds the midpoint node of the given linked list.
//
// If there are an even number of nodes, then the one to the left of the
// midpoint will be returned. For example, if there are 4 nodes in the linked
// list, then the node at index 2 will be returned. The time complexity is
// O(n), where n is the number of elements in the linked list.
func FindMidpoint(list *LinkedList) *Node {
	// Variables that are declared but not initialized are zero valued.
	var midpoint *Node
	var size int

	iterator := list.Iterate()

	for current, err := iterator.Next(); err != NoMoreElements; current, err = iterator.Next() {
		size++
		// Case: null midpoint node
		if midpoint == nil {
			midpoint = current
		}
		// Case: even number of nodes
		if size % 2 == 0 {
			midpoint = midpoint.next
		}
	}

	return midpoint
}

func main() {
	// Rhodes' "Thousand-Million' Thought Experiment
	for _, length := range []int{10, 1000, 1000000, 10000000} {
		list := NewLinkedList(length)
		start := time.Now()
		FindMidpoint(list)
		elapsed := time.Since(start).Seconds()
		// Width assigns a fixed width whereas precision determines at least
		// how many digits are displayed after the decimal point. The latter
		// may pad zeroes to whole numbers.
		fmt.Printf(
			"When n is equal to %8d, the operation takes %.9f seconds.\n",
			length,
			elapsed)
	}
}

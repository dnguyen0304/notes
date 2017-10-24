// Package main implements an algorithm to marshall a list into a tree data
// structure.
//
// # What type of list is this? Would it be acceptable to assume the list is a
//   dynamically-sized array?
// # Is the list sorted? Would it be acceptable to assume the list is sorted?
// # What type of tree is this? Would it be acceptable to assume the tree is a
//   binary tree?
// # How is the tree represented in memory? Would it be acceptable to assume
//   the tree is akin to a linked list where the algorithm is passed the head
//   node (or the root node in this case) and traverses it as needed?
//
// # What type of data does the data structure contain?
// # Am I allowed to use other data structures?
// # Am I allowed to implement other data structures?
// # Does the entire data structure fit into memory?
//
// How do you declare generic containers?
package main

import (
	"errors"
	"fmt"
	"time"
)

var ROOT_INDEX = 0

var ErrEmptyQueue = errors.New("The queue is empty.")
var ErrNodeConflict = errors.New("There is already an existing node.")

// ArrayQueue is a node queue backed by an array (i.e. slice).
//
// The container package and subsequently the vector data structure were
// removed from the standard library in Go version 1.0. The recommendation
// moving forward is to use channels or slices.
//
// The vector in Go should not be confused with the vector in Java. The latter
// is a dynamically-sized array. Unlike the ArrayList though, it is
// synchronized to some degree, resizes by 100% to amortize insertion cost, and
// conventionally is viewed as deprecated.
type ArrayQueue struct {
	array []*Node
}

// Size gets the number of elements in the queue.
//
// The time complexity is O(1). This operation is not thread-safe.
func (queue *ArrayQueue) Size() int {
	return len(queue.array)
}

// Push adds an element to the tail of the queue.
//
// The time complexity is O(1) amortized depending on the language
// implementation. This operation is not thread-safe.
func (queue *ArrayQueue) Push(element *Node) {
	queue.array = append(queue.array, element)
}

// Pop removes an element from the head of the queue.
//
// This returns ErrEmptyQueue if the queue is empty. The time complexity is
// O(1). This operation is not thread-safe.
func (queue *ArrayQueue) Pop() (*Node, error) {
	if queue.Size() == 0 {
		return nil, ErrEmptyQueue
	}
	element := queue.array[0]
	queue.array = queue.array[1:]
	return element, nil
}

// NewArrayQueue creates a node queue.
//
// The time complexity is O(1).
func NewArrayQueue() *ArrayQueue {
	// Because the size is now known, a zero-valued slice variable is declared
	// rather than declaring a slice literal or using the make function.
	var array []*Node
	queue := ArrayQueue{array: array}
	return &queue
}

// Node is a node for an integer binary tree.
//
// The fields are defined as pointer types to avoid an "invalid recursive
// type" error. This works because the compiler can now calculate the size of
// pointers
type Node struct {
	data	int
	left	*Node
	right	*Node
}

// LeftAdd adds a node to the binary tree as the left child.
//
// This returns ErrNodeConflict if there is already an existing node. The time
// complexity is O(1). This operation is not thread-safe.
func (node *Node) LeftAdd(other *Node) error {
	// Base Case: null other node
	if other == nil {
		return nil
	}
	// Base Case: existing left node
	if node.left != nil {
		return ErrNodeConflict
	}
	// Base Case: null left node
	node.left = other
	return nil
}

// RightAdd adds a node to the binary tree as the right child.
//
// This returns ErrNodeConflict if there is already an existing node. The time
// complexity is O(1). This operation is not thread-safe.
func (node *Node) RightAdd(other *Node) error {
	// Base Case: null other node
	if other == nil {
		return nil
	}
	// Base Case: existing right node
	if node.right != nil {
		return ErrNodeConflict
	}
	// Base Case: null right node
	node.right = other
	return nil
}

// NewRange creates an integer slice of the given size and with sequential
// values.
//
// The time complexity is O(n), where n is the number of elements in the slice.
func NewRange(length int) []int {
	list := make([]int, length)
	for i := range list {
		list[i] = i
	}
	return list
}

// MarshallListToTree marshalls a list into a complete binary tree.
//
// The list must be sorted. The time complexity is O(n), where n is the number
// of elements in the list.
func MarshallListToTree(list []int) *Node {
	// Base Case: null or zero-valued list
	if list == nil || len(list) == 0 {
		return nil
	}

	root := &Node{data: list[ROOT_INDEX], left: nil, right: nil}
	queue := NewArrayQueue()
	queue.Push(root)

	var current *Node
	var err error
	var left *Node
	var right *Node

	for i, j := ROOT_INDEX+1, ROOT_INDEX+2; i < len(list); i, j = i+2, j+2 {
		current, err = queue.Pop()
		if err == ErrEmptyQueue {
			break
		}

		left = &Node{data: list[i], left: nil, right: nil}
		current.LeftAdd(left)
		queue.Push(current.left)

		if j < len(list) {
			right = &Node{data: list[j], left: nil, right: nil}
			current.RightAdd(right)
			queue.Push(current.right)
		}
	}

	return root
}

func main() {
	// Rhodes' "Thousand-Million" Thought Experiment
	for _, length := range []int{10, 1000, 1000000, 10000000} {
		list := NewRange(length)
		start := time.Now()
		MarshallListToTree(list)
		elapsed := time.Since(start).Seconds()
		// Width assigns a fixed width whereas precision determines at least
		// how many digits are displayed after the decimal point.
		fmt.Printf(
			"When n is equal to %8d, the operation takes %.9f seconds.\n",
			length,
			elapsed)
	}
}

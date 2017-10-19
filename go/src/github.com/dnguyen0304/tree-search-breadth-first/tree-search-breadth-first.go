// Package main implements a breadth-first-search algorithm for tree data
// structures.
//
// What type of tree is this? Would it be acceptable to assume the tree is a
// binary tree? How is the tree represented in memory? Would it be acceptable
// to assume the tree is akin to a linked list where the algorithm is passed
// the head node (or the root node in this case) and traverses it as needed?
// Am I allowed to use other data structures? Am I allowed to implement other
// data structures? Would you prefer an iterative or recursive approach? Does
// the entire tree fit in memory? What should be done with the node upon
// visiting it? Would it be acceptable merely to print its data?
//
// How do you declare generic containers?
package main

import (
	"errors"
	"fmt"
)

var ErrEmptyQueue = errors.New("The queue is empty.")

// ArrayQueue is a queue backed by an array (i.e. slice).
//
// Go version 1.0 removes the container package from the standard library and
// subsequently the vector. The recommendation moving forward is to use
// channels or slices.
//
// The vector in Go should not be confused with that in Java. The Vector in
// Java is a dynamically-sized array. Unlike the ArrayList, it is synchronized
// to some degree, resizes by 100% to amortize insertion cost, and
// conventionally is viewed as deprecated.
type ArrayQueue struct {
	array []*Node
}

// Push adds an element to the tail of the queue.
//
// The time complexity is O(1) amortized depending on the language
// implementation. This operation is not thread-safe.
func (queue *ArrayQueue) Push(node *Node) {
	queue.array = append(queue.array, node)
}

// Pop removes an element from the head of the queue.
//
// This method returns ErrEmptyQueue if the queue is empty. The time
// complexity is O(1). This operation is not thread-safe.
func (queue *ArrayQueue) Pop() (*Node, error) {
	if len(queue.array) == 0 {
		return nil, ErrEmptyQueue
	}
	node := queue.array[0]
	queue.array = queue.array[1:]
	return node, nil
}

// Node is a string node for a binary tree.
type Node struct {
    // The fields must be defined as pointers to avoid a "invalid recursive
	// type" error. This works because the compiler can calculate the size of
    // pointers.
	data string
	left *Node
	right *Node
}

func NewTree() *Node {
	//   A
    //  / \
	// B   C
	//    / \
	//   D   E
	//    \
	//     F

	a := Node{"A", nil, nil}
	b := Node{"B", nil, nil}
	c := Node{"C", nil, nil}
	d := Node{"D", nil, nil}
	e := Node{"E", nil, nil}
	f := Node{"F", nil, nil}

	a.left = &b
	a.right = &c

	c.left = &d
	c.right = &e

	d.left = &f

	return &a
}

// SearchTree traverses the tree using breadth-first-search starting at the
// given node.
//
// The time complexity is O(n), where n is the number of nodes in the tree.
func SearchTree(queue *ArrayQueue, node *Node) {
	// Base Case
	if node == nil {
		return
	}

	// Iterative Case
	queue.Push(node)

	for current, err := queue.Pop(); err == nil; current, err = queue.Pop() {
		if current.left != nil {
			queue.Push(current.left)
		}
		if current.right != nil {
			queue.Push(current.right)
		}
		fmt.Println(current.data)
	}
}

func main() {
	// When the size is not known, declare a zero-valued slice variable rather
	// than declaring a slice literal or using the make function.
	var array []*Node
	queue := &ArrayQueue{array: array}

	tree := NewTree()

	SearchTree(queue, tree)
}

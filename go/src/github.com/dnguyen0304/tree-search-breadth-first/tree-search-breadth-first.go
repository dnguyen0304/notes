package main

import (
	"errors"
	"fmt"
)

var ErrEmptyQueue = errors.New("The queue is empty.")

// ArrayQueue is a queue data structure backed by an array (i.e. slice). Go
// version 1.0 removes the container package from the standard library and
// subsequently the vector data structure. The recommendation moving forward
// is to use channels or slices.
type ArrayQueue struct {
	array []*Node
}

// Push adds an element to the tail of the queue. This operation is not
// thread-safe. The time complexity is O(1) amortized depending on the language
// implementation.
func (queue *ArrayQueue) Push(node *Node) {
	queue.array = append(queue.array, node)
}

// Pop removes an element from the head of the queue. This method returns
// ErrEmptyQueue if the queue is empty. This operation is not thread-safe. The
// time complexity is O(1).
func (queue *ArrayQueue) Pop() (*Node, error) {
	if len(queue.array) == 0 {
		return nil, ErrEmptyQueue
	}
	node := queue.array[0]
	queue.array = queue.array[1:]
	return node, nil
}

// Node is a string node data structure.
type Node struct {
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

// SearchTree traverses the tree data structure using breadth-first searching.
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
	// When there is no specified size, declare a zero value slice variable
	// rather than declaring a slice literal or using the make function.
	var array []*Node
	queue := &ArrayQueue{array: array}

	tree := NewTree()

	SearchTree(queue, tree)
}

// package main implements an algorithm to determine if a tree data structure
// is full.
//
// # What type of tree is this? Would it be acceptable to assume the tree is
//   a binary tree?
// # How is the tree represented in memory? Would it be acceptable to assume
//   the algorithm is passed an adjacency matrix? Would it be acceptable to
//   assume the first row in the adjacency matrix is the root?
//
// # What type of data does the data structure contain?
// # Am I allowed to use other data structures?
// # Am I allowed to implement other data structures?
// # Does the entire tree fit into memory?
//
// How do you declare generic containers?
package main

import (
	"errors"
	"fmt"
)

var ROOT_INDEX = 0
var FULL_CHILDREN_COUNT = 2
var LEAF_CHILDREN_COUNT = 0

var ErrEmptyQueue = errors.New("The queue is empty.")

// ArrayQueue is a integer queue backed by an array (i.e. slice).
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
	array []int
}

// Push adds a integer element to the tail of the queue.
//
// The time complexity is O(1) amortized depending on the language
// implementation. This operation is not thread-safe.
func (queue *ArrayQueue) Push(element int) {
	queue.array = append(queue.array, element)
}

// Pop removes a integer element from the head of the queue.
//
// This returns ErrEmptyQueue if the queue is empty. The time complexity is
// O(1). This operation is not thread-safe.
func (queue *ArrayQueue) Pop() (int, error) {
	if len(queue.array) == 0 {
		return 0, ErrEmptyQueue
	}
	element := queue.array[0]
	queue.array = queue.array[1:]
	return element, nil
}

// NewArrayQueue creates a integer queue.
//
// The time complexity is O(1).
func NewArrayQueue() *ArrayQueue {
	// When the size is not known, declare a zero-valued slice variable rather
	// than declaring a slice literal or using the make function.
	var array []int
	queue := ArrayQueue{array: array}
	return &queue
}

// IsFull asserts whether the given tree is full.
//
// The time complexity is O(n^2), where n is the number of nodes in the tree.
func IsFull(tree [][]bool) bool {
	// Base Case: null tree
	if tree == nil {
		return false
	}

	// Iterative Case
	isFull := true
	count := 0
	queue := NewArrayQueue()
	queue.Push(ROOT_INDEX)

	for current, err := queue.Pop(); err != ErrEmptyQueue; current, err = queue.Pop() {
		count = 0
		for i, isAdjacent := range tree[current] {
			if isAdjacent {
				count++
				queue.Push(i)
			}
		}
		if count != FULL_CHILDREN_COUNT && count != LEAF_CHILDREN_COUNT {
			isFull = false
			break
		}
	}

	return isFull
}

// NewFullTree creates a boolean tree.
//
// The tree is full. The time complexity is O(1).
func NewFullTree() [][]bool {
	//      A
	//    /   \
	//   B     C
	//  / \   / \
	// D   E F   G

	rowA := []bool{false, true, true, false, false, false, false}
	rowB := []bool{false, false, false, true, true, false, false}
	rowC := []bool{false, false, false, false, false, true, true}
	rowD := []bool{false, false, false, false, false, false, false}
	rowE := []bool{false, false, false, false, false, false, false}
	rowF := []bool{false, false, false, false, false, false, false}
	rowG := []bool{false, false, false, false, false, false, false}

	tree := [][]bool{rowA, rowB, rowC, rowD, rowE, rowF, rowG}

	return tree
}

// NewNotFullTree creates a boolean tree.
//
// The tree is not full. The time complexity is O(1).
func NewNotFullTree() [][]bool {
	//   A
	//  / \
	// B   C
	//    / \
	//   D   E
	//    \
	//     F

	rowA := []bool{false, true, true, false, false, false}
	rowB := []bool{false, false, false, false, false, false}
	rowC := []bool{false, false, false, true, true, false}
	rowD := []bool{false, false, false, false, false, true}
	rowE := []bool{false, false, false, false, false, false}
	rowF := []bool{false, false, false, false, false, false}

	tree := [][]bool{rowA, rowB, rowC, rowD, rowE, rowF}

	return tree
}

func main() {
	var tree [][]bool

	tree = NewFullTree()
	fmt.Println(IsFull(tree))

	tree = NewNotFullTree()
	fmt.Println(IsFull(tree))
}

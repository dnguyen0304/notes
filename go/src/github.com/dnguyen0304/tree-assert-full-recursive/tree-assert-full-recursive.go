// Package main implements an algorithm to assert whether a tree data
// structure is full.
//
// # What type of tree is this? Would it be acceptable to assume the tree is
//   a binary tree?
// # How is the tree represented in memory? What it be acceptable to assume
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
	"fmt"
)

var ROOT_INDEX = 0
var FULL_CHILDREN_COUNT = 2
var LEAF_CHILDREN_COUNT = 0

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

// AssertFull asserts whether the given tree is full.
//
// The time complexity is O(n^2), where n is the number of nodes in the tree.
func AssertFull(tree [][]bool) bool {
	// Base Case: null tree
	if tree == nil {
		return false
	}

	return assertFull(tree, ROOT_INDEX)
}

func assertFull(tree [][]bool, current int) bool {
	// Base Case: null tree
	if tree == nil {
		return false
	}

	isFull := true
	children := make([]int, 0, FULL_CHILDREN_COUNT)

	for node, isAdjacent := range tree[current] {
		if isAdjacent {
			children = append(children, node)
		}
	}

	// Base Case: 0-child node (leaf)
	if len(children) == LEAF_CHILDREN_COUNT {
		return true
	}
	// Base Case: 1-child node
	if len(children) != FULL_CHILDREN_COUNT {
		return false
	}
	// Recursive Case: 2-child node
	for _, child := range children {
		if !assertFull(tree, child) {
			isFull = false
			break
		}
	}

	return isFull
}

func main() {
	var tree [][]bool

	tree = NewFullTree()
	fmt.Println(AssertFull(tree))

	tree = NewNotFullTree()
	fmt.Println(AssertFull(tree))
}

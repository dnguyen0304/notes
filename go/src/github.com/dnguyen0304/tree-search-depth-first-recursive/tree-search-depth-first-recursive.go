// Package main implements a depth-first-search algorithm for tree data
// structures.
//
// # What type of tree is this? Would it be acceptable to assume the tree is
//   a binary tree?
// # How is the tree represented in memory? Would it be acceptable to assume
//   the tree is akin to a linked list where the algorithm is passed the head
//   node (or root node in this case) and traverses as needed?
// # What type of data does the tree's nodes contain?
// # What should be done with the node upon visiting it? Would it be
//   acceptable merely to print its data?
// # Am I allowed to use other data structures?
// # Am I allowed to impelement other data structures?
// # Does the entire tree fit in memory?
package main

import (
	"fmt"
)

// Node is a string node for a binary tree.
type Node struct {
	// The fields must be defined as pointers to avoid an invalid recursive
	// type error. This works because the compiler can now calculate the size
	// of pointers.
	data  string
	left  *Node
	right *Node
}

func NewTree() *Node {
	//   A
	//  / \
	// B   C
	//    / \
	//   D   F
	//    \
	//     E

	a := Node{data: "A", left: nil, right: nil}
	b := Node{data: "B", left: nil, right: nil}
	c := Node{data: "C", left: nil, right: nil}
	d := Node{data: "D", left: nil, right: nil}
	e := Node{data: "E", left: nil, right: nil}
	f := Node{data: "F", left: nil, right: nil}

	a.left = &b
	a.right = &c

	c.left = &d
	c.right = &f

	d.right = &e

	return &a
}

func SearchTree(node *Node) {
	// Base Case
	if node == nil {
		return
	}

	fmt.Println(node.data)

	// Recursive Case
	if node.left != nil {
		SearchTree(node.left)
	}
	// Recursive Case
	if node.right != nil {
		SearchTree(node.right)
	}
}

func main() {
	tree := NewTree()
	SearchTree(tree)
}

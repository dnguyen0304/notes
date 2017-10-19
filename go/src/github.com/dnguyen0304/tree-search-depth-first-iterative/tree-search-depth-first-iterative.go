// Package main implements an iterative depth-first-search algorithm for tree
// data structures.
//
// # What type of tree is this? Would it be acceptable to assume the tree is
//   a binary tree?
// # How is the tree represented in memory? Would it be acceptable to assume
//   the tree is akin to a linked list where the algorithm is passed the head
//   node (or root node in this case) and traverses as needed?
// # Am I allowed to use other data structures?
// # Am I allowed to implement other data structures?
// # Am I allowed to use iteration instead of recursion?
// # What type of data does the tree's nodes contain?
// # Does the entire tree fit in memory?
// # What should be done with the node upon visiting it? Would it be
//   acceptable merely to print its data?
package main

import (
	"errors"
	"fmt"
)

var ErrEmptyStack = errors.New("The stack is empty.")

// ArrayStack is a stack backed by an array (i.e. slice).
type ArrayStack struct {
	array []*Node
}

// Push adds an element to the head of the stack.
//
// The time complexity is O(1) amoritized depending on the language
// implementation. This operation is not thread-safe.
func (stack *ArrayStack) Push(node *Node) {
	stack.array = append(stack.array, node)
}

// Pop removes an element from the head of the stack.
//
// This method returns ErrEmptyStack if the stack is empty. The time
// complexity is O(1). This operation is not thread-safe.
func (stack *ArrayStack) Pop() (*Node, error) {
	if len(stack.array) == 0 {
		return nil, ErrEmptyStack
	}
	node := stack.array[len(stack.array)-1]
	stack.array = stack.array[:len(stack.array)-1]
	return node, nil
}

// Node is a string node for a binary tree.
type Node struct {
	// The fields must be defined as pointers to avoid an invalid recursive
	// type error. This works because the compiler can calculate the size of
	// pointers.
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

// SearchTree traverses the tree using depth-first-search starting at the
// given node.
//
// The time complexity is O(n), where n is the number of nodes in the tree.
func SearchTree(stack *ArrayStack, node *Node) {
	// Base Case
	if node == nil {
		return
	}

	// Iterative Case
	stack.Push(node)

	for current, err := stack.Pop(); err == nil; current, err = stack.Pop() {
		fmt.Println(current.data)
		if current.right != nil {
			stack.Push(current.right)
		}
		if current.left != nil {
			stack.Push(current.left)
		}
	}
}

func main() {
	// When the size is not known, declare a zero-valued slice variable rather
	// than declaring a slice literal or using the make function.
	var array []*Node
	stack := &ArrayStack{array: array}

	tree := NewTree()

	SearchTree(stack, tree)
}

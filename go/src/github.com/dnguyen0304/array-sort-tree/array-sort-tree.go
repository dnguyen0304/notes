package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Define the fields as pointers to avoid the "invalid recursive type" error.
// This technique works because the compiler can calculate the size of
// pointers.
type Node struct {
	value int
	left *Node
	right *Node
}

// Add adds the new node to the tree. This method receiver must be a pointer
// receiver. The parameter type could be of type Node pointer to be more
// space efficient.
func (node *Node) Add(other Node) {
	if other.value < node.value {
		if node.left == nil {
			// Base Case #1
			node.left = &other
			return
		} else {
			// Recursive Case #1
			node.left.Add(other)
		}
	} else {
		if node.right == nil {
			// Base Case #2
			node.right = &other
		} else {
			// Recursive Case #2
			node.right.Add(other)
		}
	}
}

// randomRange creates a slice with the specified length and then fills it with
// randomly generated integers.
func randomRange(length int) []int {
	list := make([]int, length)
	for i, _ := range list {
		list[i] = rand.Intn(length)
	}
	return list
}

// TreeSort sorts the array in-place. The worst-case time complexity is n^2
// (quadratic). However, this only occurs when the array is already sorted or
// nearly sorted. The average-case time complexity is n * log(n).
func TreeSort(list []int) {
	// Create a binary search tree. The worst-case time complexity is
	// n * log(n), where n is the number of nodes.
	root := Node{list[0], nil, nil}
	for _, value := range list[1:] {
		node := Node{value, nil, nil}
		root.Add(node)
	}
	// Traverse the tree (in-order). The worst-case time complexity is n.
	nodes := traverse(root)
	// Sort the array. The array must not be used for iteration as it is being
	// mutated in-place. The worst-case time complexity is n.
	for i, node := range nodes {
		list[i] = node.value
	}
}

// Traverse the tree in-order.
func traverse(node Node) []Node {
	var nodes []Node

	// Base Case
	if node.left == nil && node.right == nil {
		nodes = append(nodes, node)
		return nodes
	}

	// Recursive Case
	if node.left != nil {
		nodes = append(nodes, traverse(*node.left)...)
	}
	nodes = append(nodes, node)
	if node.right != nil {
		nodes = append(nodes, traverse(*node.right)...)
	}
	return nodes
}

func main() {
	seconds := time.Now().UTC().Unix()
	rand.Seed(seconds)

	for _, length := range []int{10, 100, 1000, 1000000} {
		list := randomRange(length)
		TreeSort(list)
		start := time.Now()
		TreeSort(list)
		elapsed := time.Since(start).Seconds()
		fmt.Printf(
			"When n is equal to %7d, the operation takes %.9f seconds.\n",
			length,
			elapsed)
	}
}

package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type Node struct {
	data int
	left *Node
	right *Node
}

// Add adds the other node to the tree. The time complexity is O(n) where n
// is the size of the tree.
func (node *Node) Add(other *Node) {
	// Base Case #1
	if other == nil {
		return
	}

	if other.data < node.data {
		if node.left == nil {
			// Base Case #2
			node.left = other
		} else {
			// Recursive Case #1
			node.left.Add(other)
		}
	} else {
		if node.right == nil {
			// Base Case #3
			node.right = other
		} else {
			// Recursive Case #2
			node.right.Add(other)
		}
	}
}

// newRandom creates a rand.Rand that uses the current time as its seed. This
// object is not thread-safe. The time complexity is O(1).
func newRandom() *rand.Rand {
	seed := time.Now().UTC().Unix()
	source := rand.NewSource(seed)
	random := rand.New(source)
	return random
}

// newList creates a slice according to the specified length. The values are
// sorted. The time complexity is O(n^2), where n is the length of the array.
func newList(random *rand.Rand, length int) []int {
	list := make([]int, length)
	tree := newTree(random, length)
	nodes := traverse(tree)
	for i, node := range nodes {
		list[i] = node.data
	}
	return list
}

// newTree creates a binary tree according to the specified size and where the
// nodes contain pseudo-randomly generated integers. The time complexity is
// O(n^2), where n is the size of the tree.
func newTree(random *rand.Rand, size int) *Node {
	// The time complexity is O(1).
	root := Node{random.Intn(size), nil, nil}
	// The time complexity is O(n).
	for i := 0; i < size - 1; i++ {
		node := Node{random.Intn(size), nil, nil}
		// The time complexity is O(n).
		root.Add(&node)
	}
	return &root
}

// traverse visits the nodes using in-order tree traversal. The time
// complexity is O(n) where n is the size of the tree.
func traverse(node *Node) []*Node {
	var nodes []*Node

	// Base Case #1
	if node == nil {
		return nodes
	}

	if node.left != nil {
		// Recursive Case #1
		nodes = append(nodes, traverse(node.left)...)
	}

	// Base Case #2
	nodes = append(nodes, node)

	if node.right != nil {
		// Recursive Case #2
		nodes = append(nodes, traverse(node.right)...)
	}

	return nodes
}

// BlockSearch searches the list for the specified value in blocks. If the
// value is found, it returns the index of the first match. If the value is
// not found, it returns -1. The time complexity is O(n^(1/2)), where n is
// the length of the array.
func BlockSearch(list []int, other int) int {
	interval := int(math.Sqrt(float64(len(list))))

	// Base Case
	if len(list) == 0 {
		return -1
	}
	// Iterative Case
	for start, end, current := 0, interval, 0; ; start, end = start + interval, end + interval {
		if end > len(list) {
			end = len(list)
		}
		if current = list[end - 1]; current == other {
			return current
		} else if current > other {
			return start + LinearSearch(list[start:end], other)
		}
	}
}

// LinearSearch searches the list for the specified value. If the value is
// found, it returns the index of the first match. Otherwise, it returns -1.
// The time complexity is O(n), where n is the length of the array.
func LinearSearch(list []int, other int) int {
	for i, data := range list {
		if data == other {
			return i
		}
	}
	return -1
}

func main() {
	random := newRandom()

	for _, length := range []int{10, 100, 1000, 1000000} {
		list := newList(random, length)
		position := random.Intn(length)
		other := list[position]

		start := time.Now().UTC()
		BlockSearch(list, other)
		elapsed := time.Since(start).Seconds()

		message := fmt.Sprintf(
			"When n is equal to %7d, the operation takes %.9f seconds.",
			length,
			elapsed)
		fmt.Println(message)
	}
}

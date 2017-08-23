package main

type Node struct {
	data int
	left *Node
	right *Node
}

func Height(node Node) int {
	// Base Case
	if node == nil {
		return 0
	}

	// Recursive Case
	left = Height(*node.left)
	right = Height(*node.right)

	// The math.Max function takes arguments of type float64.
	if left >= right {
		return left + 1
	} else {
		return right + 1
	}
}

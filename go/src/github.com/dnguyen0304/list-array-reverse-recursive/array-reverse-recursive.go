// main implements a reversal algorithm for array data structures.
//
// # Should the array be reversed in-place?
// # What type of data do the array's elements contain?
// # Am I allowed to use other data structures?
// # Does the entire array fit into memory?
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// NewRandomRange creates an integer slice of the given size and with random
// values.
//
// The time complexity is O(n), where n is the number of elements in the slice.
func NewRandomRange(length int) []int {
	list := make([]int, length)
	for i := range list {
		list[i] = rand.Intn(length)
	}
	return list
}

// Reverse reverses an integer slice.
//
// The operation is applied in-place. The time complexity is O(n), where n is
// the number of elements in the slice.
func Reverse(list []int) {
	reverse(list, 0, len(list)-1)
}

func reverse(list []int, left int, right int) {
	// Base Case
	if list == nil {
		return
	}
	// Base Case
	if left >= right {
		return
	}
	// Recursive Case
	list[left], list[right] = list[right], list[left]
	reverse(list, left+1, right-1)
}

func main() {
	seconds := time.Now().UTC().Unix()
	// This is global state. However, the standard library does provide APIs
	// for creating local sources and random number generators.
	rand.Seed(seconds)

	// Rhodes' "Thousand-Million" Thought Experiment
	for _, length := range []int{10, 100, 1000, 1000000} {
		list := NewRandomRange(length)
		start := time.Now()
		Reverse(list)
		elapsed := time.Since(start).Seconds()
		// Width assigns a fixed width whereas precision determines at least
		// how many digits are displayed after the decimal point. The latter
		// may pad zeroes to whole numbers.
		fmt.Printf(
			"When n is equal to %7d, the operation takes %.9f seconds.\n",
			length,
			elapsed)
	}
}

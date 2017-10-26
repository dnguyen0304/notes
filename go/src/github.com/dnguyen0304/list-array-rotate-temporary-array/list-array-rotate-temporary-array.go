// package main implements a rotation algorithm for array data structures.
//
// # Should the array be rotated to the left or to the right? Would it be
//   acceptable if I use left rotation?
// # Should the array be rotated in-place?
// # Should iteration or recursion be used? Would it be acceptable if I use
//   iteration?
// # What type of data do the array's elements contain?
// # Am I allowed to use other data structures?
// # Does the entire array fit into memory?
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// NewRange creates an integer slice of the given size and with sequential
// values.
//
// The time complexity is O(n), where n is the number of elements in the slice.
func NewRange(length int) []int {
	list := make([]int, length)
	for i := range list {
		list[i] = i
	}
	return list
}

// LeftRotate left rotates an integer slice by the given factor.
//
// The operation is applied in-place. The time complexity is O(n), where n is
// the number of elements in the slice.
func LeftRotate(list []int, factor int) {
	// Base Case
	if list == nil {
		return
	}
	// Base Case
	// Only pointer-types (i.e. pointers, slices, channels, interfaces, maps,
	// and functions) may be compared to nil.
	if factor == 0 {
		return
	}
	// Iterative Case
	temporary := make([]int, factor)
	copy(temporary, list[:factor])

	for i := 0; i < len(list)-factor; i++ {
		list[i] = list[i+factor]
	}

	copy(list[len(list)-factor:], temporary)
}

// RightRotate right rotates an integer slice by the given factor.
//
// The operation is applied in-place. The time complexity is O(n), where n is
// the number of elements in the slice.
func RightRotate(list []int, factor int) {
	LeftRotate(list, len(list)-factor)
}

func main() {
	seconds := time.Now().UTC().Unix()
	// This is global state. However, the standard library does provide APIs
	// for creating local sources and random number generators.
	rand.Seed(seconds)

	// Rhodes' "Thousand-Million" Thought Experiment
	for _, n := range []int{10, 1000, 1000000, 10000000} {
		list := NewRange(n)
		factor := rand.Intn(n)
		start := time.Now()
		LeftRotate(list, factor)
		elapsed := time.Since(start).Seconds()
		// Width assigns a fixed width whereas precision determines at least
		// how many digits are displayed after the decimal point. The latter
		// may pad zeroes to whole numbers.
		fmt.Printf(
			"When n is equal to %7d, the operation takes %.9f seconds.\n",
			n,
			elapsed)
	}
}

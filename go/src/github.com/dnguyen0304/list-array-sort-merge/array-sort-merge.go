package main

import (
	"fmt"
	"math/rand"
	"time"
)

func RandomRange(length int) []int {
	list := make([]int, length)
	for i, _ := range list {
		list[i] = rand.Intn(length)
	}
	return list
}

func MergeSort(list []int) []int {
	// Base Case
	if len(list) == 0 || len(list) == 1 {
		return list
	}
	// Recursive Case
	// Arithmetic operations have the same return type as the first operand.
	// Therefore, if the slice contains an odd number of elements, then the
	// return value will be floor rounded.
	midpoint := len(list) / 2
	left := list[:midpoint]
	right := list[midpoint:]
	l := merge(MergeSort(left), MergeSort(right))
	return l
}

func merge(left []int, right []int) []int {
	merged := make([]int, len(left)+len(right))
	for i, j, k := 0, 0, 0; i < len(left) || j < len(right); k++ {
		if i == len(left) {
			merged[k] = right[j]
			j++
		} else if j == len(right) {
			merged[k] = left[i]
			i++
		} else if left[i] <= right[j] {
			merged[k] = left[i]
			i++
		} else if left[i] > right[j] {
			merged[k] = right[j]
			j++
		}
	}
	return merged
}

func main() {
	seconds := time.Now().UTC().Unix()
	rand.Seed(seconds)

	for _, length := range []int{10, 100, 1000, 1000000} {
		list := RandomRange(length)
		start := time.Now()
		MergeSort(list)
		elapsed := time.Since(start).Seconds()
		fmt.Printf(
			"When n is equal to %7d, the operation takes %.9f seconds.\n",
			length,
			elapsed)
	}
}

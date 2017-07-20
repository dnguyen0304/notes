package main

import (
    "fmt"
    "math/rand"
    "time"
)

func RandomRange(length int) []int {
    list := make([]int, length)
    for i := range list {
        list[i] = rand.Intn(length)
    }
    return list
}

func BubbleSort(list []int) []int {
    swaps := -1
    for swaps != 0 {
        swaps = 0
        for i := 0; i < len(list) - 1; i++ {
            if list[i] > list[i + 1] {
                list[i], list[i + 1] = list[i + 1], list[i]
                swaps++
            }
        }
    }
    return list
}

func main() {
    // Set the global seed for the random number generator.
    // To print the current timestamp formatted according to RFC 3339,
    // use the Format method and the time.RFC3339 constant.
    seconds := time.Now().UTC().Unix()
    rand.Seed(seconds)

    // Rhodes "thousand-million" Thought Experiment
    for _, length := range []int{10, 100, 1000, 1000000} {
        list := RandomRange(length)
        start := time.Now()
        BubbleSort(list)
        elapsed := time.Since(start).Seconds()
        fmt.Printf("When n is equal to %7d, the operation takes %.9f seconds.\n",
                   length,
                   elapsed)
    }
}

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
    // The complexity is approximately a magnitude worse than factorial
    // time (n^2). Having to copy the source array is probably what
    // introduces the additional overhead.
    tmp := make([]int, len(list))
    copy(tmp, list)

    // This private function mutates its argument in-place.
    bubbleSort(tmp)

    return tmp
}

func bubbleSort(list []int) {
    // Base Case
    if len(list) == 0 || len(list) == 1 {
        // BEST PRACTICE: "Return early."
        // BEST PRACTICE: "Note if nothing is supposed to happen."
        return
    }
    // BEST PRACTICE: Avoid the else clause.
    // BEST PRACTICE: Flat is better than nested.
    // Recursive Case
    for i := 0; i < len(list) - 1; i++ {
        if list[i] > list[i + 1] {
            list[i], list[i + 1] = list[i + 1], list[i]
        }
    }
    bubbleSort(list[:len(list) - 1])
}

func main() {
    seconds := time.Now().UTC().Unix()
    rand.Seed(seconds)

    // Rhodes "thousand-million" Thought Experiment
    for _, length := range []int{10, 100, 1000, 1000000} {
        list := RandomRange(length)
        start := time.Now()
        BubbleSort(list)
        elapsed := time.Since(start).Seconds()
        // Width assigns a fixed width whereas precision determines at
        // least how many digits are displayed are the decimal point.
        // The latter may pad zeroes to whole numbers.
        fmt.Printf("When n is equal to %7d, the operation takes %.9f seconds.\n",
                   length,
                   elapsed)
    }
}
